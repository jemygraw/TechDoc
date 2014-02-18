# Go变量
Go是静态类型语言，变量是有明确类型的。编译器会检查函数调用中，变量类型的正确性。

使用`var`关键字来定义变量。

Go 的基本类型有：

- bool
- string
- int  int8  int16  int32  int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte // uint8 的别名
- rune // int32 的别名 代表一个Unicode码
- float32 float64
- complex64 complex128

看看下面的例子

```go
package main

import "fmt"

func main() {
	// `var` 关键字用来定义一个或者多个变量
	var a string = "initial"
	fmt.Println(a)

	// 你一次可以定义多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go会推断出具有初始值的变量的类型
	var d = true
	fmt.Println(d)

	//定义变量时，没有给出初始值的变量被默认初始化为零值
	//整型的零值就是0
	var e int
	fmt.Println(e)

	//":=" 语法是同时定义和初始化变量的快捷方式
	f := "short"
	fmt.Println(f)
}
```
输出结果为
```
initial
1 2
true
0
short
```