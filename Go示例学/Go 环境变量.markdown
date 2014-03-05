# Go 环境变量
环境变量是一种很普遍的将配置信息传递给Unix程序的机制。
```go
package main

import "os"
import "strings"
import "fmt"
func main() {
	// 为了设置一个key/value对，使用`os.Setenv`
	// 为了获取一个key的value，使用`os.Getenv`
	// 如果所提供的key在环境变量中没有对应的value，
	// 那么返回空字符串
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// 使用`os.Environ`来列出环境变量中所有的key/value对
	// 你可以使用`strings.Split`方法来将key和value分开
	// 这里我们打印所有的key
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
```
这里我们设置了FOO环境变量，所以我们取到了它的值，但是没有设置BAR环境变量，所以值为空。另外我们列出了系统的所有环境变量，当然这个输出根据不同的系统设置可能并不相同。

输出结果
```
FOO: 1
BAR:

TERM_PROGRAM
TERM
SHELL
TMPDIR
Apple_PubSub_Socket_Render
OLDPWD
USER
SSH_AUTH_SOCK
__CF_USER_TEXT_ENCODING
__CHECKFIX1436934
PATH
PWD
ITERM_PROFILE
SHLVL
COLORFGBG
HOME
ITERM_SESSION_ID
LOGNAME
LC_CTYPE
GOPATH
_
FOO
```
