# Go 可变长参数列表

支持可变长参数列表的函数可以支持任意个传入参数，比如fmt.Println函数就是一个支持可变长参数列表的函数。

```go
package main

import "fmt"

// 这个函数可以传入任意数量的整型参数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// 支持可变长参数的函数调用方法和普通函数一样
	// 也支持只有一个参数的情况
	sum(1, 2)
	sum(1, 2, 3)

	// 如果你需要传入的参数在一个切片中，像下面一样
	// "func(slice...)"把切片打散传入
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
```

输出结果为

```
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10
```

需要注意的是，可变长参数应该是函数定义的最右边的参数，即最后一个参数。
