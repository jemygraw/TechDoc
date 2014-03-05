# Go 结构体

Go语言结构体数据类是将各个类型的变量定义的集合，通常用来表示记录。

```go
package main

import "fmt"

// 这个person结构体有name和age成员
type person struct {
	name string
	age  int
}

func main() {

	// 这个语法创建一个新结构体变量
	fmt.Println(person{"Bob", 20})

	// 可以使用"成员:值"的方式来初始化结构体变量
	fmt.Println(person{name: "Alice", age: 30})

	// 未显式赋值的成员初始值为零值
	fmt.Println(person{name: "Fred"})

	// 可以使用&来获取结构体变量的地址
	fmt.Println(&person{name: "Ann", age: 40})

	// 使用点号(.)来访问结构体成员
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 结构体指针也可以使用点号(.)来访问结构体成员
	// Go语言会自动识别出来
	sp := &s
	fmt.Println(sp.age)

	// 结构体成员变量的值是可以改变的
	sp.age = 51
	fmt.Println(sp.age)
}
```

输出结果为

```
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
Sean
50
51
```