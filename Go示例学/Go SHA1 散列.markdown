# Go SHA1 散列
SHA1散列经常用来计算二进制或者大文本数据的短标识值。git版本控制系统用SHA1来标识受版本控制的文件和目录。这里介绍Go中如何计算SHA1散列值。
Go在`crypto/*`包里面实现了几个常用的散列函数。
```go
package main

import "crypto/sha1"
import "fmt"

func main() {
	s := "sha1 this string"

	// 生成一个hash的模式是`sha1.New()`，`sha1.Write(bytes)`
	// 然后是`sha1.Sum([]byte{})`，下面我们开始一个新的hash
	// 示例
	h := sha1.New()

	// 写入要hash的字节，如果你的参数是字符串，使用`[]byte(s)`
	// 把它强制转换为字节数组
	h.Write([]byte(s))

	// 这里计算最终的hash值，Sum的参数是用来追加而外的字节到要
	// 计算的hash字节里面，一般来讲，如果上面已经把需要hash的
	// 字节都写入了，这里就设为nil就可以了
	bs := h.Sum(nil)

	// SHA1散列值经常以16进制的方式输出，例如git commit就是
	// 这样，所以可以使用`%x`来将散列结果格式化为16进制的字符串
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
```
运行结果
```
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94
```