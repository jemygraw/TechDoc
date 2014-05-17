#功能
返回一个double类型参数的小数部分并将整数部分以整型指针方式返回。
#原型
`#include<math.h>`  
`double modf(double x, double *integer)`
#实例
```c
#include<stdio.h>
#include<math.h>

int main ()
{
   double x, fractpart, intpart;

   x = 8.123456;
   fractpart = modf(x, &intpart);

   printf("Integral part = %lf\n", intpart);
   printf("Fraction Part = %lf \n", fractpart);

   return(0);
}
```
输出结果为：
```shell
Integral part = 8.000000
Fraction Part = 0.123456
```
#备注
这个函数就是把一个double类型的数拆分为小数部分和整数部分。
