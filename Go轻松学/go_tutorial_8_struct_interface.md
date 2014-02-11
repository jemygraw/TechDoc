# Go结构体和指针

基本上到这里的时候，就是上了一个台阶了。Go的精华特点即将展开。

**结构体定义**

上面我们说过Go的指针和C的不同，结构体也是一样的。Go是一门删繁就简的语言，一切令人困惑的特性都必须去掉。

简单来讲，Go提供的`结构体`就是把`使用各种数据类型定义`的`不同变量组合起来`的`高级数据类型`。闲话不多说，看例子:

	type Rect struct {
		width float64
		length float64
	}

上面我们定义了一个矩形结构体，首先是关键是`type`表示要`定义一个新的数据类型了`，然后是新的数据类型名称`Rect`，最后是`struct`关键字，表示这个高级数据类型是结构体类型。在上面的例子中，因为`width和length的数据类型相同`，还可以写成如下格式：


	type Rect struct {
		width, length float64
	}


好了，来用结构体干点啥吧，计算一下矩形面积。

	package main

	import (
		"fmt"
	)

	type Rect struct {
		width, length float64
	}

	func main() {
		var rect Rect
		rect.width = 100
		rect.length = 200
		fmt.Println(rect.width * rect.length)
	}

从上面的例子看到，其实结构体类型和基础数据类型使用方式差不多，唯一的区别就是结构体类型可以通过`.`来访问内部的成员。包括`给内部成员赋值`和`读取内部成员值`。

在上面的例子中，我们是用var关键字先定义了一个Rect变量，然后对它的成员赋值。我们也可以使用初始化的方式来给Rect变量的内部成员赋值。

	package main

	import (
		"fmt"
	)

	type Rect struct {
		width, length float64
	}

	func main() {
		var rect = Rect{width: 100, length: 200}

		fmt.Println(rect.width * rect.length)
	}


当然`如果你知道结构体成员定义的顺序`，也可以不使用`key:value`的方式赋值，`直接按照结构体成员定义的顺序给它们赋值`。

	package main

	import (
		"fmt"
	)

	type Rect struct {
		width, length float64
	}

	func main() {
		var rect = Rect{100, 200}

		fmt.Println("Width:", rect.width, "* Length:",
			rect.length, "= Area:", rect.width*rect.length)
	}

输出结果为

	Width: 100 * Length: 200 = Area: 20000

**结构体参数传递方式**

我们说过，`Go函数的参数传递方式是值传递`，这句话`对结构体也是适用的`。


	package main

	import (
		"fmt"
	)

	type Rect struct {
		width, length float64
	}

	func double_area(rect Rect) float64 {
		rect.width *= 2
		rect.length *= 2
		return rect.width * rect.length
	}
	func main() {
		var rect = Rect{100, 200}
		fmt.Println(double_area(rect))
		fmt.Println("Width:", rect.width, "Length:", rect.length)
	}

上面的例子输出为:

	80000
	Width: 100 Length: 200

也就说虽然在double_area函数里面我们将结构体的宽度和长度都加倍，但仍然没有影响main函数里面的rect变量的宽度和长度。


**结构体组合函数**

上面我们在main函数中计算了矩形的面积，但是我们觉得矩形的面积如果能够作为矩形结构体的“内部函数”提供会更好。这样我们就可以直接说这个矩形面积是多少，而不用另外去取宽度和长度去计算。现在我们看看结构体“内部函数”定义方法：


	package main

	import (
		"fmt"
	)

	type Rect struct {
		width, length float64
	}

	func (rect Rect) area() float64 {
		return rect.width * rect.length
	}

	func main() {
		var rect = Rect{100, 200}

		fmt.Println("Width:", rect.width, "Length:", rect.length,
			"Area:", rect.area())
	}

咦？这个是什么“内部方法”，根本没有定义在Rect数据类型的内部啊？

