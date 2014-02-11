# Go for循环

for循环是Go语言唯一的循环结构。这里有三个基本的for循环类型。

```go
package main

import "fmt"

func main() {

    // 最基本的一种，单一条件循环
    // 这个可以代替其他语言的while循环
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // 经典的循环条件初始化/条件判断/循环后条件变化
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    // 无条件的for循环是死循环，除非你使用break跳出循环或者
    // 使用return从函数返回
    for {
        fmt.Println("loop")
        break
    }
}
```
输出结果
```
1
2
3
7
8
9
loop
```
在后面的例子中，你将会看到其他的循环方式，比如使用range函数循环数组，切片和字典，或者用select函数循环channel通道。