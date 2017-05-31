# 变量和常量定义
现在我们讨论一下Go语言的变量定义。
 
**变量定义** 

所谓的变量就是一个拥有指定`名称`和`类型`的`数据存储位置`。  
在上面我们使用过变量的定义，现在我们来仔细看一个例子。  
  
	package main

	import (
		"fmt"
	)

	func main() {
		var x string = "hello world"
		fmt.Println(x)
	}

变量的定义首先使用`var`关键字，然后指定变量的名称`x`，再指定变量的类型`string`，在本例中，还对变量`x`进行了赋值，然后在命令行输出该变量。Go这种变量定义的方式和其他的语言有些不同，但是在使用的过程中，你会逐渐喜欢的。当然上面的变量定义方式还可以如下，即先定义变量，再赋值。

	package main

	import (
		"fmt"
	)

	func main() {
		var x string
		x = "hello world"
		fmt.Println(x)
	}

或者是直接赋值，让Go语言推断变量的类型。如下：

	package main

	import (
		"fmt"
	)

	func main() {
		var x = "hello world"
		fmt.Println(x)
	}

当然，上面变量的定义还有一种`快捷方式`。如果你知道变量的初始值，完全可以像下面这样定义变量，完全让`Go来推断语言的类型`。这种定义的方式连关键字`var`都省略掉了。

	package main

	import (
		"fmt"
	)

	func main() {
		x := "hello world"
		fmt.Println(x)
	}
注意：上面这种使用`:=`方式定义变量的方式`只能用在函数内部`。

	package main

	import (
		"fmt"
	)

	x:="hello world"
	func main() {
		y := 10
		fmt.Println(x)
		fmt.Println(y)
	}
	
对于上面的变量定义x是无效的。会导致编译错误：


	./test_var_quick.go:7: non-declaration statement outside function body


不过我们对上面的例子做下修改，比如这样是可以的。也就是使用var关键字定义的时候，如果给出初始值，就不需要显式指定变量类型。

	package main

	import (
		"fmt"
	)

	var x = "hello world"

	func main() {
		y := 10
		fmt.Println(x)
		fmt.Println(y)
	}


`变量`之所以称为变量，就是因为`它们的值在程序运行过程中可以发生变化`，但是`它们的变量类型是无法改变的`。因为`Go语言是静态语言`，并`不支持`程序运行过程中`变量类型发生变化`。比如如果你强行将一个字符串值赋值给定义为int的变量，那么会发生编译错误。即使是强制类型转换也是不可以的。`强制类型转换只支持同类的变量类型`。比如数值类型之间强制转换。

下面我们看几个例子：

	package main

	import (
		"fmt"
	)

	func main() {
		var x string = "hello world"
		fmt.Println(x)
		x = "i love go language"
		fmt.Println(x)
	}

本例子演示变量的值在程序运行过程中发生变化，结果输出为

	hello world
	i love go language

我们尝试不同类型的变量之间转换

	package main

	import (
		"fmt"
	)

	func main() {
		var x string = "hello world"
		fmt.Println(x)
		x = 11
		fmt.Println(x)
	}

在本例子中，如果试图将一个数值赋予字符串变量x，那么会发生错误：

	./test_var.go:10: cannot use 11 (type int) as type string in assignment

上面的意思就是无法将整型数值11当作字符串赋予给字符串变量。

但是同类的变量之间是可以强制转换的，如浮点型和整型之间的转换。

	package main

	import (
		"fmt"
	)

	func main() {
		var x float64 = 32.35
		fmt.Println(x)
		fmt.Println(int(x))
	}
输出的结果为

	32.35
	32


**变量命名**

上面我们看了一些变量的使用方法，那么定义一个变量名称，有哪些要求呢？
这里我们要注意，`Go的变量名称必须以字母或下划线(_)开头，后面可以跟字母，数字，或者下划线(_)`。除此之外，Go语言并不关心你如何定义变量。我们通用的做法是定义一个用户友好的变量。假设你需要定义一个狗狗的年龄，那么使用dog_age作为变量名称要好于用x来定义变量。

**变量作用域**

现在我们再来讨论一下变量的作用域。所谓作用域就是可以有效访问变量的区域。比如很简单的，你不可能在一个函数func_a里面访问另一个函数func_b里面定义的局部变量x。所以变量的作用域目前分为两类，一个是`全局变量`，另一个是`局部变量`。下面我们看个全局变量的例子：

	package main

	import (
		"fmt"
	)

	var x string = "hello world"

	func main() {
		fmt.Println(x)
	}
这里变量x定义在main函数之外，但是main函数仍然可以访问x。全局变量的作用域是该包中所有的函数。

	package main

	import (
		"fmt"
	)

	var x string = "hello world"

	func change() {
		x = "i love go"
	}
	func main() {
		fmt.Println(x)
		change()
		fmt.Println(x)
	}
在上面的例子用，我们用了change函数改变了x的值。输出结果如下：

	hello world
	i love go


我们再看一下局部变量的例子。

	package main

	import (
		"fmt"
	)

	func change() {
		x := "i love go"
	}
	func main() {
		fmt.Println(x)
	}

该例子中main函数试图访问change函数中定义的局部变量x，结果发生了下面的错误(未定义的变量x)：

	./test_var.go:11: undefined: x



**常量**

Go语言也支持常量定义。所谓`常量就是在程序运行过程中保持值不变的变量定义`。常量的定义和变量类似，只是用`const`关键字替换了var关键字，另外常量在定义的时候`必须有初始值`。

	package main

	import (
		"fmt"
	)

	func main() {
		const x string = "hello world"
		const y = "hello world"
		fmt.Println(x)
		fmt.Println(y)
	}

这里有一点需要注意，变量定义的类型推断方式`:=`不能够用来定义常量。因为常量的值是在编译的时候就已经确定的，但是变量的值则是运行的时候才使用的。这样常量定义就无法使用变量类型推断的方式了。

常量的值在运行过程中是无法改变的，强制改变常量的值是无效的。


	package main

	import (
		"fmt"
	)

	func main() {
		const x string = "hello world"
		fmt.Println(x)
		x = "i love go language"
		fmt.Println(x)
	}
比如上面的例子就会报错

	./test_var.go:10: cannot assign to x

我们再看一个Go包math里面定义的常量Pi，用它来求圆的面积。

	package main

	import (
		"fmt"
		"math"
	)

	func main() {
		var radius float64 = 10
		var area = math.Pow(radius, 2) * math.Pi
		fmt.Println(area)
	}


**多变量或常量定义**

Go还提供了一种`同时定义多个变量或者常量`的快捷方式。

	package main

	import (
		"fmt"
	)

	func main() {
		var (
			a int     = 10
			b float64 = 32.45
			c bool    = true
		)
		const (
			Pi   float64 = 3.14
			True bool    = true
		)

		fmt.Println(a, b, c)
		fmt.Println(Pi, True)
	}
