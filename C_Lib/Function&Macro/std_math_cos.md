#功能
计算一个弧度法表示的角度x的余弦值。
#原型
`#include<math.h>`  
`double cos(double x)`
#实例
我们来看一个小例子：
```c
#include <stdio.h>
#include <math.h>

#define PI 3.14159265

int
main (int argc, char *argv[])
{
   double x, ret, val;

   x = 60.0;
   val =  PI/180;
   ret = cos( x*val );
   printf("The cosine of %lf is %lf\n", x, ret);

   x = 90.0;
   val =   PI/180;
   ret = cos( x*val );
   printf("The cosine of %lf is %lf\n", x, ret);

   return 0;
}
```
输出结果为：
```shell
The cosine of 60.000000 is 0.500000
The cosine of 90.000000 is 0.000000
```
#备注
角度有两种表示方法，一种就是角度数，另一种是弧度数。两者的转换如上面例子所示。
`弧度数=角度数*(PI/180)`。这个函数参数是角度的弧度数表示。