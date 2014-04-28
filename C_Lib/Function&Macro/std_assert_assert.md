#功能
断言是C语言中用来进行异常处理的一种方式。在程序中启用断言时，当特定的
断言表达式值为假的时候，程序将终止执行并向stderr输出一条信息。
#原型

`#include <assert.h>`  
`void assert(int expression)`

#实例
首先看一下基本用法
```c
#include<stdio.h>
#define NDEBUG
#include<assert.h>
#include<stdlib.h>
int
main(int argc, char *argv[])
{
	FILE           *fp;
	fp = fopen("test1.txt", "w");
	assert(fp);
	fclose(fp);

	fp = fopen("test2.txt", "r");
	assert(fp);
	fclose(fp);
	return 0;
}
```
假设上面的test1.txt和test2.txt都不存在，那么第一个以只写方式打开
test1.txt时，fp不为NULL，所以assert的表达式为真。当第二次以只读
的方式读取不存在的test2.txt时，那么fp就为NULL了，这样assert断言
就为假，所以这段程序的输出结果为：
```shell
Assertion failed: (fp), function main, file assert1.c, line 11.
Abort trap: 6
```
#备注
1. assert是一个宏，并不是一个函数。
2. 使用assert的缺点是，频繁的调用会极大地影响程序的性能，增加额外的开销。
3. assert功能只有在调试模式下生效。
4. 可以在调试结束后，通过在引用`#include <assert.h>`之前添加
`#define NDEBUG`来禁用assert调用。
5. 断言可以用来在函数开始处检查传入参数的合法性。
6. 每个assert只检验一个条件，因为同时检验多个条件时，如果断言失败，无法只管地
判断是那个条件失败了。
7. 不能使用判断断言条件时同时改变表达式中变量值的语句。比如`assert(i++<10)`
这样的语句一旦断言失败，`i++`也不会去执行。
