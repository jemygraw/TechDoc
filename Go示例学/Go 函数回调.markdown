# Go 函数回调

Go支持函数回调，你可以把函数名称作为参数传递给另外一个函数，然后在别的地方实现这个函数。

```go
package main

import "fmt"

type Callback func(x, y int) int

func main() {
	x, y := 1, 2
	fmt.Println(test(x, y, add))
}

//提供一个接口，让外部去实现
func test(x, y int, callback Callback) int {
	return callback(x, y)
}

func add(x, y int) int {
	return x + y
}

```

运行结果

```
3
```
