# Go 进程触发
有的时候，我们需要从Go程序里面触发一个其他的非Go进程来执行。

```go
package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {

	// 我们从一个简单的命令开始，这个命令不需要任何参数
	// 或者输入，仅仅向stdout输出一些信息。`exec.Command`
	// 函数创建了一个代表外部进程的对象
	dateCmd := exec.Command("date")

	// `Output`是另一个运行命令时用来处理信息的函数，这个
	// 函数等待命令结束，然后收集命令输出。如果没有错误发
	// 生的话，`dateOut`将保存date的信息
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// 下面我们看一个需要从stdin输入数据的命令，我们将
	// 数据输入传给外部进程的stdin，然后从它输出到stdout
	// 的运行结果收集信息
	grepCmd := exec.Command("grep", "hello")

	// 这里我们显式地获取input/output管道，启动进程，
	// 向进程写入数据，然后读取输出结果，最后等待进程结束
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// 在上面的例子中，我们忽略了错误检测，但是你一样可以
	// 使用`if err!=nil`模式来进行处理。另外我们仅仅收集了
	// `StdoutPipe`的结果，同时你也可以用一样的方法来收集
	// `StderrPipe`的结果
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 注意，我们在触发外部命令的时候，需要显式地提供
	// 命令和参数信息。另外如果你想用一个命令行字符串
	// 触发一个完整的命令，你可以使用bash的-c选项
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
```
所触发的程序的执行结果和我们直接执行这些程序的结果是一样的。
运行结果
```
> date
Wed Oct 10 09:53:11 PDT 2012
> grep hello
hello grep
> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go
```