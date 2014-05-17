#功能
这个函数和pow功能相同，也是计算x的y次方。区别在于传入的参数为float类型。
#原型
`#include<math.h>`  
`float powf(float x, float y)`
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
这个函数和pow的区别在于所传入的参数为float，结果也为float。
