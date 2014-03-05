# Go 打点器
Timer是让你等待一段时间然后去做一件事情，这件事情只会做一次。而Ticker是让你按照一定的时间间隔循环往复地做一件事情，除非你手动停止它。

```go
package main

import "time"
import "fmt"

func main() {

	// Ticker使用和Timer相似的机制，同样是使用一个通道来发送数据。
	// 这里我们使用range函数来遍历通道数据，这些数据每隔500毫秒被
	// 发送一次，这样我们就可以接收到
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// Ticker和Timer一样可以被停止。一旦Ticker停止后，通道将不再
	// 接收数据，这里我们将在1500毫秒之后停止
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
```
输出结果
```
Tick at 2014-02-18 05:42:50.363640783 +0800 CST
Tick at 2014-02-18 05:42:50.863793985 +0800 CST
Tick at 2014-02-18 05:42:51.363532887 +0800 CST
Ticker stopped
```
在这个例子中，我们让Ticker一个独立协程上每隔500毫秒执行一次，然后在main函数所在协程上等待1500毫秒，然后停止Ticker。所以只输出了3次`Ticker at`信息。