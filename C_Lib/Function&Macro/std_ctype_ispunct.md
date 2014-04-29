#功能
检查参数c是否是标点字符。如果是则返回非零值，否则返回零值。 
所谓的标点字符是指那些不是字母或数字字符(isalnum)的可显示字符(isgraph)。
#原型
`#include<ctype.h>`  
`extern int ispunct(int c)`
#实例
我们看一个简单的例子。找出所有标点字符。
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
    if (ispunct(*p))
    {
      printf ("'%c' is punctuation character\n", *p);
    }
    p++;
    }
  return 0;
}
```
输出结果为：
```shell
',' is punctuation character
';' is punctuation character
```
#备注
这个函数只支持英文标点符号。编程的时候注意输入法的切换。
