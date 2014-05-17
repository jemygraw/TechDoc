#功能
计算参数x的指数值，指数的底为e。
#原型
`#include<math.h>`  
`double exp(double x)`
#实例
```c
#include <stdio.h>
#include <math.h>

int main ()
{
   double x = 0;

   printf("The exponential value of %lf is %lf\n", x, exp(x));
   printf("The exponential value of %lf is %lf\n", x+1, exp(x+1));
   printf("The exponential value of %lf is %lf\n", x+2, exp(x+2));

   return(0);
}
```
输出结果为：
```shell
The exponential value of 0.000000 is 1.000000
The exponential value of 1.000000 is 2.718282
The exponential value of 2.000000 is 7.389056
```
#备注
这个函数对应的数学公式为e<sup>x</sup>。
