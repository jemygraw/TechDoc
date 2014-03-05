# Go 指针

Go支持指针，可以用来给函数传递变量的引用。

```go
package main

import "fmt"

// 我们用两个不同的例子来演示指针的用法
// zeroval函数有一个int类型参数，这个时候传递给函数的是变量的值
func zeroval(ival int) {
	ival = 0
}

// zeroptr函数的参数是int类型指针，这个时候传递给函数的是变量的地址
// 在函数内部对这个地址所指向的变量的任何修改都会反映到原来的变量上。
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &操作符用来取得i变量的地址
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 指针类型也可以输出
	fmt.Println("pointer:", &i)
}
```
输出结果为

```
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0xc084000038
```
