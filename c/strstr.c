#include <stdio.h>
#include <string.h>
int main() {

  char *path = "/home/test/file/test";

  char *ret = strstr(path, "/home/test");
  printf("result ret= %s\n", ret);

  char *ret2 = strstr(path, "/test/file");
  printf("result ret2 = %s \n", ret2);

  char *ret3 = strstr(path, "file");
  printf("result ret3 = %s \n", ret3);

  printf("strstr(/home/test/1/1/1/1/, /home/test) = %s \n",
         strstr("/home/test/1/1/1/1/", "/home/test"));

  if (strstr(path, "/home/3t/") == NULL) {

    printf("strstr(path,/home/3t) = %s \n", strstr(path, "/home/3t"));
    printf("test string exit \n");
  }
}
