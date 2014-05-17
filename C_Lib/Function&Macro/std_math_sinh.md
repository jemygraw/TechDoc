#功能
计算参数x的双曲正弦函数值。
#原型
`#include<math.h>`  
`double sinh(double x)`
#实例
```c
#include <stdio.h>
#include <math.h>

int main ()
{
   double x, ret;
   x = 0.5;

   ret = sinh(x);
   printf("The hyperbolic sine of %lf is %lf\n", x, ret);

   return(0);
}
```
输出结果为：
```shell
The hyperbolic sine of 0.500000 is 0.521095
```
#备注
双曲正弦函数的公式为`sinh(x)=(exp(x) - exp(-x)) / 2.0`。
