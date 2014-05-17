#功能
计算参数x的平方根。
#原型
`#include<math.h>`  
`double sqrt(double x)`
#实例
```c
#include <stdio.h>
#include <math.h>

int main ()
{

   printf("Square root of %lf is %lf\n", 4.0, sqrt(4.0) );
   printf("Square root of %lf is %lf\n", 5.0, sqrt(5.0) );

   return(0);
}
```
输出结果为：
```shell
Square root of 4.000000 is 2.000000
Square root of 5.000000 is 2.236068
```
#备注
很简单，没啥补充的。
