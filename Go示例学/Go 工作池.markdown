# Go 工作池
在这个例子中，我们来看一下如何使用gorouotine和channel来实现工作池。

```go
package main

import "fmt"
import "time"

// 我们将在worker函数里面运行几个并行实例，这个函数从jobs通道
// 里面接受任务，然后把运行结果发送到results通道。每个job我们
// 都休眠一会儿，来模拟一个耗时任务。
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {

    // 为了使用我们的工作池，我们需要发送工作和接受工作的结果，
    // 这里我们定义两个通道，一个jobs，一个results
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // 这里启动3个worker协程，一开始的时候worker阻塞执行，因为
    // jobs通道里面还没有工作任务
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // 这里我们发送9个任务，然后关闭通道，告知任务发送完成
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // 然后我们从results里面获得结果
    for a := 1; a <= 9; a++ {
        <-results
    }
```
运行结果
```
worker 1 processing job 1
worker 2 processing job 2
worker 3 processing job 3
worker 1 processing job 4
worker 3 processing job 5
worker 2 processing job 6
worker 1 processing job 7
worker 3 processing job 8
worker 2 processing job 9
```