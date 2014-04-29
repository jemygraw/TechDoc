#功能
检查参数c是否是空白(white-space)字符。如果是则返回非零值，否则返回零值。
#原型
`#include<ctype.h>`  
`extern int isspace(int c)`
#实例
我们看一个简单的例子。找出所有空白字符。
```c
#include<stdio.h>
#include<ctype.h>
int
main (int argc, char *argv[])
{
  char str[7] = { 'a', '1', ',', ' ', '\n', '\t', ';' };
  char *p = str;
  while (*p)
  {
    if (isspace(*p))
    {
      printf ("%d - '%c' is white-space character\n", *p, *p);
    }
    p++;
    }
  return 0;
}
```
输出结果为：
```shell
32 - ' ' is white-space character
10 - '
' is white-space character
9 - '	' is white-space character
```
#备注
标准的空白字符有：
|字符       |ASCII值   |描述                 |
|----------|----------|---------------------|
|' '       |(0x20)    |空格字符(SPC)         |
|'\t'	   |(0x09)    |水平制表位(TAB)       |
|'\n'	   |(0x0a)    |新行(LF)             |
|'\v'	   |(0x0b)    |垂直制表位(VT)        |
|'\f'	   |(0x0c)    |feed(FF)             |
|'\r'	   |(0x0d)    |回车符(CR)           |

从上面我们可以看到空白字符不仅仅指的是空格字符(' ')。
