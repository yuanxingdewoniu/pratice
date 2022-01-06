#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX 26

typedef struct TrieNode
{
    int nCount;//记录该字符出现次数
    struct TrieNode* next[MAX];
}TrieNode;

TrieNode Memory[100000];
int allocp=0;

/*初始化*/
void InitTrieRoot(TrieNode* *pRoot)
{
    *pRoot=NULL;
}

/*创建新结点*/
TrieNode* CreateTrieNode()
{
    int i;
    TrieNode *p;

    p=&Memory[allocp++];
    p->nCount=1;

    for(i=0;i<MAX;i++)
    {
        p->next[i]=NULL;
    }

    return p;
}
 
/*插入*/
void InsertTrie(TrieNode* *pRoot,char *s)
{
    int i=0,k=0;
    TrieNode *p;
    int  nCount = 0;
    if(!(p=*pRoot))
    {
        p=*pRoot=CreateTrieNode();
    }
    i=0;

    while(s[i])
    {
        k=s[i++]-'a';//确定branch

        if(!p->next[k])
            p->next[k]=CreateTrieNode();
        else
            p->next[k]-nCount++;

        p=p->next[k];
    }
}

//查找
int SearchTrie(TrieNode* *pRoot,char *s)
{
    TrieNode *p;
    int i,k;

    if(!(p=*pRoot))
    {
        return 0;
    }

    i=0;

    while(s[i])
    {
        k=s[i++]-'a';
        if(p->next[k]==NULL) 
	    return 0;

        p=p->next[k];
    }

    return p->nCount;

} 

int main()
{
    char name[100]={0};
    TrieNode *root;

    InitTrieRoot(&root);    //初始化树根
    do
    {
        printf("scanf symbol \n"); 
	scanf("%s", name);
    InsertTrie(&root,name);

    }while(strcmp(name,"bye")!=0);  //输入"Bye"结束插入

    do
    {
        scanf("%s", name);
        printf("found %d of it!\n",SearchTrie(&root,name));

    }while( strcmp(name,"bye")!=0); //输入"Bye"结束查找

    return 0;
}

