# Go通道的同步功能
我们使用通道来同步协程之间的执行。
下面的例子是通过获取同步通道数据来阻塞程序执行的方法来等待另一个协程运行结束的。
也就是说main函数所在的协程在运行到`<-done`语句的时候将一直等待worker函数所在的协程执行完成，向通道写入数据才会（从通道获得数据）继续执行。

```go
package main

import "fmt"
import "time"

// 这个worker函数将以协程的方式运行
// 通道`done`被用来通知另外一个协程这个worker函数已经执行完成
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 向通道发送一个数据，表示worker函数已经执行完成
	done <- true
}

func main() {

	// 使用协程来调用worker函数，同时将通道`done`传递给协程
	// 以使得协程可以通知别的协程自己已经执行完成
	done := make(chan bool, 1)
	go worker(done)

	// 一直阻塞，直到从worker所在协程获得一个worker执行完成的数据
	<-done
}
```
运行结果
```
working...done
```
如果我们从main函数里面移除`<-done`语句，那么main函数在worker协程开始运行之前就结束了。