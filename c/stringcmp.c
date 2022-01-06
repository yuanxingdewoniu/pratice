#include <stdio.h>
#include <string.h>
int main() {

  char *Str1 = "/a/b/c/44d/e";
  char *Str2 = Str1;
  char *Str3 = "/a/b/c";

  printf("Str1-Str3= %s \n", Str1 - Str3);
  return 0;

  char *Str4 = "/a/b/c/44d/e";
  char *Str2_path;
  while (*Str4) {
    if (*Str4 == 'd') {

      printf("Str4 len  = %d, Str4 = %s \n", strlen(Str4), Str4);
      int len = strlen(Str1) - strlen(Str4);
      printf("strlen(Str1)= %d.strlen(Str4) =%d \n", strlen(Str1),
             strlen(Str4));

      //  strncpy(Str2_path, Str1, len);
      printf("Str2_path  %s \n", Str2_path);
      /*
           *Str4++;
          if( *Str4 == '/')
          {
          printf("Str4 = %s \n", Str4 );
              }


              printf(" 移动Str1 = %s, Str3 = %s \n",Str1, Str3);
          if( *Str1 == *Str4) {

          printf(" Str1 == Str3     Str1 = %s, Str2 = %s \n ", Str1, Str2);
              } else {:
              printf(" Str1!= Str2 Str2 =%s \n",Str2);
              }
         */
    }
    ++Str4;
  }
}
