# Go 时间戳
程序的一个通常需求是计算从Unix起始时间开始到某个时刻的秒数，毫秒数，微秒数等。
我们来看看Go里面是怎么做的。
```go
package main

import "fmt"
import "time"

func main() {

	// 使用Unix和UnixNano来分别获取从Unix起始时间
	// 到现在所经过的秒数和微秒数
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// 注意这里没有UnixMillis方法，所以我们需要将
	// 微秒手动除以一个数值来获取毫秒
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// 反过来，你也可以将一个整数秒数或者微秒数转换
	// 为对应的时间
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
```
运行结果
```
2014-03-02 23:11:31.118666918 +0800 CST
1393773091
1393773091118
1393773091118666918
2014-03-02 23:11:31 +0800 CST
2014-03-02 23:11:31.118666918 +0800 CST
```