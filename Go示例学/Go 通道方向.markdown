# Go通道方向
当使用通道作为函数的参数时，你可以指定该通道是只读的还是只写的。这种设置有时候会提高程序的参数类型安全。

```go
package main

import "fmt"

// 这个ping函数只接收能够发送数据的通道作为参数，试图从这个通道接收数据
// 会导致编译错误，这里只写的定义方式为`chan<- string`表示这个类型为
// 字符串的通道为只写通道
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong函数接收两个通道参数，一个是只读的pings，使用`<-chan string`定义
// 另外一个是只写的pongs，使用`chan<- string`来定义
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
```
运行结果
```
passed message
```
其实这个例子就是把信息首先写入pings通道里面，然后在pong函数里面再把信息从pings通道里面读出来再写入pongs通道里面，最后在main函数里面将信息从pongs通道里面读出来。
在这里，pings和pongs事实上是可读且可写的，不过作为参数传递的时候，函数参数限定了通道的方向。不过pings和pongs在ping和pong函数里面还是可读且可写的。只是ping和pong函数调用的时候把它们当作了只读或者只写。