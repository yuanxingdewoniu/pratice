#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int match_str(const char *str, int state) {

  if (state == 1) {

    if (strcmp(str, "/home/test") == 0) {
      printf("return value  3 \n");
      return 3;
    } else {
      return 0;
    }

  } else {
    return 0;
  }
}

int main() {

  int flag = 1;
  int result = 0;
  char *str_tmp = NULL;

  char str1[40] = "/home/test/Hi/Hello/Hair";
  char *str = str1;

  result = match_str(str, flag);
  printf("---- line 40 match_str result = %d \n", result);
  if ((flag == 1) && (result == 0)) {
    int len = strlen(str) - strlen(strrchr(str, '/'));

    str_tmp = (char *)malloc(len + 1);
    if (str_tmp == NULL) {
      printf(" malloc memory failed \n");
      return 0;
    }

    while (len) {

      // str_tmp = (char *) malloc ( len);
      // if (str_tmp == NULL){
      //  printf(" malloc memory failed \n" );
      //  return 0;
      // }

      memset(str_tmp, '\0', len + 1); //字符串要置空
      printf(" len  = %d, str_tmp = %s \n", len, str_tmp);
      strncpy(str_tmp, str, len);
      printf("begin str = %s ,str_tmp = %s \n", str, str_tmp);
      result = match_str(str_tmp, 1);
      printf(" 222 str_tmp = %s , result = %d\n", str_tmp, result);
      len = strlen(str_tmp) - strlen(strrchr(str_tmp, '/'));

      if (result > 0) {
        break;
        printf(" end str = %s ,str_tmp = %s len = %d\n", str, str_tmp, len);
      }
    }
  }

  if (str_tmp)
    free(str_tmp);

  printf("result = %d \n", result);
  return 0;
}
