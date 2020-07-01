# Go HTTP 默认的Client对象使用注意事项

## 前言

前面写过一篇帖子介绍了Go HTTP中重用TCP连接需要注意的问题，看上去阅读量很好。所有打算再介绍下Go HTTP中使用默认的Client时需要注意的事项。

## 问题

在一个系统中，通过HTTP调用第三方的服务是很常见的，而且稍微大的一点的系统，可能需要调用不同的组件来完成一项工作。这个时候就需要注意一些细节问题。

我们知道Go的`net/http`库里面有很多方法可以直接使用。比如常见的 `http.Get`和`http.Post`。一般情况下我们用这两个方法来写个脚本什么的都没啥问题。但是需要注意的是在大型系统中，千万不要直接使用这些方法。具体原因就在这些方法的定义中。

### http.Get

```go
func Get(url string) (resp *Response, err error) {
	return DefaultClient.Get(url)
}
```

### http.Post

```go
func Post(url, contentType string, body io.Reader) (resp *Response, err error) {
	return DefaultClient.Post(url, contentType, body)
}
```

### http.DefaultClient

上面的两个方法里面都使用了`http.DefaultClient`对象，这个对象是`net/http`包里面提供的可以即时使用的HTTP Client对象，它的定义如下：

```go
// DefaultClient is the default Client and is used by Get, Head, and Post.
var DefaultClient = &Client{}
```

问题出在哪里？问题就出在这个`DefaultClient`是个指针，换句话说这个对象在 Go 的并发编程中是不安全的。

## 实验

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 初始化超时时间为 1 秒
	http.DefaultClient.Timeout = time.Second
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		count := 1
		for {
			select {
			case <-ticker.C:
				// 每隔 5 秒，更新一下超时时间
				http.DefaultClient.Timeout = time.Second * time.Duration(count)
				count++
			}
		}
	}()

	// 不断请求 Google，会触发超时，如果没有超时，说明你已经违法，😄
	for i := 0; i < 100; i++ {
		startTime := time.Now()
		func() {
			resp, err := http.Get("https://www.google.com")
			if err != nil {
				return
			}
			defer resp.Body.Close()
		}()

		// 打印下运行数据，开始时间，超时时间
		fmt.Println(fmt.Sprintf("Run %d:", i+1), "Start:", startTime.Format("15:04:05"),
			"Timeout:", time.Since(startTime))

		// 每隔 1 秒请求一次
		<-time.After(time.Second)
	}
}
```

运行情况：

```s
Run 1: Start: 21:37:42 Timeout: 1.002390001s
Run 2: Start: 21:37:44 Timeout: 1.005189409s
Run 3: Start: 21:37:46 Timeout: 1.001791553s
Run 4: Start: 21:37:48 Timeout: 1.000847131s
Run 5: Start: 21:37:50 Timeout: 1.0042284s
Run 6: Start: 21:37:52 Timeout: 2.001313209s
Run 7: Start: 21:37:55 Timeout: 2.000255175s
Run 8: Start: 21:37:58 Timeout: 3.005502974s
Run 9: Start: 21:38:02 Timeout: 4.005494172s
Run 10: Start: 21:38:07 Timeout: 5.001988372s
Run 11: Start: 21:38:13 Timeout: 6.000908119s
Run 12: Start: 21:38:20 Timeout: 7.003262543s
Run 13: Start: 21:38:28 Timeout: 9.000410503s
Run 14: Start: 21:38:38 Timeout: 11.004758151s
Run 15: Start: 21:38:50 Timeout: 13.002290813s
```

其实不需要这个例子，你也能够明白在并发环境不能直接用共享的变量，否则会出问题的。比如A系统超时时间和B系统超时时间完全不同，结果因为用了共享的 `http.DefaultClient`，就混在一起了。

## 小结

学无止境，小心翼翼。