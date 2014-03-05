# Go 方法

一般的函数定义叫做函数，定义在结构体上面的函数叫做该结构体的方法。

示例1：

```go
package main

import "fmt"

type rect struct {
	width, height int
}

// 这个area方法有一个限定类型*rect，
// 表示这个函数是定义在rect结构体上的方法
func (r *rect) area() int {
	return r.width * r.height
}

// 方法的定义限定类型可以为结构体类型
// 也可以是结构体指针类型
// 区别在于如果限定类型是结构体指针类型
// 那么在该方法内部可以修改结构体成员信息
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// 调用方法
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go语言会自动识别方法调用的参数是结构体变量还是
	// 结构体指针，如果你要修改结构体内部成员值，那么使用
	// 结构体指针作为函数限定类型，也就是说参数若是结构体
	//变量，仅仅会发生值拷贝。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
```

输出结果为

```
area:  50
perim: 30
area:  50
perim: 30
```
示例2：

从某种意义上说，方法是函数的“语法糖”。当函数与某个特定的类型绑定，那么它就是一个方法。也证因为如此，我们可以将方法“还原”成函数。

instance.method(args)->(type).func(instance,args)

为了区别这两种方式，官方文档中将左边的称为`Method Value`，右边则是`Method Expression`。Method Value是包装后的状态对象，总是与特定的对象实例关联在一起（类似闭包，拐带私奔），而Method Expression函数将Receiver作为第一个显式参数，调用时需额外传递。

注意：对于Method Expression，T仅拥有T Receiver方法，*T拥有（T+*T）所有方法。
```go
package main

import (
	"fmt"
)

func main() {
	p := Person{2, "张三"}

	p.test(1)
	var f1 func(int) = p.test
	f1(2)
	Person.test(p, 3)
	var f2 func(Person, int) = Person.test
	f2(p, 4)

}

type Person struct {
	Id   int
	Name string
}

func (this Person) test(x int) {
	fmt.Println("Id:", this.Id, "Name", this.Name)
	fmt.Println("x=", x)
}
```
输出结果：
```
Id: 2 Name 张三
x= 1
Id: 2 Name 张三
x= 2
Id: 2 Name 张三
x= 3
Id: 2 Name 张三
x= 4
```
示例3：

使用匿名字段，实现模拟继承。即可直接访问匿名字段（匿名类型或匿名指针类型）的方法这种行为类似“继承”。访问匿名字段方法时，有隐藏规则，这样我们可以实现override效果。
```go
package main

import (
	"fmt"
)

func main() {
	p := Student{Person{2, "张三"}, 25}
	p.test()

}

type Person struct {
	Id   int
	Name string
}

type Student struct {
	Person
	Score int
}

func (this Person) test() {
	fmt.Println("person test")
}

func (this Student) test() {
	fmt.Println("student test")
}

```
输出结果为：
```
student test
```