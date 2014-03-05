# Go 命令行参数标记
命令行参数标记是为命令行程序指定选项参数的常用方法。例如，在命令`wc -l`中，`-l`就是一个命令行参数标记。

Go提供了`flag`包来支持基本的命令行标记解析。我们这里将要使用这个包提供的方法来实现带选项的命令行程序。

```go
package main

import "flag"
import "fmt"

func main() {

	// 基础的标记声明适用于string，integer和bool型选项。
	// 这里我们定义了一个标记`word`，默认值为`foo`和一
	// 个简短的描述。`flag.String`函数返回一个字符串指
	// 针（而不是一个字符串值），我们下面将演示如何使
	// 用这个指针
	wordPtr := flag.String("word", "foo", "a string")

	// 这里定义了两个标记，一个`numb`，另一个是`fork`，
	// 使用和上面定义`word`标记相似的方法
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// 你也可以程序中任意地方定义的变量来定义选项，只
	// 需要把该变量的地址传递给flag声明函数即可
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 当所有的flag声明完成后，使用`flag.Parse()`来分
	// 解命令行选项
	flag.Parse()

	// 这里我们仅仅输出解析后的选项和任何紧跟着的位置
	// 参数，注意我们需要使用`*wordPtr`的方式来获取最
	// 后的选项值
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
```
为了运行示例，你需要先将程序编译为可执行文件。
```
go build command-line-flags.go
```
下面分别看看给予该命令行程序不同选项参数的例子：

(1) 给所有的选项设置一个参数

```
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []
```
(2) 如果你不设置flag，那么它们自动采用默认的值
```
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []
```
(3) 尾部的位置参数可以出现在任意一个flag后面
```
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3]
```
(4) 注意flag包要求所有的flag都必须出现在尾部位置参数的前面，否则这些flag将被当作位置参数处理
```
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
trailing: [a1 a2 a3 -numb=7]
```
(5) 使用`-h`或者`--help`这两个flag来自动地生成命令行程序的帮助信息
```
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string
```
(6) 如果你提供了一个程序不支持的flag，那么程序会打印一个错误信息和帮助信息
```
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./go_cmd_flag:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string
```