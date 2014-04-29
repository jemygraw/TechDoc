#功能
检查传入的参数c是否是可打印字符。如果是则返回非零值，否则返回零值。  
所谓的可打印字符就是所有的非控制字符。
#原型
`#include<ctype.h>`  
`extern int isprint(int c)`
#实例
我们看一个简单的例子。找出所有可打印的字符。
```c
#include<stdio.h>
#include<ctype.h>
int
main (int argc, char *argv[])
{
  char str[7] = { 'a', '1', '\r', ' ', '\n', '\t', '  ' };
  char *p = str;
  while (*p)
  {
    if (iscntrl (*p))
    {
      printf ("'%c' is printable character\n", *p);
    }
    p++;
    }
  return 0;
}
```
输出结果为：
```shell
'a' is printable character
'1' is printable character
' ' is printable character
```
#备注
所有的非控制字符都是可打印字符。可打印字符比可显示字符多了空格字符(' ')，相关函数为isgraph。
