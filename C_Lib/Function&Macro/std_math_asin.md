#功能
计算参数x的反正弦值，x的值范围为[-1,1]。
#原型
`#include<math.h>`  
`double asin(double x)`
#实例
现在的例子介绍了asin的用法。
```c
#include <stdio.h>
#include <math.h>

#define PI 3.14159265

int main ()
{
   double x, ret, val;
   x = 0.9;
   val = 180.0 / PI;

   ret = asin(x) * val;
   printf("The arc sine of %lf is %lf degrees\n", x, ret);

   return(0);
}
```
输出结果为：
```shell
The arc sine of 0.900000 is 64.158067 degrees
```
#备注
上面的例子中，我们将结果转换为了角度数显示。因为asin的直接返回结果是弧度数。
