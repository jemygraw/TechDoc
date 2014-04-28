#功能
判断字符c是否为英文字母。
即当c为字母a-z或A-Z时，返回非零值，否则返回零。
#原型
`#include <ctype.h>`  
`extern int isalpha(int c)`
#实例
我们写一个小程序，将输入数据中的字母都过滤出来。
输入数据以';'作为结尾。
```c
#include<stdio.h>
#include<ctype.h>
int
main(int argc, char *argv[])
{
	int		c;
	while ((c = getchar()) != ';') {
		if (isalpha(c)) {
			putchar(c);
		}
	}
	putchar('\n');
	return 0;
}
```
测试结果：
```shell
apple12orange13banana16 melon10;
appleorangebananamelon
```
#备注
这里的字母都是ASCII字母，从'a'-'z'和'A'-'Z'。
