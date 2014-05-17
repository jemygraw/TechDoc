#功能
这个函数返回x对y求余的余数。x是分子，y是分母。
#原型
`#include<math.h>`  
`double fmod(double x, double y)`
#实例
```c
#include <stdio.h>
#include <math.h>

int main ()
{
   float a, b;
   int c;
   a = 9.2;
   b = 3.7;
   c = 2;
   printf("Remainder of %f / %d is %lf\n", a, c, fmod(a,c));
   printf("Remainder of %f / %f is %lf\n", a, b, fmod(a,b));

   return(0);
}
```
输出结果为：
```shell
Remainder of 9.200000 / 2 is 1.200000
Remainder of 9.200000 / 3.700000 is 1.800000
```
#备注
这个函数是浮点数求余。
