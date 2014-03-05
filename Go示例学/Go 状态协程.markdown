# Go 状态协程
在上面的例子中，我们演示了如何通过使用mutex来在多个协程之间共享状态。另外一种方法是使用协程内置的同步机制来实现。这种基于通道的方法和Go的通过消息共享内存，保证每份数据为单独的协程所有的理念是一致的。

```go
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 在这个例子中，将有一个单独的协程拥有这个状态。这样可以
// 保证这个数据不会被并行访问所破坏。为了读写这个状态，其
// 他的协程将向这个协程发送信息并且相应地接受返回信息。
// 这些`readOp`和`writeOp`结构体封装了这些请求和回复
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// 我们将计算我们执行了多少次操作
	var ops int64 = 0

	// reads和writes通道将被其他协程用来从中读取或写入数据
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// 这个是拥有`state`的协程，`state`是一个协程的私有map
	// 变量。这个协程不断地`select`通道`reads`和`writes`，
	// 当有请求来临的时候进行回复。一旦有请求，首先执行所
	// 请求的操作，然后给`resp`通道发送一个表示请求成功的值。
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 这里启动了100个协程来向拥有状态的协程请求读数据。
	// 每次读操作都需要创建一个`readOp`，然后发送到`reads`
	// 通道，然后等待接收请求回复
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// 我们开启10个写协程
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// 让协程运行1秒钟
	time.Sleep(time.Second)

	// 最后输出操作数量ops的值
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
}
```
运行结果
```
ops: 880578
```
运行这个程序，我们会看到基于协程的状态管理每秒可以处理800, 000个操作。对于这个例子来讲，基于协程的方法比基于mutex的方法更加复杂一点。当然在某些情况下还是很有用的。例如你有很多复杂的协程，而且管理多个mutex可能导致错误。
当然你可以选择使用任意一种方法，只要你保证这种方法让你觉得很舒服而且也能保证程序的正确性。