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