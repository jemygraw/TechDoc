#功能
计算y/x的反正切值，返回结果范围为[-PI,PI]。其中y为纵坐标的值，x为横坐标的值。
返回值会自动根据y和x的正负来确定，比如y=1，x=-1和y=-1，x=1，虽然它们的y/x的
值都是-1，但是y/x的反正切值是不一样的。
#原型
`#include<math.h>`  
`double atan2(double y, double x)`
#实例
```c
#include <stdio.h>
#include <math.h>

#define PI 3.14159265

int main ()
{
   double x, y, ret, val;

   x = -7.0;
   y = 7.0;
   val = 180.0 / PI;

   ret = atan2 (y,x) * val;
   printf("The arc tangent of x = %lf, y = %lf ", x, y);
   printf("is %lf degrees\n", ret);

   return(0);
}
```
输出结果：
```shell
The arc tangent of x = -7.000000, y = 7.000000 is 135.000000 degrees
```
#备注
需要注意的就是在y/x的值一样的情况下，所返回的atan2值也有可能会不同，因为所返回的值还会根据y和x的正负来判断。
