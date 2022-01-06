/*
 * Copyright (C) 2020 ~ 2021 统信软件技术有限公司
 *
 * Author:  Aaron <zhangya@uniontech.com>
 *
 * Maintainer: Aaron <zhangya@uniontech.com>
 *
 */

#include "../include/module/trie_tree.h"
#include <linux/limits.h>
#include "../include/common/dlp_macro.h"

#ifndef max
#define max(a,b) (((a) > (b)) ? (a) : (b))
#define min(a,b) (((a) < (b)) ? (a) : (b))
#endif

static void (*trie_free_func)(void* ) = NULL;

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

void node_print(TireNode* tree, int nest)
{
        if (tree) {
                for (int i = 0; i < nest; ++i) {
                        DLP_PRINT("--");
                }
                DLP_PRINT("|%s\n", tree->node_name);
                if (tree->data) {
                        for (int i = 0; i < nest; ++i) {
                                DLP_PRINT("  ");
                        }
                }
                trie_print(tree->data);
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

void trie_free(TireNode* tree)
{
        if (tree == NULL) {
                return;
        }

        if (tree->childs && tree->child_count > 0) {
                for (int i = 0; i < tree->child_count; ++i) {
                        trie_free(*(tree->childs + i));
                        DLP_FREE(*(tree->childs + i));
                }

                DLP_FREE(tree->childs);
                tree->childs = NULL;
        }

        if (tree->data) {
                trie_free_func(tree->data);
                DLP_FREE(tree->data);
                tree->data = NULL;
        }
}

TireNode* tire_node_create(const char* node_name, unsigned short node_len,
                           void* data)
{
        TireNode* node = (TireNode* ) DLP_MALLOC(sizeof(TireNode));

        node->node_name = (char* ) DLP_MALLOC(node_len + 1);

        memcpy(node->node_name, node_name, node_len);
        node->node_name[node_len] = '\0';

        node->node_len = node_len;

        node->data = data;
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

void* trie_get_child(TireNode* tree, const char* path,
                          unsigned short path_len)
{
        //DLP_PRINT("trie_get_child path = %s tree = %p", path, tree);
        if (0 == strncmp(tree->node_name, path, PATH_MAX)) {
                return tree->data;
        }

        int position;
        int res = trie_search(tree, path, path_len, &position);
        if (0 != res) {
                DLP_PRINT("trie_get_child res != 0");
                return NULL;
        }

        TireNode* child = *(tree->childs + position);
        return child->data;
}

TireNode* trie_insert_child(TireNode* tree, const char* path,
                                 unsigned short path_len, void* data)
{
        TireNode* child = NULL;
        int position;
        TireNode** temp;

        int res = trie_search(tree, path, path_len, &position);

        //已经存在的数据直接替换
        if(res == 0) {
                child = *(tree->childs + position);
                if (data) {
                        trie_free_func(child->data);
                        child->data = data;
                }

                return child;
        }

        child = tire_node_create(path, path_len, data);

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

int trie_insert(TireNode* tree, const char* path, unsigned short path_len,
                         void* data)
{
        int get_node_name = 0;
        unsigned short word_size = 0;
        const char* node_start = 0;
        unsigned short i = 0;
        TireNode* node;

        //DLP_PRINT("trie_insert path = %s data = %p\n", path, data);

        if (tree == NULL) {
                return -1;
        }

        if (path == NULL) {
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
                        node = trie_insert_child(tree, node_start, word_size,
                                                   NULL);
                        return trie_insert(node, path + i, path_len - i,
                                                data);
                }

                word_size++;
        }

        if (i == path_len) {
                trie_insert_child(tree, node_start, word_size, data);
                return 0;
        }

        return -1;
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
                        // DLP_PRINT
                        //         ("trie_bsearch child_count = %d path = %s ,path_len = %d position = %d",
                        //          child_count, path, path_len, *position);
                        *position = __idx;
                        return 0;
                }
        }

        if (__comparison > 0) {
                *position = __idx + 1;
                // DLP_PRINT
                //         ("trie_bsearch child_count = %d path = %s ,path_len = %d position = %d",
                //          child_count, path, path_len, *position);
        } else {
                *position = __idx;
                // DLP_PRINT
                //         ("trie_bsearch child_count = %d path = %s ,path_len = %d position = %d",
                //          child_count, path, path_len, *position);
        }

        return -1;
}

int trie_search(TireNode* tree, const char* path, unsigned short path_len,
                         int* position)
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
        void* result = tree->data;
        TireNode* child;
        void* pd;
        int res;

        //DLP_PRINT("trie_get path = %s, path_len = %d tree = %p \n", path, path_len, tree);

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
                                return result;
                        }

                        child = *(tree->childs + position);
                        pd = trie_get(child, path + i, path_len - i);
                        if (pd) {
                                return pd;
                        }

                        return result;
                }

                word_size++;
        }

        if (i == path_len) {
                pd = trie_get_child(tree, node_start, word_size);
                if (pd) {
                        result = pd;
                }
        }

        return result;
}

void trie_set_free_func(void (*free_func)(void* ))
{
        trie_free_func = free_func;
}

