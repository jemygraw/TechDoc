#功能
计算参数x的反正切值，返回结果范围为[-PI/2,PI/2]。
#原型
`#include<math.h>`  
`double atan(double x)`
#实例
```c
#include <stdio.h>
#include <math.h>

#define PI 3.14159265

int
main()
{
	double x, ret, val;
	x = 1.0;
	val = 180.0 / PI;

	ret = atan(x) * val;
	printf("The arc tangent of %lf is %lf degrees\n", x, ret);

	return (0);
}
```
输出结果为：
```shell
The arc tangent of 1.000000 is 45.000000 degrees
```
#备注
上面的例子中，我们将结果转换为了角度数显示。因为atan的直接返回结果是弧度数。
