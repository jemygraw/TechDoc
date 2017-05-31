# Go函数
是时候讨论一下Go的函数定义了。

**什么是函数**

函数，简单来讲就是一段将`输入数据`转换为`输出数据`的`公用代码块`。当然有的时候函数的返回值为空，那么就是说输出数据为空。而真正的处理过程在函数内部已经完成了。

想一想我们为什么需要函数，最直接的需求就是代码中有太多的重复代码了，为了代码的可读性和可维护性，将这些重复代码重构为函数也是必要的。

**函数定义**

先看一个例子

	package main

	import (
		"fmt"
	)

	func slice_sum(arr []int) int {
		sum := 0
		for _, elem := range arr {
			sum += elem
		}
		return sum
	}

	func main() {
		var arr1 = []int{1, 3, 2, 3, 2}
		var arr2 = []int{3, 2, 3, 1, 6, 4, 8, 9}
		fmt.Println(slice_sum(arr1))
		fmt.Println(slice_sum(arr2))
	}

在上面的例子中，我们需要分别计算两个切片的元素和。如果我们把计算切片元素的和的代码分别为两个切片展开，那么代码就失去了简洁性和一致性。假设你预想实现同样功能的代码在拷贝粘贴的过程中发生了错误，比如忘记改变量名之类的，到时候debug到崩溃吧。因为这时很有可能你就先入为主了，因为模板代码没有错啊，是不是。所以函数就是这个用处。

我们再仔细看一下上面的函数定义：

首先是关键字`func`，然后后面是`函数名称`，`参数列表`，最后是`返回值列表`。当然如果函数没有参数列表或者返回值，那么这两项都是可选的。其中返回值两边的括号在只声明一个返回值类型的时候可以省略。

**命名返回值**

Go的函数很有趣，你甚至可以为返回值预先定义一个名称，在函数结束的时候，直接一个return就可以返回所有的预定义返回值。例如上面的例子，我们将sum作为命名返回值。

	package main

	import (
		"fmt"
	)

	func slice_sum(arr []int) (sum int) {
		sum = 0
		for _, elem := range arr {
			sum += elem
		}
		return
	}

	func main() {
		var arr1 = []int{1, 3, 2, 3, 2}
		var arr2 = []int{3, 2, 3, 1, 6, 4, 8, 9}
		fmt.Println(slice_sum(arr1))
		fmt.Println(slice_sum(arr2))
	}

这里要注意的是，如果你定义了命名返回值，那么在函数内部你将不能再重复定义一个同样名称的变量。比如第一个例子中我们用`sum:=0`来定义和初始化变量sum，而在第二个例子中，我们只能用`sum=0`初始化这个变量了。因为`:=`表示的是定义并且初始化变量。

**实参数和虚参数**

可能你听说过函数的实参数和虚参数。其实所谓的`实参数就是函数调用的时候传入的参数`。在上面的例子中，实参就是`arr1`和`arr2`，而`虚参数就是函数定义的时候表示函数需要传入哪些参数的占位参数`。在上面的例子中，虚参就是`arr`。`实参和虚参的名字不必是一样的。即使是一样的，也互不影响。`因为虚参是函数的内部变量。而实参则是另一个函数的内部变量或者是全局变量。它们的作用域不同。如果一个函数的虚参碰巧和一个全局变量名称相同，那么函数使用的也是虚参。例如我们再修改一下上面的例子。

	package main

	import (
		"fmt"
	)

	var arr = []int{1, 3, 2, 3, 2}

	func slice_sum(arr []int) (sum int) {
		sum = 0
		for _, elem := range arr {
			sum += elem
		}
		return
	}

	func main() {
		var arr2 = []int{3, 2, 3, 1, 6, 4, 8, 9}
		fmt.Println(slice_sum(arr))
		fmt.Println(slice_sum(arr2))
	}


在上面的例子中，我们定义了全局变量arr并且初始化值，而我们的slice_sum函数的虚参也是arr，但是程序同样正常工作。

**函数多返回值**

