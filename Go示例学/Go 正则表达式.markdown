# Go 正则表达式
Go内置了对正则表达式的支持，这里是一般的正则表达式常规用法的例子。

```go
package main

import "bytes"
import "fmt"
import "regexp"

func main() {

	// 测试模式是否匹配字符串，括号里面的意思是
	// 至少有一个a－z之间的字符存在
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 上面我们直接使用了字符串匹配的正则表达式，
	// 但是对于其他的正则匹配任务，你需要使用
	// `Compile`来使用一个优化过的正则对象
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 正则结构体对象有很多方法可以使用，比如上面的例子
	// 也可以像下面这么写
	fmt.Println(r.MatchString("peach"))

	// 这个方法检测字符串参数是否存在正则所约束的匹配
	fmt.Println(r.FindString("peach punch"))

	// 这个方法查找第一次匹配的索引，并返回匹配字符串
	// 的起始索引和结束索引，而不是匹配的字符串
	fmt.Println(r.FindStringIndex("peach punch"))

	// 这个方法返回全局匹配的字符串和局部匹配的字符，比如
	// 这里会返回匹配`p([a-z]+)ch`的字符串
	// 和匹配`([a-z]+)`的字符串
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// 和上面的方法一样，不同的是返回全局匹配和局部匹配的
	// 起始索引和结束索引
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 这个方法返回所有正则匹配的字符，不仅仅是第一个
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// 这个方法返回所有全局匹配和局部匹配的字符串起始索引
	// 和结束索引
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// 为这个方法提供一个正整数参数来限制匹配数量
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	//上面我们都是用了诸如`MatchString`这样的方法，其实
	// 我们也可以使用`[]byte`作为参数，并且使用`Match`
	// 这样的方法名
	fmt.Println(r.Match([]byte("peach")))

	// 当使用正则表达式来创建常量的时候，你可以使用`MustCompile`
	// 因为`Compile`返回两个值
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// regexp包也可以用来将字符串的一部分替换为其他的值
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// `Func`变量可以让你将所有匹配的字符串都经过该函数处理
	// 转变为所需要的值
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
```
运行结果
```
true
true
peach
[0 5]
[peach ea]
[0 5 1 3]
[peach punch pinch]
[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
[peach punch]
true
p([a-z]+)ch
a <fruit>
a PEACH
```
