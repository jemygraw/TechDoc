# Go 闭包函数

Go支持匿名函数，匿名函数可以形成闭包。闭包函数可以访问定义闭包的函数定义的内部变量。

示例1：

```go
package main

import "fmt"

// 这个"intSeq"函数返回另外一个在intSeq内部定义的匿名函数，
// 这个返回的匿名函数包住了变量i，从而形成了一个闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// 我们调用intSeq函数，并且把结果赋值给一个函数nextInt，
	// 这个nextInt函数拥有自己的i变量，这个变量每次调用都被更新。
	// 这里i的初始值是由intSeq调用的时候决定的。
	nextInt := intSeq()

	// 调用几次nextInt，看看闭包的效果
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 为了确认闭包的状态是独立于intSeq函数的，再创建一个。
	newInts := intSeq()
	fmt.Println(newInts())
}
```

输出结果为

```
1
2
3
1
```
示例2：

```go

package main

import "fmt"

func main() {
	add10 := closure(10)//其实是构造了一个加10函数
	fmt.Println(add10(5))
	fmt.Println(add10(6))
	add20 := closure(20)
	fmt.Println(add20(5))
}

func closure(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}

}

```

输出结果为：
```
15
16
25
```
示例3：
```go

package main

import "fmt"

func main() {

	var fs []func() int

	for i := 0; i < 3; i++ {

		fs = append(fs, func() int {

			return i
		})
	}
	for _, f := range fs {
		fmt.Printf("%p = %v\n", f, f())
	}
}

```

输出结果：
```
0x401200 = 3
0x401200 = 3
0x401200 = 3
```
示例4：
```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	result := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(result(i))
	}
}
```

输出结果为：
```
0
1
3
6
10
15
21
28
36
45
```