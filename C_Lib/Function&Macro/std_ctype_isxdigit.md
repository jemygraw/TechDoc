#功能
检查参数c是否是十六进制字符。如果是返回非零值，否则返回零值。
#原型
`#include<ctype.h>`  
`extern int isxdigit(int c)`
#实例
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
    if (isxdigit(*p))
    {
      printf ("'%c' is xdigit character\n", *p, *p);
    }
    p++;
    }
  return 0;
}
```
输出结果为:
```shell
'a' is xdigit character
'1' is xdigit character
```
#备注
十六进制数以0x开头。可以包含的字符有'0'-'9','a'-'f','A'-'F'。
