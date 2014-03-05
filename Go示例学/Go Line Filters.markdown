# Go Line Filters
Line Filters翻译一下大概是行数据过滤器。简单一点就是一个程序从标准输入stdin读取数据，然后处理一下，将处理的结果输出到标准输出stdout。grep和sed就是常见的行数据过滤器。
这里有一个行数据过滤器的例子，是把一个输入文本转换为大写的文本。你可以使用这种方式来实现你自己的Go Line Filters。
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 使用缓冲scanner来包裹无缓冲的`os.Stdin`可以让我们
	// 方便地使用`Scan`方法，这个方法会将scanner定位到下
	// 一行的位置
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text`方法从输入中返回当前行
		ucl := strings.ToUpper(scanner.Text())

		// 输出转换为大写的行
		fmt.Println(ucl)
	}

	// 在`Scan`过程中，检查错误。文件结束不会被当作一个错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
```
运行结果
```
hello world
HELLO WORLD
how are you
HOW ARE YOU
```