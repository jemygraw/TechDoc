主要先介绍如下的两组方法：

```
func ToTitle(s string) string
func Title(s string) string
```

和

```
func TrimRight(s string, cutset string) string
func TrimSuffix(s, suffix string) string
```

但从方法的字面意义上面看，好像效果差不多，其实差距大着呢。

例如 `ToTitle`的意思是：
>ToTitle returns a copy of the string s with all Unicode letters mapped to their title case.
>将字符串中所有的Unicode字符都替换为对应的大写形式，然后返回一个新的字符串

而`Title`的意思是：
>Title returns a copy of the string s with all Unicode letters that begin words mapped to their title case.
>将字符串中每个单词的首字母替换为对应的大写形式，然后返回一个新的字符串

**差别就在于`所有的Unicode字符`和`每个单词的首字母。**

所以如下的代码输出就好理解了。

```
package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := "i am a message to test func Title and ToTitle"

	//I AM A MESSAGE TO TEST FUNC TITLE AND TOTITLE
	fmt.Println(strings.ToTitle(msg))

	//I Am A Message To Test Func Title And ToTitle
	fmt.Println(strings.Title(msg))
}
```

------

然后就是`TrimRight`和`TrimSuffix`了，这两个方法甚至坑过大牛(:-P)。

```
func TrimRight(s string, cutset string) string
func TrimSuffix(s, suffix string) string
```

其中`TrimRight`方法的意思如下：
>TrimRight returns a slice of the string s, with all trailing Unicode code points contained in cutset removed.
>将s尾部的所有在cutset中指定的字符移除掉，然后返回一个新的切片

而`TrimSuffix`方法的意思如下：
>TrimSuffix returns s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged.
>如果s以suffix结尾，则返回一个新的移除了suffix的s的切片，如果不已suffix结尾，则返回s本身

**差别就在于一个是移除s中所有的在cutset中的字符，另外一个是移除尾部子字符串**

那么如下的代码就很容易理解了：

```
package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := "i love golang home"

	//i love golang
	fmt.Println(strings.TrimSuffix(msg, "home"))

	//i love golang
	fmt.Println(strings.TrimRight(msg, "omhe"))

	//i love golang
	fmt.Println(strings.TrimRight(msg, "home"))

	//i love golang
	fmt.Println(strings.TrimRight(msg, "meho"))

	//...
}

```

OK，如果你理解了如上方法的差异，基本上就能防坑了。

另外相似的方法差距也可以类推：

```
func TrimLeft(s string, cutset string) string
func TrimPrefix(s, prefix string) string
```

