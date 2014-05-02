#功能
计算参数x的反余弦值，x的值范围为[-1,1]，函数返回值范围为[0,PI]。
#原型
`#include<math.h>`  
`double acos(double x)`
#实例
现在的例子介绍了acos的用法。
```c
#include <stdio.h>
#include <math.h>

#define PI 3.14159265

int
main()
{
	double	x, ret, val;

	x = 0.9;
	val = 180.0 / PI;

	ret = acos(x) * val;
	printf("The arc cosine of %lf is %lf degrees\n", x, ret);

	return (0);
}
```
输出结果为：
```shell
The arc cosine of 0.900000 is 25.841933 degrees
```
#备注
上面的例子中，我们将结果转换为了角度数显示。因为acos的直接返回结果是弧度数。