记不记得你在java或者c里面需要返回多个值时还得去定义一个对象或者结构体的呢？在Go里面，你不需要这么做了。Go函数支持你返回多个值。

其实函数的多返回值，我们在上面遇见过很多次了。那就是`range`函数。这个函数用来迭代数组或者切片的时候返回的是两个值，一个是数组或切片元素的索引，另外一个是数组或切片元素。在上面的例子中，因为我们不需要元素的索引，所以我们用一个特殊的忽略返回值符号`下划线(_)`来忽略索引。

假设上面的例子我们除了返回切片的元素和，还想返回切片元素的平均值，那么我们修改一下代码。

	package main

	import (
		"fmt"
	)

	func slice_sum(arr []int) (int, float64) {
		sum := 0
		avg := 0.0
		for _, elem := range arr {
			sum += elem
		}
		avg = float64(sum) / float64(len(arr))
		return sum, avg
	}

	func main() {
		var arr1 = []int{3, 2, 3, 1, 6, 4, 8, 9}
		fmt.Println(slice_sum(arr1))
	}


很简单吧，当然我们还可以将上面的参数定义为命名参数

	package main

	import (
		"fmt"
	)

	func slice_sum(arr []int) (sum int, avg float64) {
		sum = 0
		avg = 0.0
		for _, elem := range arr {
			sum += elem
		}
		avg = float64(sum) / float64(len(arr))
		//return sum, avg
		return
	}

	func main() {
		var arr1 = []int{3, 2, 3, 1, 6, 4, 8, 9}
		fmt.Println(slice_sum(arr1))
	}


在上面的代码里面，将`return sum, avg`给注释了而直接使用`return`。其实这两种返回方式都可以。

**变长参数**

想一想我们的fmt包里面的Println函数，它怎么知道你传入的参数个数呢？

	package main

	import (
		"fmt"
	)

	func main() {
		fmt.Println(1)
		fmt.Println(1, 2)
		fmt.Println(1, 2, 3)
	}

这个要归功于Go的一大特性，支持可变长参数列表。

首先我们来看一个例子

	package main

	import (
		"fmt"
	)

	func sum(arr ...int) int {
		sum := 0
		for _, val := range arr {
			sum += val
		}
		return sum
	}
	func main() {
		fmt.Println(sum(1))
		fmt.Println(sum(1, 2))
		fmt.Println(sum(1, 2, 3))
	}


在上面的例子中，我们将原来的切片参数修改为可变长参数，然后使用range函数迭代这些参数，并求和。
从这里我们可以看出至少一点那就是`可变长参数列表里面的参数类型都是相同的`（*如果你对这句话表示怀疑，可能是因为你看到Println函数恰恰可以输出不同类型的可变参数，这个问题的答案要等到我们介绍完Go的接口后才行*）。

另外还有一点需要注意，那就是`可变长参数定义只能是函数的最后一个参数`。比如下面的例子：

	package main

	import (
		"fmt"
	)

	func sum(base int, arr ...int) int {
		sum := base
		for _, val := range arr {
			sum += val
		}
		return sum
	}
	func main() {
		fmt.Println(sum(100, 1))
		fmt.Println(sum(200, 1, 2))
		fmt.Println(sum(300, 1, 2, 3))
	}


这里不知道你是否觉得这个例子其实和那个切片的例子很像啊，在哪里呢？

	package main

	import (
		"fmt"
	)

	func sum(base int, arr ...int) int {
		sum := base
		for _, val := range arr {
			sum += val
		}
		return sum
	}
	func main() {
		var arr1 = []int{1, 2, 3, 4, 5}
		fmt.Println(sum(300, arr1...))
	}

呵呵，就是把切片“啪，啪，啪”三个耳光打碎了，传递过去啊！:-P


**闭包函数**

曾经使用python和javascript的时候就在想，如果有一天可以把这两种语言的特性做个并集该有多好。

这一天终于来了，Go支持闭包函数。

首先看一个闭包函数的例子。所谓闭包函数就是将整个函数的定义一气呵成写好并赋值给一个变量。然后用这个变量名作为函数名去调用函数体。

