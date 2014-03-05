# Go 互斥
上面的例子中，我们看过了如何在多个协程之间原子地访问计数器，对于更复杂的例子，我们可以使用`Mutex`来在多个协程之间安全地访问数据。

```go
package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)

func main() {

    // 这个例子的状态就是一个map
    var state = make(map[int]int)

    // 这个`mutex`将同步对状态的访问
    var mutex = &sync.Mutex{}

    // ops将对状态的操作进行计数
    var ops int64 = 0
    // 这里我们启动100个协程来不断地读取这个状态
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {

                // 对于每次读取，我们选取一个key来访问，
                // mutex的`Lock`函数用来保证对状态的
                // 唯一性访问，访问结束后，使用`Unlock`
                // 来解锁，然后增加ops计数器
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)

                // 为了保证这个协程不会让调度器出于饥饿状态，
                // 我们显式地使用`runtime.Gosched`释放了资源
                // 控制权，这种控制权会在通道操作结束或者
                // time.Sleep结束后自动释放。但是这里我们需要
                // 手动地释放资源控制权
                runtime.Gosched()
            }
        }()
    }
    // 同样我们使用10个协程来模拟写状态
    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
                runtime.Gosched()
            }
        }()
    }

    // 主协程Sleep，让那10个协程能够运行一段时间
    time.Sleep(time.Second)

    // 输出总操作次数
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)

    // 最后锁定并输出状态
    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}
```
运行结果
```go
ops: 3931611
state: map[0:84 2:20 3:18 1:65 4:31]
```