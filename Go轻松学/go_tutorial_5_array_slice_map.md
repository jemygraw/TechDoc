# 数组，切片和字典

在上面的章节里面，我们讲过Go内置的基本数据类型。现在我们来看一下Go内置的高级数据类型，数组，切片和字典。

**数组(Array)**

数组是一个具有`相同数据类型`的元素组成的`固定长度`的`有序集合`。比如下面的例子

	var x [5]int
	
表示数组x是一个整型数组，而且数值的长度为5。

`Go提供了几种不同的数组定义方法。`

`最基本的方式就是使用var关键字来定义，然后依次给元素赋值`。`对于没有赋值的元素，默认为零值`。比如对于整数，零值就是0，浮点数，零值就是0.0，字符串，零值就是""，对象零值就是nil。

	package main

	import (
		"fmt"
	)

	func main() {
		var x [5]int
		x[0] = 2
		x[1] = 3
		x[2] = 3
		x[3] = 2
		x[4] = 12
		var sum int
		for _, elem := range x {
			sum += elem
		}
		fmt.Println(sum)
	}

在上面的例子中，我们首先使用`var`关键字来声明，然后给出数组名称`x`，最后说明数组为整型数组，长度为5。然后我们使用索引方式给数组元素赋值。在上面的例子中，我们还使用了一种遍历数组元素的方法。该方法利用Go语言提供的内置函数range来遍历数组元素。`range函数可以用在数组，切片和字典上面`。当`range来遍历数组的时候返回数组的索引和元素值`。在这里我们是对数组元素求和，所以我们对索引不感兴趣。在Go语言里面，`当你对一个函数返回值不感兴趣的话，可以使用下划线(_)来替代它`。另外这里如果我们真的定义了一个索引，在循环结构里面却没有使用索引，Go语言编译的时候还是会报错的。所以用下划线来替代索引变量也是唯一之举了。最后我们输出数组元素的和。

还有一种方式，如果知道了数组的初始值。可以像下面这样定义。

	package main

	import (
		"fmt"
	)

	func main() {
		var x = [5]int{1, 2, 3, 4}
		x[4] = 5

		var sum int
		for _, i := range x {
			sum += i
		}
		fmt.Println(sum)
	}

当然，即使你不知道数组元素的初始值，也可以使用这样的定义方式。

	package main

	import (
		"fmt"
	)

	func main() {
		var x = [5]int{}
		x[0] = 1
		x[1] = 2
		x[2] = 3
		x[3] = 4
		x[4] = 5

		var sum int
		for _, i := range x {
			sum += i
		}
		fmt.Println(sum)
	}

`在这里我们需要特别重视数组的一个特点，就是数组是有固定长度的。`

但是如果我们有的时候也可以不显式指定数组的长度，而是使用`...`来替代数组长度，Go语言会自动计算出数组的长度。不过这种方式定义的数组一定是有初始化的值的。

	package main

	import (
		"fmt"
	)

	func main() {
		var x = [...]string{
			"Monday",
			"Tuesday",
			"Wednesday",
			"Thursday",
			"Friday",
			"Saturday",
			"Sunday"}

		for _, day := range x {
			fmt.Println(day)
		}
	}

在上面的例子中，还需要注意一点就是如果将数组元素定义在不同行上面，那么最后一个元素后面必须跟上`}`或者`,`。上面的例子也可以是这样的。

	package main

	import (
		"fmt"
	)

	func main() {
		var x = [...]string{
			"Monday",
			"Tuesday",
			"Wednesday",
			"Thursday",
			"Friday",
			"Saturday",
			"Sunday",
		}

		for _, day := range x {
			fmt.Println(day)
		}
	}
	
`Go提供的这种可以自动计算数组长度的方法在调试程序的时候特别方便，假设我们注释掉上面数组x的最后一个元素，我们甚至不需要去修改数组的长度。`
	
**切片(Slice)**

在上面我们说过数组是有固定长度的有序集合。这也就是说一旦数组长度定义，你将无法在数组里面多添加哪怕一个元素。数组的这种特点有的时候会成为很大的缺点，尤其是当数组的元素个数不确定的情况下。

所以`切片`诞生了。

切片和数组很类似，甚至你可以理解成数组的子集。但是`切片有一个数组所没有的特点，那就是切片的长度是可变的`。

严格地讲，切片有`容量(capacity)`和`长度(length)`两个属性。

首先我们来看一下切片的定义。切片有两种定义方式，一种是先声明一个变量是切片，然后使用内置函数make去初始化这个切片。另外一种是通过取数组切片来赋值。

	package main

	import (
		"fmt"
	)

	func main() {
		var x = make([]float64, 5)
		fmt.Println("Capcity:", cap(x), "Length:", len(x))
		var y = make([]float64, 5, 10)
		fmt.Println("Capcity:", cap(y), "Length:", len(y))

		for i := 0; i < len(x); i++ {
			x[i] = float64(i)
		}
		fmt.Println(x)

		for i := 0; i < len(y); i++ {
			y[i] = float64(i)
		}
		fmt.Println(y)
	}

