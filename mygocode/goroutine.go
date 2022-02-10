package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick ", times)
		N
		// 演示1秒
		time.Sleep(time.Second)
	}
}

func main() {

	// 并发执行程序
	go running()
	time.Sleep(time.Second * 2)
}
