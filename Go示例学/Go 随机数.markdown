# Go 随机数
Go的`math/rand`包提供了伪随机数的生成。
```go
package main

import "fmt"
import "math/rand"

func main() {

	// 例如`rand.Intn`返回一个整型随机数n，0<=n<100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// `rand.Float64` 返回一个`float64` `f`,
	// `0.0 <= f < 1.0`
	fmt.Println(rand.Float64())

	// 这个方法可以用来生成其他数值范围内的随机数，
	// 例如`5.0 <= f < 10.0`
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 为了使随机数生成器具有确定性，可以给它一个seed
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 如果源使用一个和上面相同的seed，将生成一样的随机数
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
}
```
运行结果
```
81,87
0.6645600532184904
7.1885709359349015,7.123187485356329
5,87
5,87
```