package main

import (
	"fmt"
	"time"
)

func main() {

	num := []int{1, 3, 5}
	//遍历数组、切片 获取索引和元素
	for k1, v1 := range num {
		fmt.Printf("key:= %d, value:= %d\n", k1, v1)
	}

	//遍历字符串
	var str = "hi"

	for sk, sv := range str {
		fmt.Printf("key: %d value: 0x%x\n", sk, sv)
	}

	//遍历map 获取map的键和值

	m := map[string]int{
		"a": 100,
		"b": 200,
	}
	for mk, mv := range m {
		fmt.Printf("map key = %s , map value = %d \n", mk, mv)
	}

	//遍历通道
	c := make(chan int)

	go func() {
		c <- 1
		c <- 3
		close(c)
	}()

	time.Sleep(2 * time.Second)
	for cv := range c {
		fmt.Println(cv)
	}

}
N