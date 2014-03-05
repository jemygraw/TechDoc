# Go 时间
Go提供了对时间和一段时间的支持。这里有一些例子。
```go
package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	// 从获取当前时间开始
	now := time.Now()
	p(now)

	// 你可以提供年，月，日等来创建一个时间。当然时间
	// 总是会和地区联系在一起，也就是时区
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// 你可以获取时间的各个组成部分
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 输出当天是周几，Monday-Sunday中的一个
	p(then.Weekday())

	// 下面的几个方法判断两个时间的顺序，精确到秒
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Sub方法返回两个时间的间隔(Duration)
	diff := now.Sub(then)
	p(diff)

	// 可以以不同的单位来计算间隔的大小
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 你可以使用Add方法来为时间增加一个间隔
	// 使用负号表示时间向前推移一个时间间隔
	p(then.Add(diff))
	p(then.Add(-diff))
}
```
运行结果
```
2014-03-02 22:54:40.561698065 +0800 CST
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
37578h19m41.910310828s
37578.328308419674
2.2546996985051804e+06
1.3528198191031083e+08
135281981910310828
2014-03-02 14:54:40.561698065 +0000 UTC
2005-08-05 02:15:16.741076409 +0000 UTC
```