输出结果为

	Capcity: 5 Length: 5
	Capcity: 10 Length: 5
	[0 1 2 3 4]
	[0 1 2 3 4]

上面我们首先用make函数定义切片x，这个时候x的容量是5，长度也是5。然后使用make函数定义了切片y，这个时候y的容量是10，长度是5。然后我们再分别为切片x和y的元素赋值，最后输出。

所以使用make函数定义切片的时候，有`两种方式`，一种`只指定长度，这个时候切片的长度和容量是相同的`。另外一种是`同时指定切片长度和容量`。虽然切片的容量可以大于长度，但是`赋值的时候要注意最大的索引仍然是len(x)－1`。否则会报索引超出边界错误。

另外一种是通过数组切片赋值，采用`[low_index:high_index]`的方式获取数值切片，其中切片元素`包括low_index的元素`，但是`不包括high_index的元素`。

	package main

	import (
		"fmt"
	)

	func main() {
		var arr1 = [5]int{1, 2, 3, 4, 5}
		var s1 = arr1[2:3]
		var s2 = arr1[:3]
		var s3 = arr1[2:]
		var s4 = arr1[:]
		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Println(s3)
		fmt.Println(s4)
	}

输出结果为

	[3]
	[1 2 3]
	[3 4 5]
	[1 2 3 4 5]
	


在上面的例子中，我们还省略了low_index或high_index。如果省略了low_index，那么等价于从索引0开始；如果省略了high_index，则默认high_index等于len(arr1)，即切片长度。

这里为了体现切片的长度可以变化，我们看一下下面的例子：

	package main

	import (
		"fmt"
	)

	func main() {
		var arr1 = make([]int, 5, 10)
		for i := 0; i < len(arr1); i++ {
			arr1[i] = i
		}
		fmt.Println(arr1)

		arr1 = append(arr1, 5, 6, 7, 8)
		fmt.Println("Capacity:", cap(arr1), "Length:", len(arr1))
		fmt.Println(arr1)
	}

输出结果为

	[0 1 2 3 4]
	Capacity: 10 Length: 9
	[0 1 2 3 4 5 6 7 8]

这里我们初始化arr1为容量10，长度为5的切片，然后为前面的5个元素赋值。然后输出结果。然后我们再使用Go内置方法append来为arr1追加四个元素，这个时候再看一下arr1的容量和长度以及切片元素，我们发现切片的长度确实变了。

另外我们再用`append`方法给arr1多追加几个元素，试图超过arr1原来定义的容量大小。

	package main

	import (
		"fmt"
	)

	func main() {
		var arr1 = make([]int, 5, 10)
		for i := 0; i < len(arr1); i++ {
			arr1[i] = i
		}

		arr1 = append(arr1, 5, 6, 7, 8, 9, 10)
		fmt.Println("Capacity:", cap(arr1), "Length:", len(arr1))
		fmt.Println(arr1)
	}

输出结果为

	Capacity: 20 Length: 11
	[0 1 2 3 4 5 6 7 8 9 10]
	
我们发现arr1的长度变为11，因为元素个数现在为11个。另外我们发现arr1的容量也变了，变为原来的两倍。这是因为`Go在默认的情况下，如果追加的元素超过了容量大小，Go会自动地重新为切片分配容量，容量大小为原来的两倍`。

上面我们介绍了，可以`使用append函数给切片增加元素`，现在我们再来介绍一个`copy函数用来从一个切片拷贝元素到另一个切片`。

	package main

	import (
		"fmt"
	)

	func main() {
		slice1 := []int{1, 2, 3, 4, 5, 6}
		slice2 := make([]int, 5, 10)
		copy(slice2, slice1)
		fmt.Println(slice1)
		fmt.Println(slice2)
	}
输出结果

	[1 2 3 4 5 6]
	[1 2 3 4 5]
	
在上面的例子中，我们将slice1的元素拷贝到slice2，因为slice2的长度为5，所以最多拷贝5个元素。

总结一下，数组和切片的区别就在于`[]`里面是否有数字或者`...`。因为数值长度是固定的，而切片是可变的。


**字典(Map)**

字典是一组`无序的`，`键值对`的`集合`。

字典也叫做`关联数组`，因为数组通过`索引`来查找元素，而字典通过`键`来查找元素。当然，很显然的，字典的键是不能重复的。如果试图赋值给同一个键，后赋值的值将覆盖前面赋值的值。

字典的定义也有两种，一种是`初始化数据`的定义方式，另一种是`使用神奇的make函数`来定义。

	package main

	import (
		"fmt"
	)

	func main() {
		var x = map[string]string{
			"A": "Apple",
			"B": "Banana",
			"O": "Orange",
			"P": "Pear",
		}

		for key, val := range x {
			fmt.Println("Key:", key, "Value:", val)
		}
	}

