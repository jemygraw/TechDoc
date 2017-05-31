# 使用包和测试管理项目

Go天生就是为了支持良好的项目管理体验而设计的。

**包**

在软件工程的实践中，我们会遇到很多功能重复的代码，比如去除字符串首尾的空格。高质量软件产品的特点就是它的部分代码是可以重用的，比如你不必每次写个函数去去除字符串首尾的空格。

我们上面讲过变量，结构体，接口和函数等，事实上所谓的包，就是把一些用的多的这些变量，结构体，接口和函数等统一放置在一个逻辑块中。并且给它们起一个名字，这个名字就叫做包名。

例如我们上面用的最多的fmt包，这个包提供了很多格式化输出的函数，你可以在自己的代码中引用这个包，来做格式化输出，而不用你自己每次去写个函数。一门成熟的语言都会提供齐全的基础功能包供人调用。

使用包有三个好处

1. 可以减少函数名称重复，因为不同包中可以存在名称相同的函数。否则得话，你得给这些函数加上前缀或者后缀以示区别。
2. 包把函数等组织在一起，方便你查找和重用。比如你想用Println()函数输出一行字符串，你可以很方便地知道它在fmt包中，直接引用过来用就可以了。
3. 使用包可以加速程序编译。因为包是预编译好的，你改动自己代码得时候，不必每次去把包编译一下。

**创建包**

我们现在来举个例子，用来演示Go的项目管理。

首先我们在目录`/Users/jemy/JemyGraw/GoLang`下面创建文件夹`pkg_demo`。然后在`pkg_demo`里面创建`src`文件夹
。然后再在`main`文件夹里面创建`main.go`文件。另外为了演示包的创建，我们在`src`目录下面创建文件夹`net.duokr`，然后再在`net.duokr`文件夹里面创建`math`文件夹，这个文件夹名称就是这个文件夹下面go文件的包名称。然后我们再创建一个`math_pkg.go`文件，之所以取这个名字而不是`math.go`只是为了说明这个文件名称和包名不需要一致。然后我们还创建了一个`math_pkg_test.go`文件作为包的测试用例文件。整体结构如下：

	.
	└── src
	    ├── main
	    │   ├── build.sh
	    │   └── main.go
	    └── net.duokr
	        └── math
	            ├── math_pkg.go
	            └── math_pkg_test.go


其中build.sh是我们为了编译这个项目而写的脚本，因为编译项目需要几条命令，把它写在脚本文件中方便使用。另外为了能够让build.sh能够执行，使用`chmod +x build.sh`为它赋予可执行权限。build.bat是Windows下面的编译脚本。
我们来看一下`math_pkg.go`的定义：

	package math

	func Add(a, b int) int {
		return a + b
	}
	func Subtract(a, b int) int {
		return a - b
	}
	func Multiply(a, b int) int {
		return a * b
	}

	func Divide(a, b int) int {
		if b == 0 {
			panic("Can not divided by zero")
		}
		return a / b
	}
		
首先是包名，然后是几个函数定义，这里我们会发现这些`函数定义首字母都是大写`，`Go规定了只有首字母大写的函数才能从包导出使用，即其他调用这个包中函数的代码只能调用那些导出的函数`。

我们再看一下`main.go`的定义：

	package main

	import (
		"fmt"
		math "net.duokr/math"
	)

	func main() {
		var a = 100
		var b = 200

		fmt.Println("Add demo:", math.Add(a, b))
		fmt.Println("Substract demo:", math.Subtract(a, b))
		fmt.Println("Multiply demo:", math.Multiply(a, b))
		fmt.Println("Divide demo:", math.Divide(a, b))
	}

在main.go里面，我们使用import关键字引用我们自定义的包math，引用的方法是从main包平行的文件夹net.duokr开始，后面跟上包名math。这里面我们给这个长长的包名起了一个别名就叫math。然后分别调用math包里面的函数。

最后我们看一下我们的编译脚本：

	export GOPATH=$GOPATH:/Users/jemy/JemyGraw/GoLang/pkg_demo
	export GOBIN=/Users/jemy/JemyGraw/GoLang/pkg_demo/bin
	go build net.duokr/math
	go build main.go
	go install main

第一行，我们将项目路径加入GOPATH中，这样待会儿编译main.go的时候才能找到我们自定义的包；

第二行，我们设置本项目的安装目录，第五行的命令将编译好的文件放到这个目录下面；

第三行，我们编译我们的自定义包；

第四行，我们编译我们main.go文件；

第五行，将编译好的文件安装到指定目录下。

这里还有一个Windows下面的编译脚本build.bat：

	@echo off
	set GOPATH=GOPATH;C:\JemyGraw\GoLang\pkg_demo
	set GOBIN=C:\JemyGraw\GoLang\pkg_demo\bin
	go build net.duokr\math
	go build main.go
	go install main

好了，运行脚本编译一下，在main文件夹和bin文件夹下面都会生成一个可执行文件。

这个时候文件夹结构为：

	.
	├── bin
	│   └── main
	├── pkg
	│   └── darwin_386
	│       └── net.duokr
	│           └── math.a
	└── src
	    ├── main
	    │   ├── build.bat
	    │   ├── build.sh
	    │   ├── main
	    │   └── main.go
	    └── net.duokr
	        └── math
	            ├── math_pkg.go
	            └── math_pkg_test.go

运行一下，输出结果为：

	Add demo: 300
	Substract demo: -100
	Multiply demo: 20000
	Divide demo: 0
	
好了，包的使用介绍完毕，我们再来看一下测试用例怎么写。

**测试**

在上面的例子中，我们发现我们自定义的包下面还有一个math_pkg_test.go文件，这个文件包含了本包的一些测试用例。而且Go会把以`_test.go`结尾的文件当作是测试文件。

测试怎么写，当然是用assert来判断程序的运行结果是否和预期的相同了。

我们来看看这个math包的测试用例。

	package math

	import (
		"testing"
	)

	func TestAdd(t *testing.T) {
		var a = 100
		var b = 200

		var val = Add(a, b)
		if val != a+b {
			t.Error("Test Case [", "TestAdd", "] Failed!")
		}
	}

	func TestSubtract(t *testing.T) {
		var a = 100
		var b = 200

		var val = Subtract(a, b)
		if val != a-b {
			t.Error("Test Case [", "TestSubtract", "] Failed!")
		}
	}

	func TestMultiply(t *testing.T) {
		var a = 100
		var b = 200

		var val = Multiply(a, b)
		if val != a*b {
			t.Error("Test Case [", "TestMultiply", "] Failed!")
		}
	}

	func TestDivideNormal(t *testing.T) {
		var a = 100
		var b = 200

		var val = Divide(a, b)
		if val != a/b {
			t.Error("Test Case [", "TestDivideNormal", "] Failed!")
		}
	}

将路径切换到测试文件所在目录，运行`go test`命令，go会自动测试所有的测试用例。

在上面的例子中，测试用例的特点是以函数名以`Test`开始，而且具有唯一参数`t *testing.T`。

**小结**

Go提供的包和用例测试是构建优秀的软件产品的基础，只要我们不断学习，努力去做，一定可以做的很好。
	

