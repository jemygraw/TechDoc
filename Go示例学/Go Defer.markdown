# Go Defer
Defer 用来保证一个函数调用会在程序执行的最后被调用。通常用于资源清理工作。

```go
package main

import "fmt"
import "os"

// 假设我们想创建一个文件，然后写入数据，最后关闭文件
func main() {
	// 在使用createFile得到一个文件对象之后，我们使用defer
	// 来调用关闭文件的方法closeFile，这个方法将在main函数
	// 最后被执行，也就是writeFile完成之后
	f := createFile("/tmp/defer.txt")
	// Windows下面使用这个语句
	// f := createFile("D:\\Temp\\defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
```
运行结果
```
creating
writing
closing
```
使用defer来调用closeFile函数可以保证在main函数结束之前，关闭文件的操作一定会被执行。