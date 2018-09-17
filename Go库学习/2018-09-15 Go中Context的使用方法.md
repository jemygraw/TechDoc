在一般的项目中，存在一个函数调用另外一个函数的情况，而这另外的一个函数还会再去调用其他的函数。一般项目越复杂，调用的层级就也多。这个时候，如果让这种调用关系一直执行到底，那么也没有什么问题，如果有错误发生就一级一级往上抛。

在服务端的编程中，可能会存在一种情况，在一个请求进来的处理过程中，后端的服务会依赖多个其他的微服务的API，在这种级连的服务调用过程中，假设需要取消这个过程的执行，那么就会需要多个API的调用同时取消并且完成相应的清理工作。

在 Golang 中，标准库提供了一个协程安全的包 context，这个包里面提供了几个有用的方法来实现这个过程。

其中一个叫 context.WithCancel 可以用来派生出一个新的 context 对象和一个cancelFunc函数，该派生出的新的 context 对象的方法 Done() 返回一个用于标志取消操作的 channel 对象，当函数cancelFunc被调用的时候，派生出的 context 对象在调用 Done() 函数时会收到取消的信号。

下面我们用一个例子来演示这个过程，我们让一个主 goroutine 发送这个取消的信号，然后等待每个子协程完成取消操作，最后返回。

```
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("work%d", i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(ctx, name)
		}()
	}
	<-time.After(time.Second * 10)
	cancelFunc()

	//wait for all the work done
	wg.Wait()
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			//quit the work when done() fired
			fmt.Println("quit", name)
			return
		default:
			<-time.After(time.Second * 1)
		}
		fmt.Println(name, time.Now().String())
	}
}
```