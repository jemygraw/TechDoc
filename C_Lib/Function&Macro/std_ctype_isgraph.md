#功能
判断一个字符是否为可显示字符，即输出后能看到的字符。
#原型
`#include<ctype.h>`  
`extern int isgraph(int c)`
#实例
我们一样看看一个字符数组中有哪些可显示字符。
```c
#include<stdio.h>
#include<ctype.h>
#define ARRAY_SIZE 7
int
main(int argc, char *argv[])
{
	char		test_array[ARRAY_SIZE] = {' ', '\t', 'a', ',', '2', ')', '='};
	int		i;
	for (i = 0; i < ARRAY_SIZE; i++) {
		if (isgraph(test_array[i])) {
			printf("'%c' is graph character\n", test_array[i]);
		}
	}
	return 0;
}
```
输出结果为：
```shell
'a' is graph character
',' is graph character
'2' is graph character
')' is graph character
'=' is graph character
```
#备注
这个函数可isprint很相似，唯一的区别是isgraph强调的是可显示的字符，所以不包括那些空格，制表位等。
