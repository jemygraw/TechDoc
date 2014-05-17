#功能
计算参数x的底数为10的对数值。
#原型
`#include<math.h>`  
`double log10(double x)`
#实例
```c
#include <stdio.h>
#include <math.h>

int main ()
{
   double x, ret;
   x = 10000;

   ret = log10(x);
   printf("log10(%lf) = %lf\n", x, ret);

   return(0);
}
```
输出结果为：
```shell
log10(10000.000000) = 4.000000
```
#备注
这个函数对应的数学公式为lg(x)。