确实如此，我们看到，虽然main函数中的rect变量可以直接调用函数area()来获取矩形面积，但是area()函数确实没有定义在Rect结构体内部，这点和C语言的有很大不同。`Go使用组合函数的方式来为结构体定义结构体方法`。我们仔细看一下上面的area()函数定义。

首先是关键字`func`表示这是一个函数，第二个参数是`结构体类型和实例变量`，第三个是`函数名称`，第四个是`函数返回值`。这里我们可以看出area()函数和普通函数定义的`区别就在于`area()函数`多了一个结构体类型限定`。这样一来Go就知道了这是一个为结构体定义的`方法`。

这里需要注意一点就是`定义在结构体上面的函数(function)`一般叫做`方法(method)`。

**结构体和指针**

我们在指针一节讲到过，`指针的主要作用就是在函数内部改变传递进来变量的值`。对于上面的计算矩形面积的例子，我们可以修改一下代码如下：

	package main

	import (
		"fmt"
	)

	type Rect struct {
		width, length float64
	}

	func (rect *Rect) area() float64 {
		return rect.width * rect.length
	}

	func main() {
		var rect = new(Rect)
		rect.width = 100
		rect.length = 200
		fmt.Println("Width:", rect.width, "Length:", rect.length,
			"Area:", rect.area())
	}


上面的例子中，使用了new函数来创建一个结构体指针rect，也就是说rect的类型是\*Rect，结构体遇到指针的时候，你`不需要使用*去访问结构体的成员`，直接使用`.`引用就可以了。所以上面的例子中我们直接使用`rect.width=100` 和`rect.length=200`来设置结构体成员值。因为这个时候rect是结构体指针，所以我们定义area()函数的时候结构体限定类型为`*Rect`。

其实在计算面积的这个例子中，我们不需要改变矩形的宽或者长度，所以定义area函数的时候结构体限定类型仍然为`Rect`也是可以的。如下：

	package main

	import (
		"fmt"
	)

	type Rect struct {
		width, length float64
	}

	func (rect Rect) area() float64 {
		return rect.width * rect.length
	}

	func main() {
		var rect = new(Rect)
		rect.width = 100
		rect.length = 200
		fmt.Println("Width:", rect.width, "Length:", rect.length,
			"Area:", rect.area())
	}

这里Go足够聪明，所以rect.area()也是可以的。

至于`使不使用结构体指针和使不使用指针的出发点是一样的`，那就是`你是否试图在函数内部改变传递进来的参数的值`。再举个例子如下：

	package main

	import (
		"fmt"
	)

	type Rect struct {
		width, length float64
	}

	func (rect *Rect) double_area() float64 {
		rect.width *= 2
		rect.length *= 2
		return rect.width * rect.length
	}

	func main() {
		var rect = new(Rect)
		rect.width = 100
		rect.length = 200
		fmt.Println(*rect)
		fmt.Println("Double Width:", rect.width, "Double Length:", rect.length,
			"Double Area:", rect.double_area())
		fmt.Println(*rect)
	}

这个例子的输出是：

	{100 200}
	Double Width: 200 Double Length: 400 Double Area: 80000
	{200 400}


**结构体内嵌类型**

我们可以在一个`结构体内部定义另外一个结构体类型的成员`。例如iPhone也是Phone，我们看下例子：

	package main

	import (
		"fmt"
	)

	type Phone struct {
		price int
		color string
	}

	type IPhone struct {
		phone Phone
		model string
	}

	func main() {
		var p IPhone
		p.phone.price = 5000
		p.phone.color = "Black"
		p.model = "iPhone 5"
		fmt.Println("I have a iPhone:")
		fmt.Println("Price:", p.phone.price)
		fmt.Println("Color:", p.phone.color)
		fmt.Println("Model:", p.model)
	}

输出结果为

	I have a iPhone:
	Price: 5000
	Color: Black
	Model: iPhone 5