我们将刚刚的例子修改一下：

	package main

	import (
		"fmt"
	)

	func main() {
		var arr1 = []int{1, 2, 3, 4, 5}
		
		var sum = func(arr ...int) int {
			total_sum := 0
			for _, val := range arr {
				total_sum += val
			}
			return total_sum
		}
		fmt.Println(sum(arr1...))
	}

从这里我们可以看出，其实闭包函数也没有什么特别之处。因为Go不支持在一个函数的内部再定义一个嵌套函数，所以使用闭包函数能够实现在一个函数内部定义另一个函数的目的。

这里我们需要注意的一个问题是，闭包函数对它外层的函数中的变量具有`访问`和`修改`的权限。例如：

	package main

	import (
		"fmt"
	)

	func main() {
		var arr1 = []int{1, 2, 3, 4, 5}
		var base = 300
		var sum = func(arr ...int) int {
			total_sum := 0
			total_sum += base
			for _, val := range arr {
				total_sum += val
			}
			return total_sum
		}
		fmt.Println(sum(arr1...))
	}


这个例子，输出315，因为total_sum加上了base的值。

	package main

	import (
		"fmt"
	)

	func main() {
		var base = 0
		inc := func() {
			base += 1
		}
		fmt.Println(base)
		inc()
		fmt.Println(base)
	}


在上面的例子中，闭包函数修改了main函数的局部变量base。

最后我们来看一个闭包的示例，生成偶数序列。

	package main

	import (
		"fmt"
	)

	func createEvenGenerator() func() uint {
		i := uint(0)
		return func() (retVal uint) {
			retVal = i
			i += 2
			return
		}
	}
	func main() {
		nextEven := createEvenGenerator()
		fmt.Println(nextEven())
		fmt.Println(nextEven())
		fmt.Println(nextEven())
	}


这个例子很有意思的，因为我们定义了一个`返回函数定义`的函数。而所返回的函数定义就是`在这个函数的内部定义的闭包函数`。这个闭包函数在外层函数调用的时候，每次都生成一个新的偶数（加2操作）然后返回闭包函数定义。

其中`func() uint`就是函数createEvenGenerator的返回值。在createEvenGenerator中，这个返回值是return返回的闭包函数定义。

	func() (retVal uint) {
        	retVal = i
        	i += 2
        	return
    	}

因为createEvenGenerator函数返回的是一个函数定义，所以我们再把它赋值给一个代表函数的变量，然后用这个代表闭包函数的变量去调用函数执行。

**递归函数**

每次谈到递归函数，必然绕不开阶乘和斐波拉切数列。

阶乘

	package main

	/**
	    n!=1*2*3*...*n
	*/
	import (
		"fmt"
	)

	func factorial(x uint) uint {
		if x == 0 {
			return 1
		}
		return x * factorial(x-1)
	}

	func main() {
		fmt.Println(factorial(5))
	}


如果x为0，那么返回1，因为0!=1。如果x是1，那么f(1)=1*f(0)，如果x是2，那么f(2)=2*f(1)=2*1*f(0)，依次推断f(x)=x*(x-1)*...*2*1*f(0)。

从上面看出所谓递归，就是在函数的内部重复调用一个函数的过程。需要注意的是这个函数必须能够一层一层分解，并且有出口。上面的例子出口就是0。

斐波拉切数列

求第N个斐波拉切元素

	package main

	/**
		f(1)=1
		f(2)=2
		f(n)=f(n-2)+f(n-1)
	*/
	import (
		"fmt"
	)

	func fibonacci(n int) int {
		var retVal = 0
		if n == 1 {
			retVal = 1
		} else if n == 2 {
			retVal = 2
		} else {
			retVal = fibonacci(n-2) + fibonacci(n-1)
		}
		return retVal

	}
	func main() {
		fmt.Println(fibonacci(5))
	}


斐波拉切第一个元素是1，第二个元素是2，后面的元素依次是前两个元素的和。