输出结果为

	Key: A Value: Apple
	Key: B Value: Banana
	Key: O Value: Orange
	Key: P Value: Pear
	
在上面的例子中，我们定义了一个string:string的字典，其中`[]`之间的是键类型，右边的是值类型。另外我们还看到了`range函数，此函数一样神奇，可以用来迭代字典元素，返回key:value键值对`。当然如果你对键或者值不感兴趣，一样可以使用`下划线(_)`来忽略返回值。


	package main

	import (
		"fmt"
	)

	func main() {
		var x map[string]string

		x = make(map[string]string)

		x["A"] = "Apple"
		x["B"] = "Banana"
		x["O"] = "Orange"
		x["P"] = "Pear"

		for key, val := range x {
			fmt.Println("Key:", key, "Value:", val)
		}
	}

上面的方式就是使用了make函数来初始化字典，`试图为未经过初始化的字典添加元素会导致运行错误`，你可以把使用make函数初始化的那一行注释掉，然后看一下。

当然上面的例子中，我们可以把定义和初始化合成一句。

	package main

	import (
		"fmt"
	)

	func main() {
		x := make(map[string]string)

		x["A"] = "Apple"
		x["B"] = "Banana"
		x["O"] = "Orange"
		x["P"] = "Pear"

		for key, val := range x {
			fmt.Println("Key:", key, "Value:", val)
		}
	}

现在我们再来看一下字典的数据访问方式。如果你访问的元素所对应的键存在于字典中，那么没有问题，如果不存在呢？

这个时候会返回零值。对于字符串零值就是""，对于整数零值就是0。但是对于下面的例子：

	package main

	import (
		"fmt"
	)

	func main() {
		x := make(map[string]int)

		x["A"] = 0
		x["B"] = 20
		x["O"] = 30
		x["P"] = 40

		fmt.Println(x["C"])
	}

在这个例子中，很显然不存在键C，但是程序的输出结果为0，这样就和键A对应的值混淆了。

Go提供了一种方法来解决这个问题：

	package main

	import (
		"fmt"
	)

	func main() {
		x := make(map[string]int)

		x["A"] = 0
		x["B"] = 20
		x["O"] = 30
		x["P"] = 40

		if val, ok := x["C"]; ok {
			fmt.Println(val)
		}
	}
上面的例子中，我们可以看到事实上使用`x["C"]`的返回值有两个，一个是值，另一个是是否存在此键的bool型变量，所以我们看到ok为true的时候就输出键C的值，如果ok为false，那就是字典中不存在这个键。

现在我们再来看看`Go提供的内置函数delete，这个函数可以用来从字典中删除元素`。

	package main

	import (
		"fmt"
	)

	func main() {
		x := make(map[string]int)

		x["A"] = 10
		x["B"] = 20
		x["C"] = 30
		x["D"] = 40

		fmt.Println("Before Delete")
		fmt.Println("Length:", len(x))
		fmt.Println(x)

		delete(x, "A")

		fmt.Println("After Delete")
		fmt.Println("Length:", len(x))
		fmt.Println(x)
	}

输出结果为

	Before Delete
	Length: 4
	map[A:10 B:20 C:30 D:40]
	After Delete
	Length: 3
	map[B:20 C:30 D:40]

我们在删除元素前查看一下字典长度和元素，删除之后再看一下。这里面我们还可以看到`len函数也可以用来获取字典的元素个数`。当然如果你试图删除一个不存在的键，那么程序也不会报错，只是不会对字典造成任何影响。

最后我们再用一个稍微复杂的例子来结束字典的介绍。

我们有一个学生登记表，登记表里面有一组学号，每个学号对应一个学生，每个学生有名字和年龄。

	package main

	import (
		"fmt"
	)

	func main() {
		var facebook = make(map[string]map[string]int)
		facebook["0616020432"] = map[string]int{"Jemy": 25}
		facebook["0616020433"] = map[string]int{"Andy": 23}
		facebook["0616020434"] = map[string]int{"Bill": 22}

		for stu_no, stu_info := range facebook {
			fmt.Println("Student:", stu_no)
			for name, age := range stu_info {
				fmt.Println("Name:", name, "Age:", age)
			}
			fmt.Println()
		}
	}

输出结果为

	Student: 0616020432
	Name Jemy Age 25

	Student: 0616020433
	Name Andy Age 23

	Student: 0616020434
	Name Bill Age 22

当然我们也可以用初始化的方式定义字典：

	package main

	import (
		"fmt"
	)

	func main() {
		var facebook = map[string]map[string]int{
			"0616020432": {"Jemy": 25},
			"0616020433": {"Andy": 23},
			"0616020434": {"Bill": 22},
		}

		for stu_no, stu_info := range facebook {
			fmt.Println("Student:", stu_no)
			for name, age := range stu_info {
				fmt.Println("Name:", name, "Age:", age)
			}
			fmt.Println()
		}
	}


输出结果是一样的。




