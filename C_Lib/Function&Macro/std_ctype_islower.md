#功能
判断参数c是否是小写英文字母，如果是返回非零值，否则返回零值。
#原型
`#include<ctype.h>`  
`extern int islower(int c)`
#实例
我们看一个简单的例子。
```c
#include<stdio.h>
#include<ctype.h>
int
main(int argc, char *argv[])
{
	char str[8] = "AbcDEFG";
	char *p = str;
	while (*p) {
		if (islower(*p)) {
			printf("%c is lower-case character\n", *p);
		}
		p++;
	}
	return 0;
}
```
输出结果为：
```shell
b is lower-case character
c is lower-case character
```
#备注
这个函数仅支持英文字母。相关的函数为isupper。
