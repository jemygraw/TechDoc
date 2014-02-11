# Go Switch语句

当条件判断分支太多的时候，我们会使用switch语句来优化逻辑。

```go
package main

import "fmt"
import "time"

func main() {

	// 基础的switch用法
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 你可以使用逗号来在case中分开多个条件。还可以使用default语句，
	// 当上面的case都没有满足的时候执行default所指定的逻辑块。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	// 当switch没有跟表达式的时候，功能和if/else相同，这里我们
	// 还可以看到case后面的表达式不一定是常量。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
}
```

输出结果为

```
write 2 as two
it's a weekday
it's before noon
```
