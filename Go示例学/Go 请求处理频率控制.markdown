# Go 请求处理频率控制
频率控制是控制资源利用和保证服务高质量的重要机制。Go可以使用goroutine，channel和ticker来以优雅的方式支持频率控制。
```go
package main

import "time"
import "fmt"

func main() {

	// 首先我们看下基本的频率限制。假设我们得控制请求频率，
	// 我们使用一个通道来处理所有的这些请求，这里向requests
	// 发送5个数据，然后关闭requests通道
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 这个limiter的Ticker每隔200毫秒结束通道阻塞
	// 这个limiter就是我们频率控制处理器
	limiter := time.Tick(time.Millisecond * 200)

	// 通过阻塞从limiter通道接受数据，我们将请求处理控制在每隔200毫秒
	// 处理一个请求，注意`<-limiter`的阻塞作用。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 我们可以保持正常的请求频率限制，但也允许请求短时间内爆发
	// 我们可以通过通道缓存来实现，比如下面的这个burstyLimiter
	// 就允许同时处理3个事件。
	burstyLimiter := make(chan time.Time, 3)

	// 填充burstyLimiter，先发送3个数据
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 然后每隔200毫秒再向burstyLimiter发送一个数据，这里是不断地
	// 每隔200毫秒向burstyLimiter发送数据
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 这里模拟5个请求，burstyRequests的前面3个请求会连续被处理，
	// 因为burstyLimiter被先连续发送3个数据的的缘故，而后面两个
	// 则每隔200毫秒处理一次
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
```
运行结果
```
request 1 2014-02-21 14:20:05.2696437 +0800 CST
request 2 2014-02-21 14:20:05.4696637 +0800 CST
request 3 2014-02-21 14:20:05.6696837 +0800 CST
request 4 2014-02-21 14:20:05.8697037 +0800 CST
request 5 2014-02-21 14:20:06.0697237 +0800 CST
request 1 2014-02-21 14:20:06.0697237 +0800 CST
request 2 2014-02-21 14:20:06.0697237 +0800 CST
request 3 2014-02-21 14:20:06.0707238 +0800 CST
request 4 2014-02-21 14:20:06.2707438 +0800 CST
request 5 2014-02-21 14:20:06.4707638 +0800 CST
```
我们从输出的结果上可以看出最后的5个输出结果中，前三个的时间是连续的，而后两个的时间是隔了200毫秒。