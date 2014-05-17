#功能
计算参数x的双曲正切函数值。
#原型
`#include<math.h>`  
`double tanh(double x)`
#实例
```c
#include <stdio.h>
#include <math.h>

int main ()
{
   double x, ret;
   x = 0.5;

   ret = tanh(x);
   printf("The hyperbolic tangent of %lf is %lf\n", x, ret);

   return(0);
}
```
输出结果为：
```shell
The hyperbolic tangent of 0.500000 is 0.462117
```
#备注
双曲正切函数的公式为`tanh(x) = sinh(x) / cosh(x)`。
