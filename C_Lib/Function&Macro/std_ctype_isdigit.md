#功能
判断字符c是否是数字字符，即如果是'0'-'9'，则返回非零值，否则返回零值。
#原型
`#include<ctype.h>`  
`extern int isdigit(int c)`
#实例
我们来检测一个字符数组里面有哪些字符是数字字符。
```c
#include<stdio.h>
#include<ctype.h>
#define ARRAY_SIZE 7
int
main(int argc, char *argv[])
{
	char test_array[ARRAY_SIZE] = {'a', '1', '3', 'b', 'd', '4', '0'};
	int i;
	for (i = 0; i < ARRAY_SIZE; i++) {
		if (isdigit(test_array[i])) {
			printf("%c", test_array[i]);
		}
	}
	printf("\n");
	return 0;
}
```
输出结果为：
```shell
1340
```
#备注
这个函数检查的是数字字符，而不是数字。
