# Go 时间格式化和解析
Go使用模式匹配的方式来支持日期格式化和解析。
```go
package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	// 这里有一个根据RFC3339来格式化日期的例子
	t := time.Now()
	p(t.Format("2006-01-02T15:04:05Z07:00"))

	// Format 函数使用一种基于示例的模式匹配方式，
	// 它使用已经格式化的时间模式来决定所给定参数
	// 的输出格式
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))

	// 对于纯数字表示的时间来讲，你也可以使用标准
	// 的格式化字符串的方式来格式化时间
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// 时间解析也是采用一样的基于示例的方式
	withNanos := "2006-01-02T15:04:05.999999999-07:00"
	t1, e := time.Parse(
		withNanos,
		"2012-11-01T22:08:41.117442+00:00")
	p(t1)
	kitchen := "3:04PM"
	t2, e := time.Parse(kitchen, "8:41PM")
	p(t2)

	// Parse将返回一个错误，如果所输入的时间格式不对的话
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)

	// 你可以使用一些预定义的格式来格式化或解析时间
	p(t.Format(time.Kitchen))
}
```
运行结果
```
2014-03-03T22:39:31+08:00
10:39PM
Mon Mar  3 22:39:31 2014
2014-03-03T22:39:31.647077+08:00
2014-03-03T22:39:31-00:00
2012-11-01 22:08:41.117442 +0000 +0000
0000-01-01 20:41:00 +0000 UTC
parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
10:39PM
```