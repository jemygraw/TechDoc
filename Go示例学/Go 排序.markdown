# Go 排序

Go的sort包实现了内置数据类型和用户自定义数据类型的排序功能。我们先看看内置数据类型的排序。

```go
package main

import "fmt"
import "sort"

func main() {

	// 这些排序方法都是针对内置数据类型的。
	// 这里的排序方法都是就地排序，也就是说排序改变了
	// 切片内容，而不是返回一个新的切片
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// 对于整型的排序
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// 我们还可以检测切片是否已经排序好
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
```
输出结果
```
Strings: [a b c]
Ints:    [2 4 7]
Sorted:  true
```