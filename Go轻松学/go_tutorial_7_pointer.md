# Go指针
不要害怕，Go的指针是好指针。

**定义**

所谓`指针其实你可以把它想像成一个箭头，这个箭头指向（存储）一个变量的地址`。

因为这个箭头本身也需要变量来存储，所以也叫做指针变量。

Go的指针`不支持那些乱七八糟的指针移位`。`它就表示一个变量的地址`。看看这个例子：

	package main

	import (
		"fmt"
	)

	func main() {
		var x int
		var x_ptr *int

		x = 10
		x_ptr = &x

		fmt.Println(x)
		fmt.Println(x_ptr)
		fmt.Println(*x_ptr)
	}


上面例子输出`x的值`，`x的地址`和`通过指针变量输出x的值`，而`x_ptr就是一个指针变量`。

	10
	0xc084000038
	10

认真理清楚这两个符号的意思。

**&** `取一个变量的地址`

**\*** `取一个指针变量所指向的地址的值`


考你一下，上面的例子中，如何输出x_ptr的地址呢？

	package main

	import (
		"fmt"
	)

	func main() {
		var x int
		var x_ptr *int

		x = 10
		x_ptr = &x

		fmt.Println(&x_ptr)
	}

此例看懂，指针就懂了。

永远记住一句话，`所谓指针就是一个指向（存储）特定变量地址的变量`。没有其他的特别之处。

再变态一下，看看这个：

	package main

	import (
		"fmt"
	)

	func main() {
		var x int
		var x_ptr *int

		x = 10
		x_ptr = &x

		fmt.Println(*&x_ptr)
	}

1. x_ptr 是一个`指针变量`，它`指向(存储)x的地址`；
2. &x_ptr 是`取这个指针变量x_ptr的地址`，这里可以设想`有另一个指针变量x_ptr_ptr(指向)存储`这个`x_ptr指针的地址`；
3. *&x_ptr 等价于`*x_ptr_ptr`就是`取这个x_ptr_ptr指针变量`所`指向(存储)`的`地址所对应的变量的值` ，也就是`x_ptr的值`，也就是`指针变量x_ptr指向(存储)的地址`，也就是`x的地址`。 这里可以看到，其实`*&`这两个运算符在一起就相互抵消作用了。

**用途**

`指针的一大用途就是可以将变量的指针作为实参传递给函数，从而在函数内部能够直接修改实参所指向的变量值。`

Go的变量传递都是值传递。

	package main

	import (
		"fmt"
	)

	func change(x int) {
		x = 200
	}
	func main() {
		var x int = 100
		fmt.Println(x)
		change(x)
		fmt.Println(x)
	}


上面的例子输出结果为

	100
	100

很显然，change函数`改变的`仅仅是`内部变量x`的`值`，而`不会改变`传递进去的`实参`。其实，也就是说Go的函数一般关心的是输出结果，而输入参数就相当于信使跑到函数门口大叫，你们这个参数是什么值，那个是什么值，然后就跑了。你函数根本就不能修改它的值。不过如果是传递的实参是指针变量，那么函数一看，小子这次你地址我都知道了，哪里跑。那么就是下面的例子：

	package main

	import (
		"fmt"
	)

	func change(x *int) {
		*x = 200
	}
	func main() {
		var x int = 100
		fmt.Println(x)
		change(&x)
		fmt.Println(x)
	}


上面的例子中，change函数的虚参为`整型指针变量`，所以在main中调用的时候`传递的是x的地址`。然后在change里面使用`*x=200`修改了这个x的地址的值。所以`x的值就变了`。这个输出是：

	100
	200


**new**

new这个函数挺神奇，因为它的用处太多了。这里还可以通过new来`初始化一个指针`。上面说过指针指向(存储)的是一个变量的地址，但是指针本身也需要地址存储。先看个例子：

	package main

	import (
		"fmt"
	)

	func set_value(x_ptr *int) {
		*x_ptr = 100
	}
	func main() {
		x_ptr := new(int)
		set_value(x_ptr)
		//x_ptr指向的地址
		fmt.Println(x_ptr)
		//x_ptr本身的地址
		fmt.Println(&x_ptr)
		//x_ptr指向的地址值
		fmt.Println(*x_ptr)
	}


上面我们定义了一个x_ptr变量，然后用`new申请`了一个`存储整型数据的内存地址`，然后将这个`地址赋值`给`x_ptr指针变量`，也就是说`x_ptr指向（存储）的是一个可以存储整型数据的地址`，然后用set_value函数将`这个地址中存储的值`赋值为100。所以第一个输出是`x_ptr指向的地址`，第二个则是`x_ptr本身的地址`，而`*x_ptr`则是`x_ptr指向的地址中存储的整型数据的值`。


	0xc084000040
	0xc084000038
	100

**小结**

好了，现在用个例子再来回顾一下指针。

交换两个变量的值。

	package main

	import (
		"fmt"
	)

	func swap(x, y *int) {
		*x, *y = *y, *x
	}
	func main() {
		x_val := 100
		y_val := 200
		swap(&x_val, &y_val)
		fmt.Println(x_val)
		fmt.Println(y_val)
	}


很简单吧，这里利用了Go提供的`交叉赋值`的功能，另外由于是使用了指针作为参数，所以在swap函数内，x_val和y_val的值就被交换了。


