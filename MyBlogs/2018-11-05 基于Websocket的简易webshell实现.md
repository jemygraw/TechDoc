我们在很多场合都看到过基于浏览器的 shell，你可以在里面输入一些和你本机相同的命令，然后从远程服务器获得对应的输出。

本篇文章就是用来讲解这个基于 web 的 shell 的实现方法的。我们之所以研究这个问题，另一方面也是因为 kubernetes 的 Dashboard 里面也包含了这个功能。
在研究 kubernetes 的 Dashboard 的时候，我们会发现那个功能是基于 WebSocket 来实现的。所以本篇也就是讲解基于 WebSocket 的 shell 实现方法。

## 思路

我们如果仔细地思考一下，其实这个 web shell 的主要功能就是将这个命令发送到远程服务器，然后远程服务器执行这个命令，然后把结果返回给客户端就可以了。所以在这个客户端和服务器的交互场景下，有很多的方案可以选择，比如直接使用 HTTP 协议或者使用 TCP 协议，那么为什么 kubernetes 在实现的使用使用 web socket协议呢？

在进行一个技术方案的选型的时候，最重要的就是深入了解各个方案的利弊，以及它们最适用的场景。所以我们可以对比下基于 TCP，HTTP 和 WebSocket 三种协议实现这个 webshell 的优缺点。

|协议|类别|特点|
|---|---|---|
|WebSocket|七层(应用层)|兼容HTTP的80端口和HTTPS的443端口，可以运行在HTTP或HTTPS协议上，全双工协议，基于事件驱动的交互方式，客户端不需要轮询服务端的执行结果|
|HTTP(s)|七层(应用层)|HTTP需要保持长连接来维持客户端和服务器之间的不断的命令执行交互，否则频繁的短连接性能损耗严重，另外客户端需要主动轮训服务端的执行结果|
|TCP|四层(传输层)|我就是个裸的传输层协议啦，HTTPS(s)和WebSocket都最终依赖我|


从原理上讲，大家最后都是需要依赖TCP协议来进行数据传输，所以如果坚持用TCP协议实现 webshell 当然是可以的，没有任何问题。
但是协议的抽象目的就是简化问题的解决方案以及解决旧有方案的缺点，HTTP的出现就是一种规范化的TCP协议应用，否则按照大家各自定义自己的数据格式的做法，这个互联网还是不要搞了，没法搞。你能脑补出每个公司每天都在互相接入对方的协议开发自己的应用么？画面太美，不敢想，不敢想。

那么我们就看看 WebSokcet 的出现简化和解决了 HTTP 协议的哪些问题就可以了。

对于交互式场景的应用，最重要的就是等待回复的不确定性，比如你发个消息给对方，对方什么时候回复是不确定的，你执行了一个远端的命令，这个命令什么时候执行完毕也是不确定的，在HTTP协议中，解决这种不确定性的方案是什么呢？轮询！你不是没有办法告诉我么？我自己去问行不？可以，来问吧，周期性地询问一下。

轮询的做法有什么问题呢？首先就是轮询周期的设定，你怎么设定这个时间呢？周期太短，白白浪费那么多建立连接，断开连接的动作，就像很久以前谈恋爱的小伙子，每隔一分钟去问一下传达室的大爷，今天刚寄出的信有没有回复。大爷很快就口吐白沫了。但是要是每隔一天去问，看上去好像不错，但是万一那个信是上午到的，结果你下午才去拿，那白白痛苦等待了一个上午不是。所以这个方案不好，但是没办法。

刚刚说了轮询的第一个缺点，第二个就是你总是轮询，整个路上全是你来来往往的身影，占用带宽不是，浪费连接不是。

那么 WebSocket 的出现，就可以解决这个问题了，大爷说，小伙子信发出去了不着急，等信到了大爷通知你，甚至主动把信送给你，你看好不好。

这就完美解决问题了嘛。

## 代码

BB那么久，好歹说清楚 WebSocket 哪里好了，简单贴个能跑的代码吧。能跑就是说能展示原理，但是别直接拷贝就上线了，不好。

先上客户端，就是模拟每次发一个命令过去，然后命令后面接上换行符，算是简单的协议格式。


remoteshell-client.go

```
package main

import (
	"flag"
	"golang.org/x/net/websocket"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var origin string
	var url string
	flag.StringVar(&origin, "origin", "", "websocket origin")
	flag.StringVar(&url, "url", "", "websocket remote url")
	flag.Parse()

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		panic(err)
		return
	}
	buffer := make([]byte, 40960)
	bScanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for bScanner.Scan() {
		line := bScanner.Text()
		ws.Write([]byte(line + "\r\n"))
		num, err := ws.Read(buffer)
		if err != nil {
			ws.Close()
			return
		}
		fmt.Println(string(buffer[:num]))
		fmt.Print("> ")
	}
}
```

执行方法：

```
$ ./remoteshell-client -url 'ws://localhost:9001/remote/shell' -origin 'http://localhost:9001'
```

服务端就是处理这个请求并给个回复了，因为这个连接是一直都在的，所以读数据就是直接for循环去读。我们在客户端发送数据的时候，给每条数据加了一个换行符，所以服务端就可以按行来读了。

```
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os/exec"
	"bytes"
	"strings"
	"golang.org/x/net/websocket"
	"bufio"
	"os"
)

func RemoteShell(ws *websocket.Conn) {
	bScanner := bufio.NewScanner(ws)
	currentWorkingDir, _ := os.Getwd()
	fmt.Println("current working dir", currentWorkingDir)

	for bScanner.Scan() {
		// parse command
		cmd := bScanner.Text()
		fmt.Println(cmd)
		cmdItems := strings.Split(cmd, " ")
		cmdName := cmdItems[0]

		var cmdArgs []string
		if len(cmdItems) >= 2 {
			cmdArgs = cmdItems[1:]
		}

		// execute command
		cmdOutput := bytes.NewBuffer(nil)
		cmdExec := exec.Command(cmdName, cmdArgs...)
		cmdExec.Dir = currentWorkingDir
		cmdExec.Stdout = cmdOutput
		cmdExec.Stderr = cmdOutput
		err := cmdExec.Run()
		if err != nil {
			fmt.Println(err)
			ws.Write([]byte(err.Error()))
		} else {
			ws.Write(cmdOutput.Bytes())
		}
	}
}

func main() {
	var host string
	var port int
	flag.StringVar(&host, "host", "0.0.0.0", "host to listen")
	flag.IntVar(&port, "port", 9001, "port to listen")
	flag.Parse()

	//handler
	http.Handle("/remote/shell", websocket.Handler(RemoteShell))

	//listen
	endPoint := fmt.Sprintf("%s:%d", host, port)
	err := http.ListenAndServe(endPoint, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
```

执行方法：

```
$ ./remoteshell-server 
```