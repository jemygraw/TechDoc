# Go 集合功能
我们经常需要程序去处理一些集合数据，比如选出所有符合条件的数据或者使用一个自定义函数将一个集合元素拷贝到另外一个集合。

在一些语言里面，通常是使用泛化数据结构或者算法。但是Go不支持泛化类型，在Go里面如果你的程序或者数据类型需要操作集合，那么通常是为集合提供一些操作函数。

这里演示了一些操作strings切片的集合函数，你可以使用这些例子来构建你自己的函数。注意在有些情况下，使用内联集合操作代码会更清晰，而不是去创建新的帮助函数。

```go
package main

import "strings"
import "fmt"

// 返回t在vs中第一次出现的索引，如果没有找到t，返回－1
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// 如果t存在于vs中，那么返回true，否则false
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 如果使用vs中的任何一个字符串作为函数f的参数可以让f返回true，
// 那么返回true，否则false
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 如果分别使用vs中所有的字符串作为f的参数都能让f返回true，
// 那么返回true，否则返回false
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// 返回一个新的字符串切片，切片的元素为vs中所有能够让函数f
// 返回true的元素
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 返回一个bool类型切片，切片的元素为vs中所有字符串作为f函数
// 参数所返回的结果
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	// 来，试试我们的字符串切片操作函数
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	// 上面的例子都使用匿名函数，你也可以使用命名函数
	fmt.Println(Map(strs, strings.ToUpper))
}
```
运行结果
```
2
false
true
false
[peach apple pear]
[PEACH APPLE PEAR PLUM]
```