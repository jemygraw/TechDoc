#功能
计算参数x的双曲余弦函数值。
#原型
`#include<math.h>`  
`double cosh(double x)`
#实例
```c
#include <stdio.h>
#include <math.h>

int main ()
{
   double x;

   x = 0.5;
   printf("The hyperbolic cosine of %lf is %lf\n", x, cosh(x));

   x = 1.0;
   printf("The hyperbolic cosine of %lf is %lf\n", x, cosh(x));

   x = 1.5;
   printf("The hyperbolic cosine of %lf is %lf\n", x, cosh(x));

   return(0);
}
```
输出结果为：
```shell
The hyperbolic cosine of 0.500000 is 1.127626
The hyperbolic cosine of 1.000000 is 1.543081
The hyperbolic cosine of 1.500000 is 2.352410
```
#备注
双曲余弦函数的公式为`cosh(x)=(exp(x) + exp(-x)) / 2.0`。
