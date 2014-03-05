# Go Base64编码
Go提供了对base64编码和解码的内置支持
```go
package main

// 这种导入包的语法将默认的base64起了一个别名b64，这样
// 我们在下面就可以直接使用b64表示这个包，省点输入量
import b64 "encoding/base64"
import "fmt"

func main() {

	// 这里是我们用来演示编码和解码的字符串
	data := "abc123!?$*&()'-=@~"

	// Go支持标准的和兼容URL的base64编码。
	// 我们这里使用标准的base64编码，这个
	// 函数需要一个`[]byte`参数，所以将这
	// 个字符串转换为字节数组
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 解码一个base64编码可能返回一个错误，
	// 如果你不知道输入是否是正确的base64
	// 编码，你需要检测一些解码错误
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// 使用兼容URL的base64编码和解码
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
```
运行结果
```
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
```
这两种方法都将原数据编码为base64编码，区别在于标准的编码后面是`+`，而兼容URL的编码方式后面是`-`。