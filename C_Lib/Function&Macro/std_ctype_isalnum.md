#功能
判断字符变量c是否为字母或数字。
即当c为字母a-z或A-Z或者为0-9时，返回非零值，否则返回零。
#原型
`#include <ctype.h>`  
`extern int isalnum(int c)`
#实例
这个函数很简单，这里的字符c是ASCII字符，即'a'-'z','A'-'Z','0'-'9'。
```c
#include<stdio.h>
#include<ctype.h>
int
main(int argc, char *argv[])
{
	int c = 'a';
	printf("'%c' -> %d\n", c, isalnum(c));
	c = '>';
	printf("'%c' -> %d\n", c, isalnum(c));
	c = '1';
	printf("'%c' -> %d\n", c, isalnum(c));
	c = 48;
	printf("'%c' -> %d\n", c, isalnum(c));
	return 0;
}
```
输出结果为
```shell
'a' -> 1
'>' -> 0
'1' -> 1
'0' -> 1
```
#备注
注意这里的字符`c`是ASCII字符，比如将上面的`c`赋值为`48`，那么对应ASCII字符为`'0'`。
