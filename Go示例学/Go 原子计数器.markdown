# Go 原子计数器
Go里面的管理协程状态的主要机制就是通道通讯。这些我们上面的例子介绍过。这里还有一些管理状态的机制，下面我们看看多协程原子访问计数器的例子，这个功能是由sync/atomic包提供的函数来实现的。
```go
package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {

	// 我们使用一个无符号整型来代表一个永远为正整数的counter
	var ops uint64 = 0

	// 为了模拟并行更新，我们使用50个协程来每隔1毫秒来
	// 增加一下counter值，注意这里的50协程里面的for循环，
	// 也就是说如果主协程不退出，这些协程将永远运行下去
	// 所以这个程序每次输出的值有可能不一样
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 为了能够保证counter值增加的原子性，我们使用
				// atomic包中的AddUint64方法，将counter的地址和
				// 需要增加的值传递给函数即可
				atomic.AddUint64(&ops, 1)

				// 允许其他的协程来处理
				runtime.Gosched()
			}
		}()
	}

	//等待1秒中，让协程有时间运行一段时间
	time.Sleep(time.Second)

	// 为了能够在counter仍被其他协程更新值的同时安全访问counter值，
	// 我们获取一个当前counter值的拷贝，这里就是opsFinal，需要把
	// ops的地址传递给函数`LoadUint64`
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
```
我们多运行几次，结果如下:
```
ops: 7499289
ops: 7700843
ops: 7342417
```