# Go 写入文件
Go将数据写入文件的方法和上面介绍过的读取文件的方法很类似。
```go
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 首先看一下如何将一个字符串写入文件
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// 为了实现细颗粒度的写入，打开文件后再写入
	f, err := os.Create("/tmp/dat2")
	check(err)

	// 在打开文件后通常应该立刻使用defer来调用
	// 打开文件的Close方法，以保证main函数结束
	// 后，文件关闭
	defer f.Close()

	// 你可以写入字节切片
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// 也可以使用`WriteString`直接写入字符串
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// 调用Sync方法来将缓冲区数据写入磁盘
	f.Sync()

	// `bufio`除了提供上面的缓冲读取数据外，还
	// 提供了缓冲写入数据的方法
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// 使用Flush方法确保所有缓冲区的数据写入底层writer
	w.Flush()
}
```
运行结果
```
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes
```