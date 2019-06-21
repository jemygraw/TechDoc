#Go语言包之 strings
**Contains**(s, substr string) bool  
检查substr是否在字符串s中，如果存在返回true，否则返回false。
```go
package main
import (
	"fmt"
	"strings"
)
func main(){
	s:="hello i am jemy, who are you, what's your name"
	fmt.Println(strings.Contains(s,"name"))
}
//true
```
**ContainsAny**(s, chars string) bool  
检查字符串s中是否存在chars字符串中的任何一个Unicode字符，如果存在则返回true，否则返回false。
比如下面的例子：
```go
strings.ContainsAny(s, "xa") //true
string.ContainsAny(s, "x") //false
```
**ContainsRune**(s string, r rune) bool
检查字符串s中是否存在Unicode字符r，如果存在返回true，否则返回false。
```go
s:="hello i am jemy, who are you, what's your name"
r:=`x`
fmt.Println(strings.ContainsAny(s,r))
//false
```
**Count(s, sep string)** int  
检查字符串sep在字符串s中出现的总次数
```go
s:="we are a team, we are all boys, where are the girls"
sep:="are"
fmt.Println(strings.Count(s,sep))
//3
```
**EqualFold**(s, t string) bool  
检查以UTF-8编码方式解析的字符串s和t是否相同，忽略大小写。
```go
t:="你好，我是小猪猪A"
s:="你好，我是小猪猪a"
fmt.Println(strings.EqualFold(s,t))
//true
```
**Fields**(s string) []string  
Fields方法将字符串s用空白字符分割，空白字符由**unicode.IsSpace**所定义。方法返回分割后的字符串数组。如果字符串s仅仅包含空白字符的话，则返回空数组。
```go
s1="we are a team	we can fight together "
fmt.Println(strings.Fields(s))
s2="	"
fmt.Println(strings.Fields(s))
//[we are a team we can fight together]
//[]
```

**FieldsFunc**(s string, f func(rune) bool) []string  
这个方法判断字符串s中的每个Unicode字符，如果发现用c作为函数f的参数能够让f返回true的话，则用这个字符c来分割字符串s。并返回分割后的字符串数组。如果s字符串中所有的字符都可以让函数f返回true，或者字符串s为空，那么返回一个空字符串数组。
```go
package main
import (
	"fmt"
	"strings"
)
func is_a_or_b(c rune) bool{
	result:=false
	if c>='a' && c <='b'{
		result=true
	}
	return result
}
func is_我(c rune) bool{
	result:=false
	if c == '我'{
		result=true
	}
	return result
}
func main(){
	s:="we are a team 我是大猪猪 "
	r_array:=strings.FieldsFunc(s, is_a_or_b)
	for i:=0; i<len(r_array);i++{
		fmt.Println("["+r_array[i]+"]")
	}
	fmt.Println("---------------------------")
	x_array:=strings.FieldsFunc(s,is_我)
	for i:=0; i<len(x_array);i++{
		fmt.Println("["+x_array[i]+"]")
	}
}
/*
[we ]
[re ]
[ te]
[m 我是大猪猪 ]
---------------------------
[we are a team ]
[是大猪猪 ]
*/
```

**HasPrefix**(s, prefix string) bool  
这个方法检查s是否以字符串prefix开头。
```go
s:="we are a team 我是大猪猪 "
result:=strings.HasPrefix(s, "we")
fmt.Println(result)//true
```

**HasSuffix**(s, suffix string) bool   
这个方法检查s是否以字符串suffix结尾
```go
s:="we are a team 我是大猪猪"
result:=strings.HasSuffix(s, "大猪猪")
fmt.Println(result) //true
```

**Index**(s, sep string) int  
这个方法返回字符串sep在s中第一次出现的位置索引，如果s中不存在sep，那么返回-1。
```go
s:="we are a team , who are you"
fmt.Println(strings.Index(s, "are"))//3
fmt.Println(strings.Index(s, "and"))//-1
```
**IndexAny**(s, chars string) int  
这个方法依次按Unicode字符检查chars字符串中的字符是否存在于s中，如果存在则返回s字符串中第一次出现chars中字符的位置索引，如果chars里面的字符都不存在于s中，那么返回-1。
```go
s:="we are a team , who are you"
chars:="help"
fmt.Println(strings.IndexAny(s, chars))//1
chars="x soilder"
fmt.Println(strings.IndexAny(s, chars))//1,亲空格也是字符哦
chars="go"
fmt.Println(strings.IndexAny(s, chars))//18
chars="silk"
fmt.Println(strings.IndexAny(s, chars))//-1
```

