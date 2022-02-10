package main

import (
	"fmt"
	"sync"
)

var (
	count int
	countGuard sync.Mutex
)

func GetCount() int {
	countGuard.Lock()
	//在函数退出时解除锁定
	defer countGuard.Unlock()
	return count
}



func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}


func main() {
	// 可以进行并发安全的设置
	SetCount(1)

	//可以进行并发安全的获取
	fmt.Println(GetCount())
}
