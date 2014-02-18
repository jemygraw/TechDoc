# Go 切片

切片是Go语言的关键类型之一，它提供了比数组更多的功能。

示例1：
```go
package main

import "fmt"

func main() {

	// 和数组不同的是，切片的长度是可变的。
	// 我们可以使用内置函数make来创建一个长度不为零的切片
	// 这里我们创建了一个长度为3，存储字符串的切片，切片元素
	// 默认为零值，对于字符串就是""。
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 可以使用和数组一样的方法来设置元素值或获取元素值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// 可以用内置函数len获取切片的长度
	fmt.Println("len:", len(s))

	// 切片还拥有一些数组所没有的功能。
	// 例如我们可以使用内置函数append给切片追加值，然后
	// 返回一个拥有新切片元素的切片。
	// 注意append函数不会改变原切片，而是生成了一个新切片，
	// 我们需要用原来的切片来接收这个新切片
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 另外我们还可以从一个切片拷贝元素到另一个切片
	// 下面的例子就是创建了一个和切片s长度相同的新切片
	// 然后使用内置的copy函数来拷贝s的元素到c中。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 切片还支持一个取切片的操作 "slice[low:high]"
	// 获取的新切片包含元素"slice[low]"，但是不包含"slice[high]"
	// 下面的例子就是取一个新切片，元素包括"s[2]"，"s[3]"，"s[4]"。
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 如果省略low，默认从0开始，不包括"slice[high]"元素
	l = s[:5]
	fmt.Println("sl2:", l)

	// 如果省略high，默认为len(slice)，包括"slice[low]"元素
	l = s[2:]
	fmt.Println("sl3:", l)

	// 我们可以同时声明和初始化一个切片
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 我们也可以创建多维切片，和数组不同的是，切片元素的长度也是可变的。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
```

输出结果为

```
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]
```

数组和切片的定义方式的区别在于`[]`之中是否有`固定长度`或者推断长度标志符`...`。

示例2：
```go
package main

import "fmt"

func main() {
	s1 := make([]int, 0)
	test(s1)
	fmt.Println(s1)
}

func test(s []int) {
	s = append(s, 3)
	//因为原来分配的空间不够，所以在另外一个地址又重新分配了空间，所以原始地址的数据没有变
}

```
输出结果为：
```
[]
```
若改为：
```go

package main

import "fmt"

func main() {
	s1 := make([]int, 0)
	s1 = test(s1)
	fmt.Println(s1)
}

func test(s []int) []int {
	s = append(s, 3)
	return s
}
```

输出结果为：
```
[3]//正确结果
```
示例3：

cap是slice的最大容量，append函数添加元素，如果超过原始slice的容量，会重新分配底层数组。
```
package main

import "fmt"

func main() {
	s1 := make([]int, 3, 6)
	fmt.Println("s1= ", s1, len(s1), cap(s1))
	s2 := append(s1, 1, 2, 3)
	fmt.Println("s1= ", s1, len(s1), cap(s1))
	fmt.Println("s2= ", s2, len(s2), cap(s2))
	s3 := append(s2, 4, 5, 6)
	fmt.Println("s1= ", s1, len(s1), cap(s1))
	fmt.Println("s2= ", s2, len(s2), cap(s2))
	fmt.Println("s3= ", s3, len(s3), cap(s3))

}
```
输出结果为：
```
s1=  [0 0 0] 3 6
s1=  [0 0 0] 3 6
s2=  [0 0 0 1 2 3] 6 6
s1=  [0 0 0] 3 6
s2=  [0 0 0 1 2 3] 6 6
s3=  [0 0 0 1 2 3 4 5 6] 9 12
```
示例4：

指向同一底层数组的slice之间copy时，允许存在重叠。copy数组时，受限于src和dst数组的长度最小值。
```go
package main

import "fmt"

func main() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := make([]int, 3, 20)
	var n int
	n = copy(s2, s1)
	fmt.Println(n, s2, len(s2), cap(s2))

	s3 := s1[4:6]
	fmt.Println(n, s3, len(s3), cap(s3))

	n = copy(s3, s1[1:5])
	fmt.Println(n, s3, len(s3), cap(s3))
}

```
输出结果：
```
3 [0 1 2] 3 20
3 [4 5] 2 6
2 [1 2] 2 6
```
