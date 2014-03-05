# Go Exit
使用`os.Exit`可以给定一个状态，然后立刻退出程序运行。
```go
package main

import "fmt"
import "os"

func main() {
	// 当使用`os.Exit`的时候defer操作不会被运行，
	// 所以这里的``fmt.Println`将不会被调用
	defer fmt.Println("!")

	// 退出程序并设置退出状态值
	os.Exit(3)
}
```
注意，Go和C语言不同，main函数并不返回一个整数来表示程序的退出状态，而是将退出状态作为`os.Exit`函数的参数。

如果你使用`go run`来运行程序，将会有如下输出
```
exit status 3
```
如果你使用`go build`先编译程序，然后再运行可执行文件，程序将不会有输出。
如果你想查看程序的返回值，*nix系列系统下面使用如下方法:
```
$ ./go_exit
$ echo $?
3
```