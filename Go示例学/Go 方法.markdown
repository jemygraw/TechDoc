# Go 方法

一般的函数定义叫做函数，定义在结构体上面的函数叫做该结构体的方法。

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
	// 结构体指针作为函数限定类型。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
```

输出结果为

```go
area:  50
perim: 30
area:  50
perim: 30
```