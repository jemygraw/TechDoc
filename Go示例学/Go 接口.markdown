# Go 接口

接口是一个方法签名的集合。
所谓方法签名，就是指方法的声明，而不包括实现。

```go
package main

import "fmt"
import "math"

// 这里定义了一个最基本的表示几何形状的方法的接口
type geometry interface {
    area() float64
    perim() float64
}

// 这里我们要让正方形square和圆形circle实现这个接口
type square struct {
    width, height float64
}
type circle struct {
    radius float64
}

// 在Go中实现一个接口，只要实现该接口定义的所有方法即可
// 下面是正方形实现的接口
func (s square) area() float64 {
    return s.width * s.height
}
func (s square) perim() float64 {
    return 2*s.width + 2*s.height
}

// 圆形实现的接口
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// 如果一个函数的参数是接口类型，那么我们可以使用命名接口
// 来调用这个函数
// 比如这里的正方形square和圆形circle都实现了接口geometry，
// 那么它们都可以作为这个参数为geometry类型的函数的参数。
// 在measure函数内部，Go知道调用哪个结构体实现的接口方法。
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    s := square{width: 3, height: 4}
    c := circle{radius: 5}

    // 这里circle和square都实现了geometry接口，所以
    // circle类型变量和square类型变量都可以作为measure
    // 函数的参数
    measure(s)
    measure(c)
}
```
输出结果为

```
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
```

也就是说如果结构体A实现了接口B定义的所有方法，那么A也是B类型。