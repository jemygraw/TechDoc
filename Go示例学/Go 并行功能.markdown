# Go 并行功能
goroutine是一个轻量级的线程。
```go
package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 假设我们有一个函数叫做f(s)
	// 这里我们使用通常的同步调用来调用函数
	f("direct")

	// 为了能够让这个函数以协程(goroutine)方式
	// 运行使用go f(s)
	// 这个协程将和调用它的协程并行执行
	go f("goroutine")

	// 你也可以为匿名函数开启一个协程运行
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 上面的协程在调用之后就异步执行了，所以程序不用等待它们执行完成
	// 就跳到这里来了，下面的Scanln用来从命令行获取一个输入，然后才
	// 让main函数结束
	// 如果没有下面的Scanln语句，程序到这里会直接退出，而上面的协程还
	// 没有来得及执行完，你将无法看到上面两个协程运行的结果
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
```
运行结果
```
direct : 0
direct : 1
direct : 2
goroutine : 0
goroutine : 1
goroutine : 2
going
ok
done
```