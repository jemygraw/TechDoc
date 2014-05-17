#功能
首先看一下函数对应的公式为y=x*2^exponent。其中y为函数计算的结果。
#原型
`#include<math.h>`  
`double ldexp(double x, int exponent)`
#实例
```c
#include <stdio.h>
#include <math.h>

int main ()
{
   double x, ret;
   int n;

   x = 0.65;
   n = 3;
   ret = ldexp(x ,n);
   printf("%f * 2^%d = %f\n", x, n, ret);

   return(0);
}
```
输出结果为：
```shell
0.650000 * 2^3 = 5.200000
```
#备注
这个函数和frexp是逆过程。
