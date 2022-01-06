#include <stdio.h>
#include <string.h>
 
int main ()
{
   int len;
   const char str[] = "https://www.runoob.com";
   const char ch = '.';
   char *ret;
 
   ret = strrchr(str, ch);



   printf("|%c| 之后的字符串是 - |%s|  strlen =%d  retlen = %d    \n", ch, ret, strlen(str),   strlen(ret));
   len = strlen(str)- strlen(ret);

    char *tmp = (char *) malloc (len);
    if(tmp == NULL) {
    printf("malloc memory failed \n");
    }


    memset(tmp,len, 0);
    memcpy(tmp, str,len);
    printf("str = %s,\n  tmp =%s \n",str, tmp);
    free(tmp);
    tmp = NULL;


   return(0);
}
