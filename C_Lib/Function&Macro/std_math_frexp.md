#功能
首先看一下函数对应的公式`x=尾数(mantissa)*2^指数(exponent)`。  
而这个函数就是根据参数x来分解出尾数(mantissa)和指数的(exponent)。
#原型
`#include<math.h>`  
`double frexp(double x, int *exponent)`
#实例
```c
#include <stdio.h>
#include <math.h>

int main ()
{
   double x = 1024, fraction;
   int e;

   fraction = frexp(x, &e);
   printf("x = %.2lf = %.2lf * 2^%d\n", x, fraction, e);

   return(0);
}
```
输出结果为：
```shell
x = 1024.00 = 0.50 * 2^11
```
#备注
这个函数将参数x分解为尾数和指数部分。如果参数x不为0，那么正常分解出尾数和指数，其中尾数的绝对值范围为[-0.5,1)。
如果x为0，那么尾数和指数均为0。

在上面的例子中，因为x不为0，所以即使1024可以表示为2<sup>10<sup>，也会被分解为0.5的尾数和11的指数。