在上面的例子中，我们在结构体IPhone里面定义了一个Phone变量phone，然后我们可以像正常的访问结构体成员一样访问phone的成员数据。但是我们原来的意思是`“iPhone也是(is-a)Phone”`，而这里的结构体IPhone里面定义了一个phone变量，给人的感觉就是`“iPhone有一个(has-a)Phone”`，挺奇怪的。当然Go也知道这种方式很奇怪，所以支持如下做法：


	package main

	import (
		"fmt"
	)

	type Phone struct {
		price int
		color string
	}

	type IPhone struct {
		Phone
		model string
	}

	func main() {
		var p IPhone
		p.price = 5000
		p.color = "Black"
		p.model = "iPhone 5"
		fmt.Println("I have a iPhone:")
		fmt.Println("Price:", p.price)
		fmt.Println("Color:", p.color)
		fmt.Println("Model:", p.model)
	}

输出结果为

	I have a iPhone:
	Price: 5000
	Color: Black
	Model: iPhone 5

在这个例子中，我们定义IPhone结构体的时候，`不再定义Phone变量`，`直接把结构体Phone类型定义在那里`。然后IPhone就可以`像访问直接定义在自己结构体里面的成员一样访问Phone的成员`。

上面的例子中，我们演示了结构体的内嵌类型以及内嵌类型的成员访问，除此之外，假设结构体A内部定义了一个内嵌结构体B，那么A同时也可以调用所有定义在B上面的函数。


	package main

	import (
		"fmt"
	)

	type Phone struct {
		price int
		color string
	}

	func (phone Phone) ringing() {
		fmt.Println("Phone is ringing...")
	}

	type IPhone struct {
		Phone
		model string
	}

	func main() {
		var p IPhone
		p.price = 5000
		p.color = "Black"
		p.model = "iPhone 5"
		fmt.Println("I have a iPhone:")
		fmt.Println("Price:", p.price)
		fmt.Println("Color:", p.color)
		fmt.Println("Model:", p.model)

		p.ringing()
	}

输出结果为：

	I have a iPhone:
	Price: 5000
	Color: Black
	Model: iPhone 5
	Phone is ringing...


**接口**

我们先看一个例子，关于Nokia手机和iPhone手机都能够打电话的例子。

	package main

	import (
		"fmt"
	)

	type NokiaPhone struct {
	}

	func (nokiaPhone NokiaPhone) call() {
		fmt.Println("I am Nokia, I can call you!")
	}

	type IPhone struct {
	}

	func (iPhone IPhone) call() {
		fmt.Println("I am iPhone, I can call you!")
	}
	func main() {
		var nokia NokiaPhone
		nokia.call()

		var iPhone IPhone
		iPhone.call()
	}

我们定义了NokiaPhone和IPhone，它们都有各自的方法call()，表示自己都能够打电话。但是我们想一想，是手机都应该能够打电话，所以这个不算是NokiaPhone或是IPhone的独特特点。否则iPhone不可能卖这么贵了。

再仔细看一下`接口的定义`，首先是关键字`type`，然后是`接口名称`，最后是关键字`interface`表示这个类型是接口类型。`在接口类型里面，我们定义了一组方法`。

Go语言提供了一种接口功能，它把所有的具有共性的方法定义在一起，`任何其他类型只要实现了这些方法就是实现了这个接口`，`不一定非要显式地声明`要去实现哪些接口啦。比如上面的手机的call()方法，就完全可以定义在接口Phone里面，而NokiaPhone和IPhone只要实现了这个接口就是一个Phone。

	package main

	import (
		"fmt"
	)

	type Phone interface {
		call()
	}

	type NokiaPhone struct {
	}

	func (nokiaPhone NokiaPhone) call() {
		fmt.Println("I am Nokia, I can call you!")
	}

	type IPhone struct {
	}

	func (iPhone IPhone) call() {
		fmt.Println("I am iPhone, I can call you!")
	}

	func main() {
		var phone Phone

		phone = new(NokiaPhone)
		phone.call()

		phone = new(IPhone)
		phone.call()

	}

