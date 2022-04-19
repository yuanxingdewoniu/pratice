#include<string.h>
 
#include<stdio.h>
 
int main(void){
 
   char *str1 = "I like www.dotcpp.com very much!", *str2 = "www.dotcpp.com";
 char *str3 = "3233";
	char *ptr = strstr(str1, str3);
 
   printf("The substring is: %s\n", ptr);
 
   return 0;
 
}
