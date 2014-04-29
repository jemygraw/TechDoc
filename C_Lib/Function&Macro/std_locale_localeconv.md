#功能
这个函数用来获取位置相关的信息。这些信息以结构体`lconv`的对象返回。

#原型
`#include<locale.h>`  
`struct lconv *localeconv(void)`
#实例
```c
#include <locale.h>
#include <stdio.h>

int main (int argc, char *argv[])
{
   struct lconv *lc;

   setlocale(LC_MONETARY, "it_IT");
   lc = localeconv();
   printf("Local Currency Symbol: %s\n",lc->currency_symbol);
   printf("International Currency Symbol: %s\n",lc->int_curr_symbol);

   setlocale(LC_MONETARY, "en_US");
   lc = localeconv();
   printf("Local Currency Symbol: %s\n",lc->currency_symbol);
   printf("International Currency Symbol: %s\n",lc->int_curr_symbol);

   setlocale(LC_MONETARY, "en_GB");
   lc = localeconv();
   printf ("Local Currency Symbol: %s\n",lc->currency_symbol);
   printf ("International Currency Symbol: %s\n",lc->int_curr_symbol);

   printf("Decimal Point = %s\n", lc->decimal_point);

   return 0;
}
```
输出结果为:
```shell
Local Currency Symbol: EUR
International Currency Symbol: EUR
Local Currency Symbol: $
International Currency Symbol: USD
Local Currency Symbol: £
International Currency Symbol: GBP
Decimal Point = .
```
#备注
这个lconv的结构体定义为：
```c
typedef struct {
   char *decimal_point;
   char *thousands_sep;
   char *grouping;	
   char *int_curr_symbol;
   char *currency_symbol;
   char *mon_decimal_point;
   char *mon_thousands_sep;
   char *mon_grouping;
   char *positive_sign;
   char *negative_sign;
   char int_frac_digits;
   char frac_digits;
   char p_cs_precedes;
   char p_sep_by_space;
   char n_cs_precedes;
   char n_sep_by_space;
   char p_sign_posn;
   char n_sign_posn;
} lconv;
```

