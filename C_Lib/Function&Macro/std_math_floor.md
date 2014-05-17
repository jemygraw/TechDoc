#功能
这个函数返回小于或等于参数x的最大整数值。
#原型
`#include<math.h>`  
`double floor(double x)`
#实例
```c
#include <stdio.h>
#include <math.h>

int main ()
{
   float val1, val2, val3, val4;

   val1 = 1.6;
   val2 = 1.2;
   val3 = 2.8;
   val4 = 2.3;

   printf("floor(%f) = %.1lf\n", val1, floor(val1));
   printf("floor(%f) = %.1lf\n", val2, floor(val2));
   printf("floor(%f) = %.1lf\n", val3, floor(val3));
   printf("floor(%f) = %.1lf\n", val4, floor(val4));

   return(0);
}
```
输出结果为：
```shell
floor(1.600000) = 1.0
floor(1.200000) = 1.0
floor(2.800000) = 2.0
floor(2.300000) = 2.0
```
#备注
这个函数和四舍五入不同，而是向下取得一个最靠近这个参数的整数值。
