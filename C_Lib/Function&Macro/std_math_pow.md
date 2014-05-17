#功能
计算参数x的y次方值。
#原型
`#include<math.h>`  
`double pow(double x, double y)`
#实例
```c
#include <stdio.h>
#include <math.h>

int main ()
{
   printf("Value 8.0 ^ 3 = %lf\n", pow(8.0, 3));

   printf("Value 3.05 ^ 1.98 = %lf\n", pow(3.05, 1.98));

   return(0);
}
```
输出结果为：
```shell
Value 8.0 ^ 3 = 512.000000
Value 3.05 ^ 1.98 = 9.097324
```
#备注
这个函数对应的数学公式为x<sup>y</sup>。