其实对于递归函数来讲，只要知道了函数的出口，后面的不过是让计算机去不断地推断，一直推断到这个出口。理解了这一点，递归就很好理解了。


**异常处理**

当你读取文件失败而退出的时候是否担心文件句柄是否已经关闭？抑或是你对于try...catch...finally的结构中finally里面的代码和try里面的return代码那个先执行这样的问题痛苦不已？

一切都结束了。一门完美的语言必须有一个清晰的无歧义的执行逻辑。

好，来看看Go提供的异常处理。

*defer*

	package main

	import (
		"fmt"
	)

	func first() {
		fmt.Println("first func run")
	}
	func second() {
		fmt.Println("second func run")
	}

	func main() {
		defer second()
		first()
	}

Go语言提供了关键字`defer`来在函数运行结束的时候运行一段代码或调用一个清理函数。上面的例子中，虽然second()函数写在first()函数前面，但是由于使用了defer标注，所以它是在main函数执行结束的时候才调用的。

所以输出结果

	first func run
	second func run

`defer`用途最多的在于释放各种资源。比如我们读取一个文件，读完之后需要释放文件句柄。

	package main

	import (
		"bufio"
		"fmt"
		"os"
		"strings"
	)

	func main() {
		fname := "D:\\Temp\\test.txt"
		f, err := os.Open(fname)
		defer f.Close()
		if err != nil {
			os.Exit(1)
		}
		bReader := bufio.NewReader(f)
		for {
			line, ok := bReader.ReadString('\n')
			if ok != nil {
				break
			}
			fmt.Println(strings.Trim(line, "\r\n"))
		}
	}


在上面的例子中，我们按行读取文件，并且输出。从代码中，我们可以看到在使用os包中的Open方法打开文件后，立马跟着一个defer语句用来关闭文件句柄。这样就保证了该文件句柄在main函数运行结束的时候或者异常终止的时候一定能够被释放。而且由于紧跟着Open语句，一旦养成了习惯，就不会忘记去关闭文件句柄了。


*panic* & *recover*

>当你周末走在林荫道上，听着小歌，哼着小曲，很是惬意。突然之间，从天而降瓢泼大雨，你顿时慌张（panic）起来，没有带伞啊，淋着雨感冒就不好了。于是你四下张望，忽然发现自己离地铁站很近，那里有很多卖伞的，心中顿时又安定了下来（recover），于是你飞奔过去买了一把伞（defer）。

好了，panic和recover是Go语言提供的用以处理异常的关键字。`panic用来触发异常`，而`recover用来终止异常并且返回传递给panic的值`。（注意`recover并不能处理异常`，而且`recover只能在defer里面使用，否则无效`。）

先瞧个小例子

	package main

	import (
		"fmt"
	)

	func main() {
		fmt.Println("I am walking and singing...")
		panic("It starts to rain cats and dogs")
		msg := recover()
		fmt.Println(msg)
	}

看看输出结果


	runtime.panic(0x48d380, 0xc084003210)
        C:/Users/ADMINI~1/AppData/Local/Temp/2/bindist667667715/go/src/pkg/runtime/panic.c:266 	+0xc8
	main.main()
        D:/JemyGraw/Creation/Go/freebook_go/func_d1.go:9 +0xea
	exit status 2

咦？怎么没有输出recover获取的错误信息呢？

这是因为在运行到panic语句的时候，程序已经异常终止了，后面的代码就不运行了。

那么如何才能阻止程序异常终止呢？这个时候要使用defer。因为`defer一定是在函数执行结束的时候运行的。不管是正常结束还是异常终止`。

修改一下代码

	package main

	import (
		"fmt"
	)

	func main() {
		defer func() {
			msg := recover()
			fmt.Println(msg)
		}()
		fmt.Println("I am walking and singing...")
		panic("It starts to rain cats and dogs")
	}


好了，看下输出

	I am walking and singing...
	It starts to rain cats and dogs

小结：

panic触发的异常通常是运行时错误。比如试图访问的索引超出了数组边界，忘记初始化字典或者任何无法轻易恢复到正常执行的错误。





