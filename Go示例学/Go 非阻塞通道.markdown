# Go 非阻塞通道
默认情况下，通道发送和接收数据是阻塞的。然而我们可以使用select的一个default的选项来实现无阻塞发送或接收数据，甚至可以将多个select的case选项和default选项结合起来使用。
```go
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// 这里是一个非阻塞的从通道接收数据，如果messages通道有数据
	// 可以接收，那么select将运行`<-messages`这个case，否则的话
	// 程序立刻执行default选项后面的语句
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 非阻塞通道发送数据也是一样的
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 在default前面，我们可以有多个case选项，从而实现多通道
	// 非阻塞的选择，这里我们尝试从messages和signals接收数据
	// 如果有数据可以接收，那么执行对应case后面的逻辑，否则立刻
	// 执行default选项后面的逻辑
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
```
运行结果
```
no message received
no message sent
no activity
```
这个例子中，由于我们使用了default来实现非阻塞的通道，所以开始的时候messages里面没有数据可以接收，直接输出`no message received`，而第二次由于messages通道没有相应的数据接收方，所以也不会写入数据，直接转到default，输出`no message sent`，至于第三个就很好理解了，什么也没有，输出`no activity`。
其实，我们可以把这个例子修改一下，让messages通道带缓冲区，这样例子或许更好理解一点。定义messages的时候使用`messages := make(chan string, 1)`。
```go
package main

import "fmt"

func main() {
	messages := make(chan string, 1)
	signals := make(chan bool)

	// 这里是一个非阻塞的从通道接收数据，如果messages通道有数据
	// 可以接收，那么select将运行`<-messages`这个case，否则的话
	// 程序立刻执行default选项后面的语句
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 非阻塞通道发送数据也是一样的,但是由于messages带了缓冲区，
	// 即使没有数据接受端也可以发送数据，所以这里的`messages<-msg`
	// 会被执行，从而不再跳到default去了。
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 在default前面，我们可以有多个case选项，从而实现多通道
	// 非阻塞的选择，这里我们尝试从messages和signals接收数据
	// 如果有数据可以接收，那么执行对应case后面的逻辑，否则立刻
	// 执行default选项后面的逻辑
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
```
运行结果
```
no message received
sent message hi
received message hi
```