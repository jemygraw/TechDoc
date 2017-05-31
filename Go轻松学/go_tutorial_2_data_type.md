# Go语言内置基础数据类型

在自然界里面，有猫，有狗，有猪。有各种动物。每种动物都是不同的。  
比如猫会喵喵叫，狗会旺旺叫，猪会哼哼叫。。。  
Stop!!!  
好了，大家毕竟不是幼儿园的小朋友。介绍到这里就可以了。

论点就是每个东西都有自己归属的类别(Type)。  
那么在Go语言里面，每个变量也都是有类别的，这种类别叫做`数据类型(Data Type)`。  
Go的数据类型有两种：一种是`语言内置的数据类型`，另外一种是`通过语言提供的自定义数据类型方法自己定义的自定义数据类型`。

先看看语言`内置的基础数据类型`

**数值型(Number)**

数值型有`三种`，一种是`整数类型`，另外一种是`带小数的类型`(一般计算机里面叫做`浮点数类型`)，还有一种`虚数类型`。  

整数类型不用说了，和数学里面的是一样的。和数学里面不同的地方在于计算机里面`正整数和零`统称为`无符号整型`，而`负整数`则称为`有符号整型`。  

Go的内置整型有`uint8`, `uint16`, `uint32`, `uint64`, `int8`, `int16`, `int32`和`int64`。其中`u`开头的类型就是`无符号整型`。无符号类型能够表示正整数和零。而有符号类型除了能够表示正整数和零外，还可以表示负整数。
另外还有一些别名类型，比如`byte`类型，这个类型和`uint8`是一样的，表示`字节类型`。另外一个是`rune类型`，这个类型和`int32`是一样的，用来表示`unicode的代码点`，就是unicode字符所对应的整数。

Go还定义了三个`依赖系统`的类型，`uint`，`int`和`uintptr`。因为在32位系统和64位系统上用来表示这些类型的位数是不一样的。

*对于32位系统*

uint=uint32  
int=int32  
uintptr为32位的指针  

*对于64位系统*

uint=uint64  
int=int64  
uintptr为64位的指针  

至于类型后面跟的数字8，16，32或是64则表示用来表示这个类型的位不同，`位越多，能表示的整数范围越大`。
比如对于用N位来表示的整数，如果是`有符号的整数`，能够表示的整数范围为`-2^(N-1) ~ 2^(N-1)－1`；如果是`无符号的整数`，则能表示的整数范围为`0 ～ 2^N`。

Go的浮点数类型有两种，`float32`和`float64`。float32又叫`单精度浮点型`，float64又叫做`双精度浮点型`。其`最主要的区别就是小数点后面能跟的小数位数不同`。

另外Go还有两个其他语言所没有的类型，`虚数类型`。`complex64`和`complex128`。

对于数值类型，其所共有的操作为`加法(＋)`，`减法(－)`，`乘法(＊)`和`除法(/)`。另外对于`整数类型`，还定义了`求余运算(%)`

求余运算为整型所独有。如果对浮点数使用求余，比如这样

    package main

    import (
        "fmt"
    )

    func main() {
        var a float64 = 12
        var b float64 = 3

        fmt.Println(a % b)
    }


编译时候会报错

    invalid operation: a % b (operator % not defined on float64)

所以，这里我们可以知道所谓的`数据类型有两层意思`，一个是定义了`该类型所能表示的数`，另一个是定义了`该类型所能进行的操作`。
简单地说，对于一只小狗，你能想到的一定是狗的面貌和它会汪汪叫，而不是猫的面容和喵喵叫。


**字符串类型(String)**

字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由`单个字节`连接起来的。（对于汉字，通常由多个字节组成）。这就是说，传统的字符串是由字符组成的，而`Go的字符串不同`，是`由字节组成`的。这一点需要注意。

字符串的表示很简单。用(双引号"")或者(``号)来描述。

    "hello world"

或者

    `hello world`

唯一的区别是，**双引号之间的转义字符会被转义，而``号之间的转义字符保持原样不变**。

    package main

    import (
        "fmt"
    )

    func main() {
        var a = "hello \n world"
        var b = `hello \n world`

        fmt.Println(a)
        fmt.Println("----------")
        fmt.Println(b)
    }

输出结果为

    hello 
     world
    ----------
    hello \n world

字符串所能进行的一些基本操作包括:  
（1）`获取字符长度`  
（2）`获取字符串中单个字节`  
（3）`字符串连接`  

    package main

    import (
        "fmt"
    )

    func main() {
        var a string = "hello"
        var b string = "world"

        fmt.Println(len(a))
        fmt.Println(a[1])
        fmt.Println(a + b)
    }
    
输出如下 

    5
    101
    helloworld
  
这里我们看到a[1]得到的是一个整数，这就证明了上面`"Go的字符串是由字节组成的这句话"`。我们还可以再验证一下。

	package main

	import (
		"fmt"
	)

	func main() {
		var a string = "你"
		var b string = "好"
		fmt.Println(len(a))
		fmt.Println(len(b))
		fmt.Println(len(a + b))
		fmt.Println(a[0])
		fmt.Println(a[1])
		fmt.Println(a[2])
	}

输出

    3
    3
    6
    228
    189
    160

我们开始的时候，从上面的三行输出知道，"你"和"好"分别是用三个字节组成的。我们依次获取a的三个字节，输出，得到结果。


**布尔型(Bool)**

布尔型是表示`真值`和`假值`的类型。可选值为`true`和`false`。

所能进行的操作如下：
`&& and 与`
`|| or 或`
`!  not 非`

Go的布尔型取值`就是true`或`false`。`任何空值(nil)或者零值(0, 0.0, "")都不能作为布尔型来直接判断`。

	package main

	import (
    	"fmt"
	)

	func main() {
    	var equal bool
    	var a int = 10
    	var b int = 20
    	equal = (a == b)
    	fmt.Println(equal)
	}
输出结果

    false

下面是错误的用法

	package main

	import (
    	"fmt"
	)

	func main() {
    	if 0 {
        	fmt.Println("hello world")
    	}
    	if nil {
        	fmt.Println("hello world")
    	}
    	if "" {
     		fmt.Println("hello world")
    	}
	}

编译错误

    ./t.go:8: non-bool 0 (type untyped number) used as if condition
    ./t.go:11: non-bool nil used as if condition
    ./t.go:14: non-bool "" (type untyped string) used as if condition


上面介绍的是Go语言内置的基础数据类型。