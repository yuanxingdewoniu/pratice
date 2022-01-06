
/*
 * Copyright (C) 2020 ~ 2021 统信软件技术有限公司
 *
 * Author:  Aaron <zhangya@uniontech.com>
 *
 * Maintainer: Aaron <zhangya@uniontech.com>
 *
 */
#include <stddef.h>
#include "include/trie_tree.h"
#include <linux/limits.h>
//include "include/dlp_macro.h"


#include <sys/types.h>
#include <sys/stat.h>


#include <linux/module.h>

#define PATH_MAX        4096
#define FOLDER_PATH "/etc/filearmor.d/folder.config"
#define FOLDER_PATH_TMP "/etc/filearmor.d/folder_tmp.config"


#define NO_MODULE

#ifndef NO_MODULE
#include <linux/kernel.h>
#include <linux/slab.h>
#define DLP_PRINT printk
#define DLP_MALLOC(size) kmalloc((size), GFP_KERNEL)
#define DLP_FREE kfree
#else
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
  
#define DLP_PRINT(format, ...) \
            printf("[%s:%d->%s] "format, __FILE__, __LINE__, __func__, ##__VA_ARGS__)
//#define DLP_PRINT printf
#define DLP_MALLOC(size) malloc((size))
#define DLP_FREE free
#endif

#ifndef max
#define max(a,b) (((a) > (b)) ? (a) : (b))
#define min(a,b) (((a) < (b)) ? (a) : (b))
#endif

static void (*trie_free_func)(void* ) = NULL;
/*
void trie_print(const dlp_rule* dr)
{
        if (NULL == dr) {
                return;
        }

        DLP_PRINT("dlp_path=%s, path_acl_size = %u", dr->dlp_path,
                  dr->rule.rule_size);

        for (int i = 0; i < dr->rule.rule_size; ++i) {
                DLP_PRINT("    app_rules %d: app_path = %s, uid = %u, acl = %u", i,
                          (dr->rule.app_rules + i)->app_path,
                          (dr->rule.app_rules + i)->uid,
                          (dr->rule.app_rules + i)->acl);
        }
}
*/
void free_data_func(void * data)
{
 printf("free data %p \n",data);
}


void node_print(TireNode* tree, int nest)
{
        if (tree) {
                for (int i = 0; i < nest; ++i) {
                        DLP_PRINT("--");
                }
                DLP_PRINT("|%s\n", tree->node_name);
                //trie_print(tree->data);
        }

        ++nest;

        DLP_PRINT("tree->childs = %p child_count = %d, child_capacity = %d\n",
                  tree->childs, tree->child_count, tree->child_capacity);
        if (tree->childs) {
                for (unsigned int i = 0; i < tree->child_count; ++i) {
                        node_print(*(tree->childs + i), nest);
                }
        }
}

TireNode* trie_create(void)
{
        TireNode* node = (TireNode* ) DLP_MALLOC(sizeof(TireNode));

        memset(node, 0, sizeof(TireNode));

        return node;
}

void  trie_free(TireNode* tree)
{
        int i = 0 ;
        if (tree == NULL) {
                return ;
        }

        if (tree->childs && tree->child_count > 0) {
                for (int i = 0; i < tree->child_count; ++i) {
                        trie_free(*(tree->childs + i));
                        DLP_FREE(*(tree->childs + i));
                }
    
                DLP_FREE(tree->childs);
                tree->childs = NULL;
                i++;

		}

}

TireNode* tire_node_create(const char* node_name, unsigned short node_len)
{
        TireNode* node = (TireNode* ) DLP_MALLOC(sizeof(TireNode));

        node->node_name = (char* ) DLP_MALLOC(node_len + 1);

        memcpy(node->node_name, node_name, node_len);
        node->node_name[node_len] = '\0';

        node->node_len = node_len;

        node->childs = NULL;
        node->child_count = 0;
        node->child_capacity = 0;

        return node;
}

static int tire_compare(const char* path, unsigned short path_len,
                        const TireNode* node)
{
        if (NULL == path) {
                return -1;
        }

        if (NULL == node) {
                return 1;
        }

        if (NULL == node->node_name) {
                return 1;
        }

     DLP_PRINT("path = %s, node->node_name = %s path_len = %d, strncmp(path, node->node_name, path_len) = %d \n",path, node->node_name, path_len,strncmp(path, node->node_name, path_len) );
        return strncmp(path, node->node_name, path_len);
}

size_t tire_recommend(size_t __new_size, size_t cap)
{
        size_t two_cap = 2 * cap;

        return max(two_cap, __new_size);
}

int trie_delete_child(TireNode* tree, const char* path,
                           unsigned short path_len)
{
        int position;
        int res = trie_search(tree, path, path_len, &position);
        if (0 != res) {
                return -1;
        }

        TireNode* child = *(tree->childs + position);
        trie_free(child);

        if (position + 1 == tree->child_count) {
                *(tree->childs + position) = NULL;
                --tree->child_count;
                return 0;
        }

        memmove(tree->childs + position, tree->childs + position + 1,
                (tree->child_count - position) * sizeof(TireNode* ));
        --tree->child_count;
        *(tree->childs + tree->child_count) = NULL;
        return 0;
}
static int trie_bsearch(const char* path, unsigned short path_len,
                                 TireNode** childs, int child_count, int* position)
{
        int __l, __u;
        int __idx = 0;
        const TireNode* __p;
        int __comparison = 0;

        __l = 0;
        __u = child_count;
        while (__l < __u) {

                __idx = (__l + __u) / 2;
                __p = *(childs + __idx);

                __comparison = tire_compare(path, path_len, __p);
                if (__comparison < 0)
                        __u = __idx;
                else if (__comparison > 0)
                        __l = __idx + 1;
                else {
                        DLP_PRINT("-----1------trie_bsearch child_count = %d path = %s ,path_len = %d position = %d \n",
                                 child_count, path, path_len, *position);
                        *position = __idx;
                        return 0;
                }
        }
        if (__comparison > 0) {
                *position = __idx + 1;
                 DLP_PRINT ("------2-------trie_bsearch child_count = %d path = %s ,path_len = %d position = %d \n",
                         child_count, path, path_len, *position);
        } else {
                *position = __idx;
                 DLP_PRINT("--------3------trie_bsearch child_count = %d path = %s ,path_len = %d position = %d \n",
                          child_count, path, path_len, *position);
        }

        return -1;
}

int trie_search(TireNode* tree, const char* path, unsigned short path_len, int* position)
{
        if (NULL == tree || NULL == path) {
                *position = 0;
                return -1;
        }

        if (NULL == tree->childs) {
                *position = 0;
                return -1;
        }

        return trie_bsearch(path, path_len, tree->childs,
                                 tree->child_count, position);
}


TireNode* trie_insert_child(TireNode* tree, const char* path,
                                 unsigned short path_len)
{
        TireNode* child = NULL;
        int position;
        TireNode** temp;

        int res = trie_search(tree, path, path_len, &position);

        //已经存在的数据直接替换
        if(res == 0) {
                child = *(tree->childs + position);
                return child;
        }

        child = tire_node_create(path, path_len);

        if (tree->childs == NULL) {
                tree->childs = DLP_MALLOC(sizeof(TireNode* ));
                *(tree->childs) = child;
                tree->child_count = 1;
                tree->child_capacity = 1;
        } else {
                if (position == tree->child_count) {
                        if (tree->child_capacity < tree->child_count + 1) {
                                tree->child_capacity =
                                        tire_recommend(tree->child_count,
                                                        tree->child_capacity);
                                temp = tree->childs;
                                tree->childs = DLP_MALLOC(tree->child_capacity *
                                                        sizeof(TireNode* ));
                                memcpy(tree->childs, temp, (tree->child_count) *
                                                sizeof(TireNode* ));
                                DLP_FREE(temp);
                        }

                        *(tree->childs + tree->child_count) = child;
                        ++tree->child_count;
                } else {
                        if (tree->child_capacity < tree->child_count + 1) {
                                tree->child_capacity = tire_recommend(tree->child_count,
                                                        tree->child_capacity);
                                temp = tree->childs;
                                tree->childs = DLP_MALLOC(tree->child_capacity *
                                                        sizeof(TireNode* ));
                                if (position > 0) {
                                        memcpy(tree->childs, temp,
                                                position * sizeof(TireNode* ));
                                }
                                *(tree->childs + position) = child;
                                memcpy(tree->childs + position + 1,
                                        temp + position,
                                        (tree->child_count -position) * sizeof(TireNode* ));
                                DLP_FREE(temp);
                        } else {
                                memmove(tree->childs + (position + 1),
                                        tree->childs + (position),
                                        (tree->child_count - position) *
                                        sizeof(TireNode* ));
                                *(tree->childs + position) = child;
                        }

                        ++tree->child_count;
                }
        }

        return child;
}

int trie_insert(TireNode* tree, const char* path, unsigned short path_len)
{
        int get_node_name = 0;
        unsigned short word_size = 0;
        const char* node_start = 0;
        unsigned short i = 0;
        TireNode* node;

         DLP_PRINT("trie_insert path = %s  path_len = %d \n", path, path_len);

        if (tree == NULL) {
                return -1;
        }

        if (path == NULL) {
                return -1;
        }

        for (i = 0; i < path_len; ++i) {
                if (path[i] == '\\' || path[i] == '/') {
                        if (word_size > 0) {
                                get_node_name = 1; //
                        }

                        continue;
                }

                if (word_size == 0) {
                        node_start = &path[i];
                        DLP_PRINT("&path[i] = %s \n", &path[i]);
                }

                if (get_node_name) {
                        // node = trie_insert_child(tree, node_start, word_size,
                        //                            NULL);
                        node = trie_insert_child(tree, node_start, word_size);

                        return trie_insert(node, path + i, path_len - i);
                }
               
                word_size++;
                DLP_PRINT("word_size= %d \n", word_size);
        }

        if (i == path_len) {
                trie_insert_child(tree, node_start, word_size);
                return 0;
        }

        return -1;
}

//static



int trie_delete(TireNode* tree, const char* path, unsigned short path_len)
{
        TireNode* child;
        int get_node_name = 0;
        unsigned short word_size = 0;
        const char* node_start = 0;
        unsigned short i = 0;
        int position;
        int res;

        if (tree == NULL || path == NULL) {
                return -1;
        }

        for (i = 0; i < path_len; ++i) {
                if (path[i] == '\\' || path[i] == '/') {
                        if (word_size > 0) {
                                get_node_name = 1;
                        }

                        continue;
                }

                if (word_size == 0) {
                        node_start = &path[i];
                }

                if (get_node_name) {
                        res = trie_search(tree, node_start, word_size,
                                                 &position);
                        if (res != 0) {
                                return -1;
                        }

                        child = *(tree->childs + position);
                        return trie_delete(child, path + i, path_len - i);
                }

                word_size++;
        }

        if (i == path_len) {
                if (0 != trie_delete_child(tree, node_start, word_size)) {
                        return -1;
                }
        }

        return 0;
}

void* trie_get(TireNode* tree, const char* path,
                         unsigned short path_len)
{
        int get_node_name = 0;
        unsigned short word_size = 0;
        const char* node_start = 0;
        unsigned short i = 0;
        int position;
        void* result = NULL;
        TireNode* child;
        void* pd;
        int res;

        DLP_PRINT("^^^^^^^111^^^^^^trie_get path = %s, path_len = %d tree = %p \n", path, path_len, tree);

        if (tree == NULL) {
                return NULL;
        }

        if (path == NULL) {
                return NULL;
        }

        for (i = 0; i < path_len; ++i) {
                if (path[i] == '\\' || path[i] == '/') {
                        if (word_size > 0) {
                                get_node_name = 1;
                        }

                        continue;
                }

                if (word_size == 0) {
                        node_start = path + i;
                }

                if (get_node_name) {
                        res = trie_search(tree, node_start, word_size,
                                                 &position);
                      

                        if (res != 0) {
                          //返回当前的路径信息，获取父目录,
                            DLP_PRINT("------返回当前的路径信息，获取父目录 ------node_start = %s\n", node_start);
                          return (char *)node_start;
                        }


                       DLP_PRINT("node_start =  %s\n", node_start);
                        child = *(tree->childs + position);

        
                        pd = trie_get(child, path + i, path_len - i);
                        if (pd) {
                              DLP_PRINT("pd = %s \n", pd);
                                return pd;
                        }

                        return result;
                }

                word_size++;
        }

        return result;
}

void trie_set_free_func(void (*free_func)(void* ))
{
        trie_free_func = free_func;
}


char* check_folder_rules(const char *name)
{

        FILE *profile_file = NULL;
        char *return_path;
	char StrLine[256] = {0};
        profile_file  =  fopen(FOLDER_PATH, "r");

        if ( profile_file   == NULL)
        {
            DLP_PRINT("open file failed \n");

        }

        while (!feof(profile_file)) {
        fgets(StrLine, 255, profile_file);  //读取一行
         DLP_PRINT("get path strlen(StrLine)= %d, name = %s \n StrLine = %s strstr(StrLine, name) = %s  \n ",  strlen(StrLine),   name, StrLine,strstr(StrLine, name));


        int len = strlen (StrLine)-1;
         DLP_PRINT("---------len =%d \n", len);
	if ( strncmp(name,StrLine,len ) == 0)
	{
        DLP_PRINT("get path strlen(StrLine)= %d, name = %s \n strlen(StrLine) = %d ------ strstr  result = %d \n ", 
                strlen(StrLine),   name, strlen(StrLine), strncmp(name,StrLine,len ));

                return_path = (char *)malloc (len +1 );
                memset(return_path, '0', len);
                DLP_PRINT("return_path = %s, StrLine = %s, -----len = %d \n", return_path, StrLine, len);
                strncpy(return_path,StrLine,len);
                *(return_path+len) = '\0';

                DLP_PRINT(" @@@@ StrLine = %s, path = %s", StrLine, return_path);
                return return_path;
           
	}
       }
      DLP_PRINT("############  \n");
      return NULL;
}


int serialize_path(const char* path, void* data)
{
        int path_len;

        if (path == NULL) {
                return -1;
        }

        path_len = strlen(path);
        if (path_len >= PATH_MAX) {
                return -1;
        }

        *((int*) data) = path_len + 1;
        memcpy((char* ) data + sizeof(int), path, path_len + 1);

        return sizeof(int) + path_len + 1;
}

int unserialize_path(const void* data, char** path)
{
        int len = 0;

//    *path = NULL;
        if (data == NULL) {
                return -1;
        }

        len = *((int*) data);

        if (len > PATH_MAX || len < 0) {
                return -1;
        }
		
	*path =(void *)  DLP_MALLOC(len);
         memcpy(*path, data + sizeof(int), len);
        return len + sizeof(int);
}


void test_serialize_path(char* path)
{
        void* data = DLP_MALLOC(PATH_MAX);

        int size = serialize_path(path, data);
        DLP_PRINT("serialize_path size = %d\n", size);

       char* path_out1 = NULL ;

      size = unserialize_path(data, &path_out1);
      DLP_PRINT("unserialize_path size = %d, out = %s\n", size, path_out1);
     
	  DLP_FREE(data);
}


void add_path(const char *path)
{

        struct stat s_buf;
 
        /*获取文件信息，把信息放到s_buf中*/
        stat(path,&s_buf);

        /*判断是否目录,非目录不做处理*/
        DLP_PRINT("%d", S_ISDIR(s_buf.st_mode));
        if(S_ISDIR(s_buf.st_mode) == 0 )
        {
               return ;
        }
        
        int count = 0;
        FILE *profile_file = NULL;
        FILE *profile_file_tmp = NULL;
        char StrLine[256] = {0};
        profile_file  =  fopen(FOLDER_PATH, "r");
        profile_file_tmp  =  fopen(FOLDER_PATH_TMP, "w+");

        if ( profile_file   == NULL) {
            DLP_PRINT("open file failed \n");

        }

        if ( profile_file_tmp   == NULL) {
            DLP_PRINT("open file failed \n");

        }

            int len = 0;
            while( ( fgets(StrLine, 255, profile_file)) != NULL ) {

        int len = strlen(StrLine);

        DLP_PRINT("strlen(StrLine) = %d \n", strlen(StrLine));
        char StrLine_tmp[len];
        memset(StrLine_tmp, 0, len  );

        DLP_PRINT("StrLine_tmp = %s strlen(StrLine_tmp) = %d len = %d\n", StrLine_tmp, strlen(StrLine_tmp), len );
   

        strncpy(StrLine_tmp, StrLine, len-1);  //去除换行符号
        DLP_PRINT("StrLine_tmp = %s path = %s  strcmp(StrLine_tmp, path) = %d \n", StrLine_tmp, path,strcmp(StrLine_tmp, path));

        if (strcmp(StrLine_tmp, path) == 0){
        count = 1; // 该路径存在
        } 

       memset(StrLine_tmp,'\0', len);
       DLP_PRINT("StrLine_tmp = %s", StrLine_tmp);

        fputs(StrLine, profile_file_tmp);
        memset(StrLine,0,255);
        }
        
         if(count == 0){
           fputs(path, profile_file_tmp);
           fputs("\n",profile_file_tmp) ;
        } 

        fclose(profile_file_tmp);
        fclose(profile_file);

 rename(FOLDER_PATH_TMP, FOLDER_PATH);


}



void delete_path(const char *path)
{


        struct stat s_buf;
        /*获取文件信息，把信息放到s_buf中*/
        stat(path,&s_buf);

        /*判断是否目录,非目录不做处理*/
        DLP_PRINT("%d", S_ISDIR(s_buf.st_mode));
        if(S_ISDIR(s_buf.st_mode) == 0 )
        {
               return ;
        }


        int count = 0;
        FILE *profile_file = NULL;
        FILE *profile_file_tmp = NULL;
        char StrLine[256] = {0};
        profile_file  =  fopen(FOLDER_PATH, "r");
        profile_file_tmp  =  fopen(FOLDER_PATH_TMP, "w+");

        if ( profile_file   == NULL) {
            DLP_PRINT("open file failed \n");
        }

        if ( profile_file_tmp   == NULL) {
            DLP_PRINT("open file failed \n");
        }

    while( (fgets(StrLine, 255, profile_file)) != NULL ) {
            DLP_PRINT("StrLine = %s \n  path = %s\n  strcmp(path, Strline) = %d \n", StrLine, path ,strncmp(StrLine, path,strlen(path)));
          
            if ( strncmp(StrLine, path,strlen(path)) != 0 ){
                DLP_PRINT("StrLine = %s \n", StrLine);
                fputs(StrLine, profile_file_tmp);
            }
    }

        fclose(profile_file_tmp);
        fclose(profile_file);
        rename(FOLDER_PATH_TMP, FOLDER_PATH);


}




int main()
{

        TireNode* tree = trie_create();


		// trie_set_free_func(free_data_func);


       // add_path("/home/test");

       // DLP_PRINT(" add /home/test \n");

       //  add_path("/home/wxl/test/1/1");
       // DLP_PRINT(" add /home/wxl/test/1/1/2\n");

       // add_path("/home/1/test/1/1");
       // DLP_PRINT("add /home/1/test/1/1 \n");




       // add_path("/home/test");
       // DLP_PRINT(" add /home/test \n");

       //  add_path("/home/wxl/test/1/1");
       // DLP_PRINT(" add /home/wxl/test/1/1/2\n");

       // add_path("/home/1/test/1/1");
       // DLP_PRINT("add /home/1/test/1/1 \n");


       // add_path("/home/test");
       // DLP_PRINT(" add /home/test \n");

       //  add_path("/home/wxl/test/1/1");
       // DLP_PRINT(" add /home/wxl/test/1/1/2\n");

       // add_path("/home/1/test/1/1");
       // DLP_PRINT("add /home/1/test/1/1 \n");


       delete_path("/home/wxl/work/demo/c");
       DLP_PRINT(" delete_path  /home/wxl/work/demo/c\n");


        char *name = "1/home/wxl/1/test/test1/1111/3F";
        char *return_path;
        return_path = check_folder_rules(name);

       if (return_path != NULL)
       {
        printf("-----------check folder_rules return_path= %s -------------- \n", return_path);
        free( return_path);
       } else {
         DLP_PRINT("1111 test 1111 \n" );
         return 0;
       }
       DLP_PRINT("test 1111 \n" );
         DLP_PRINT("test 2222 \n" );

       


      

      // delete_path("/home/wxl/1/test/test1");
      //  DLP_PRINT(" delete_path /home/wxl/1/test/test1\n");

       
      // delete_path("/home/1/test/1/1");
      //  DLP_PRINT(" delete_path /home/1/test/1/1\n");


		//  trie_insert(tree, StrLine, strlen(StrLine), test);

        /***插入的位置可以发现，那么就在目录中****/

            // void *test7= malloc(1);
            // trie_insert(tree, 
            // "/home/5wxl/test/3/5/",
            // strlen("/home/5wxl/test/3/5"),
            // test7);
           

            // DLP_PRINT("####333####/home/5wxl/test/3/5/1/2\n");
            // const char *res1 = trie_get(tree,  "/home/5wxl/test/3/5/1/2",
            // strlen( "/home/5wxl/test/3/5/1/2"));

            // DLP_PRINT(" ****33***** res = %s \n", res1);

            // int path_len = strlen( "/home/5wxl/test/3/5/1/2") - strlen(res1) -1;
            // char *path = (char *) malloc (path_len);
            // memset(path, '0', path_len);
            
            // strncpy(path,"/home/5wxl/test/3/5/1/2",path_len);
            // printf(" path =  %s\n", path );

            // DLP_PRINT("#####333###/home/5wxl/test/3/5/1/2\n");

            // trie_free(tree);

/*
            FILE *profile_file = NULL;
            char StrLine[256] = {0};
            profile_file  =  fopen(FOLDER_PATH, "r");


            if ( profile_file   == NULL)
            {
                DLP_PRINT("open file failed \n");

            }


            while (!feof(profile_file)) {
            fgets(StrLine, 255, profile_file);  //读取一行

            DLP_PRINT("^^^^^^^^^^^^^ Strlen(StrLine) = %d  Strline = %s\n^^^^^^^^^^^^^^^^^",strlen(StrLine), StrLine );
            trie_insert(tree, StrLine, strlen(StrLine)); 
            }

            const char *res1 = trie_get(tree,  "/home/",
            strlen( "/home/"));

            const char *res2 = trie_get(tree,  "/home/wxl/test/test1/home/wxl/test/1/1/3/4/4",
            strlen( "/home/wxl/test/test1/home/wxl/test/1/1/3/4/4"));

            DLP_PRINT(" res1 = %s   res2 =  %s \n", res1, res2);
            //增加检查 路径是否匹配

            if (res2 != NULL)
            {
                
            int path_len = strlen( "/home/wxl/test/test1/home/wxl/test/1/1/3/4/4") - strlen(res2) -1;
            char *path = (char *) malloc (path_len);
            memset(path, '0', path_len);

            strncpy(path,"/home/wxl/test/test1/home/wxl/test/1/1/3/4/4",path_len);
            printf(" path =  %s\n", path );

            }

            trie_free(tree);
        */

   return 0;
}
