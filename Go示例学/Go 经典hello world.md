我们的第一个例子是打印经典的“hello world”信息，我们先看下代码。

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```
输出结果为：
```
$ ls
el_01_hello_world.go
$ go build el_01_hello_world.go 
$ ls
el_01_hello_world	el_01_hello_world.go
$ ./el_01_hello_world 
hello world
```
为了使一个`go文件`能够`编译`为`可执行文件`，包名必须是`main`，然后我们导入提供格式化输出的`fmt`包，该程序的执行入口是`func main()`函数，在函数里面，我们使用`fmt`包提供的`Println`函数来输出"hello world"字符串。
	
为了运行这个程序，我们可以使用`go run el_01_hello_world.go`来运行这个例子，这样是直接输出运行结果而不会产生任何中间文件。但是有的时候我们希望能够将程序编译为二进制文件保存起来，我们可以像上面一样使用`go build el_01_hello_world.go`来将源代码编译为二进制可执行文件。然后我们可以直接运行这个二进制可执行文件。

好了，第一个例子就这样结束了。很简单。
