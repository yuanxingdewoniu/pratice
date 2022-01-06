/*
 * Copyright (C) 2020 ~ 2021 统信软件技术有限公司
 *
 * Author:  Aaron <zhangya@uniontech.com>
 *
 * Maintainer: Aaron <zhangya@uniontech.com>
 *
 */

#pragma once

//include "dlp.h"


typedef struct _TrieNode {
        char* node_name;
        unsigned short node_len;

 

        struct _TrieNode** childs;
        int child_count;
        unsigned int child_capacity; //预申请的容量,以两倍方式扩充容量
} TireNode;

// 打印d整个trie tree, 调试用
void node_print(TireNode* tree, int nest);

// 创建一个TireNode
TireNode* trie_create(void);

// 递归释放TireNode
void trie_free(TireNode* tree);

// 插入数据
int trie_insert(TireNode* tree, const char* path, unsigned short path_len);

// 搜索数据
int trie_search(TireNode* tree, const char* path, unsigned short path_len,
                int* position);

// 删除数据
int trie_delete(TireNode* tree, const char* path, unsigned short path_len);

// 获取数据 路径的数据不存在 若dlp data为NULL，使用父路径的dlp数据
void* trie_get(TireNode* tree, const char* path, unsigned short path_len);

// 设置数据释放函数 当插入数据已经存在，需要替换时释放前值
// 删除数据时，释放
void trie_set_free_func(void (*free_func)(void* ));