在上面的例子中，我们定义了一个接口Phone，接口里面有一个方法call()，仅此而已。然后我们在main函数里面定义了一个Phone类型变量，并分别为之赋值为NokiaPhone和IPhone。然后调用call()方法，输出结果如下：

	I am Nokia, I can call you!
	I am iPhone, I can call you!

以前我们说过，`Go语言式静态类型语言，变量的类型在运行过程中不能改变`。但是在上面的例子中，phone变量好像先定义为Phone类型，然后是NokiaPhone类型，最后成为了IPhone类型，真的是这样吗？

原来，在Go语言里面，`一个类型A只要实现了接口X所定义的全部方法`，那么`A类型的变量`也是`X类型的变量`。在上面的例子中，NokiaPhone和IPhone都实现了Phone接口的call()方法，所以它们都是Phone，这样一来是不是感觉正常了一些。

我们为Phone添加一个方法sales()，再来熟悉一下接口用法。

	package main

	import (
		"fmt"
	)

	type Phone interface {
		call()
		sales() int
	}

	type NokiaPhone struct {
		price int
	}

	func (nokiaPhone NokiaPhone) call() {
		fmt.Println("I am Nokia, I can call you!")
	}
	func (nokiaPhone NokiaPhone) sales() int {
		return nokiaPhone.price
	}

	type IPhone struct {
		price int
	}

	func (iPhone IPhone) call() {
		fmt.Println("I am iPhone, I can call you!")
	}

	func (iPhone IPhone) sales() int {
		return iPhone.price
	}

	func main() {
		var phones = [5]Phone{
			NokiaPhone{price: 350},
			IPhone{price: 5000},
			IPhone{price: 3400},
			NokiaPhone{price: 450},
			IPhone{price: 5000},
		}

		var totalSales = 0
		for _, phone := range phones {
			totalSales += phone.sales()
		}
		fmt.Println(totalSales)

	}
	
输出结果：

	14200

上面的例子中，我们定义了一个手机数组，然后计算手机的总售价。可以看到，由于NokiaPhone和IPhone都实现了sales()方法，所以它们都是Phone类型，但是计算售价的时候，Go会知道调用哪个对象实现的方法。

接口类型还可以作为结构体的数据成员。

假设有个败家子，iPhone没有出的时候，买了好几款Nokia，iPhone出来后，又买了好多部iPhone，老爸要来看看这小子一共花了多少钱。

	package main

	import (
		"fmt"
	)

	type Phone interface {
		sales() int
	}

	type NokiaPhone struct {
		price int
	}

	func (nokiaPhone NokiaPhone) sales() int {
		return nokiaPhone.price
	}

	type IPhone struct {
		price int
	}

	func (iPhone IPhone) sales() int {
		return iPhone.price
	}

	type Person struct {
		phones []Phone
		name   string
		age    int
	}

	func (person Person) total_cost() int {
		var sum = 0
		for _, phone := range person.phones {
			sum += phone.sales()
		}
		return sum
	}

	func main() {
		var bought_phones = [5]Phone{
			NokiaPhone{price: 350},
			IPhone{price: 5000},
			IPhone{price: 3400},
			NokiaPhone{price: 450},
			IPhone{price: 5000},
		}

		var person = Person{name: "Jemy", age: 25, phones: bought_phones[:]}

		fmt.Println(person.name)
		fmt.Println(person.age)
		fmt.Println(person.total_cost())
	}
	
这个例子纯为演示接口作为结构体数据成员，如有雷同，纯属巧合。这里面我们定义了一个Person结构体，结构体内部定义了一个手机类型切片。另外我们定义了Person的total_cost()方法用来计算手机花费总额。输出结果如下：

	Jemy
	25
	14200
	
**小结**

Go的结构体和接口的实现方法可谓删繁就简，去除了很多别的语言令人困惑的地方，而且学习难度也不大，很容易上手。不过由于思想比较独到，也有可能会有人觉得功能太简单而无用，这个就各有看法了，不过在逐渐的使用过程中，我们会慢慢领悟到这种设计所带来的好处，以及所避免的问题。