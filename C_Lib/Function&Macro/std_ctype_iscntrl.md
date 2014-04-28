#功能
判断字符c是否为控制字符，即c的ASCII值在0x00-0x1F之间或0x7F(DEL)时，返回非零值，否则返回零值。
#原型
`#include<ctype.h>`  
`extern int iscntrl(int c)`
#实例
我们来看一个简单的例子。
```c
#include<stdio.h>
#include<ctype.h>
int
main(int argc, char *argv[])
{
	int c;
	c = 'a';
	printf("%x:%s\n", c, iscntrl(c) ? "yes" : "no");
	c = 0x0d;
	printf("%x:%s\n", c, iscntrl(c) ? "yes" : "no");
	c = 0x7f;
	printf("%x:%s\n", c, iscntrl(c) ? "yes" : "no");
	return 0;
}
```
输出结果为:
```shell
61:no
d:yes
7f:yes
```
#备注
控制字符（Control Character），出现于特定的信息文本中，表示某一控制功能的字符。  
在ASCII码中，第0～31号及第127号(共33个)是控制字符或通讯专用字符，如控制符：LF（换行）、CR（回车）、FF（换页）、DEL（删除）、BS（退格)、BEL（振铃）等；通讯专用字符：SOH（文头）、EOT（文尾）、ACK（确认）等。
