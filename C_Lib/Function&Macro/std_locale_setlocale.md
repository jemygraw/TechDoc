#功能
这个函数用来设置位置相关的信息。

**参数说明**

 - category 
这个参数是一个命名常量。用来指定locale设定所影响的函数种类，即影响哪些函数的功能。
 - locale
如果locale指定为NULL或者空字符串("")，那么locale的值为环境变量中以上面的category参数为键的环境变量值。

**返回值**

函数调用成功后，返回一个和locale集相关的字符串，如果调用失败则返回NULL。

#原型
`#include<locale.h>`  
`char *setlocale(int category, const char *locale)`
#实例
```c
#include <locale.h>
#include <stdio.h>
#include <time.h>

int
main ()
{
  time_t currtime;
  struct tm *timer;
  char buffer[80];

  time (&currtime);
  timer = localtime (&currtime);

  printf ("Locale is: %s\n", setlocale (LC_ALL, "en_GB"));
  strftime (buffer, 80, "%c", timer);
  printf ("Date is: %s\n", buffer);


  printf ("Locale is: %s\n", setlocale (LC_ALL, "de_DE"));
  strftime (buffer, 80, "%c", timer);
  printf ("Date is: %s\n", buffer);

  return (0);
}
```
输出结果为:
```shell
Locale is: en_GB
Date is: Thu 23 Aug 2012 06:39:32 MST
Locale is: de_DE
Date is: Do 23 Aug 2012 06:39:32 MST
```
#备注
函数参数category的所有常量值如下：
|常量值       |描述                            |
|-------------|-------------------------------|
|LC_ALL       |影响下面所有函数                 |
|LC_COLLATE   |影响字符串比较，查看strcoll       |
|LC_CTYPE     |影响字符分类和转换，例如strtoupper|
|LC_MONETARY  |影响货币格式化，查看localeconv    |
|LC_NUMERIC   |影响十进制分隔符，查看localeconv  |
|LC_TIME      |影响日期和时间格式化，查看strftime |
|LC_MESSAGES  |影响系统响应信息                  |
