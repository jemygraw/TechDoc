# Go 字符串操作函数
strings 标准库提供了很多字符串操作相关的函数。这里提供的几个例子是让你先对这个包有个基本了解。

```go
package main

import s "strings"
import "fmt"

// 这里给fmt.Println起个别名，因为下面我们会多处使用。
var p = fmt.Println

func main() {

	// 下面是strings包里面提供的一些函数实例。注意这里的函数并不是
	// string对象所拥有的方法，这就是说使用这些字符串操作函数的时候
	// 你必须将字符串对象作为第一个参数传递进去。
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	// 你可以在strings包里面找到更多的函数

	// 这里还有两个字符串操作方法，它们虽然不是strings包里面的函数，
	// 但是还是值得提一下。一个是获取字符串长度，另外一个是从字符串中
	// 获取指定索引的字符
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
```
运行结果
```
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       a-b
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
ToLower:    test
ToUpper:    TEST

Len:  5
Char: 101
```
