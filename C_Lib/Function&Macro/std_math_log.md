#功能
计算参数x的自然(e)对数值。
#原型
`#include<math.h>`  
`double log(double x)`
#实例
```c
#include <stdio.h>
#include <math.h>

int main ()
{
   double x, ret;
   x = 2.7;

   /* 求 log(2.7) */
   ret = log(x);
   printf("log(%lf) = %lf\n", x, ret);

   return(0);
}
```
输出结果为：
```shell
log(2.700000) = 0.993252
```
#备注
这个函数对应的数学公式为ln(x)。
