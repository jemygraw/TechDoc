#Go数值

Go有很多种数据类型，包括字符串类型，整型，浮点型，布尔型等等，这里有几个基础的例子。
```go
package main

import "fmt"

func main() {

	// 字符串可以使用"+"连接
	fmt.Println("go" + "lang")

	//整型和浮点型
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// 布尔型的几种操作符
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
```
输出结果为
```
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
```
