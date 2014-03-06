#Go String与Byte切片之间的转换

String转换到Byte数组时，每个byte(byte类型其实就是uint8)保存字符串对应字节的数值。

注意Go的字符串是UTF-8编码的，每个字符长度是不确定的，一些字符可能是1、2、3或者4个字节结尾。

示例1：

```go
package main

import "fmt"

func main() {

	s1 := "abcd"
	b1 := []byte(s1)
	fmt.Println(b1) // [97 98 99 100]

	s2 := "中文"
	b2 := []byte(s2)
	fmt.Println(b2)  // [228 184 173 230 150 135], unicode，每个中文字符会由三个byte组成

	r1 := []rune(s1)
	fmt.Println(r1) // [97 98 99 100], 每个字一个数值

	r2 := []rune(s2)
	fmt.Println(r2) // [20013 25991], 每个字一个数值

}
```