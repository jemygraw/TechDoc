#功能
将英文小写字母转换为大写字母。
#原型
`#include<ctype.h>`  
`extern int toupper(int c)`
#实例
```c
#include<stdio.h>
#include<ctype.h>
int
main (int argc, char *argv[])
{
  char *str = "duokr school";
  char *p = str;
  while (*p)
  {
      printf ("%c", toupper (*p));
      p++;
  }
  printf ("\n");
  return 0;
}

```
输出结果为:
```shell
DUOKR SCHOOL
```
#备注
如果存在对应的大写字母，则函数返回大写字母，否则保持不变。比如对于数字字符来讲，就直接返回数字字符。这个函数只支持英文字母。相关函数为tolower。
