package main

import (
	"fmt"
)

func printer(c chan string) {
	//开始无限循环等待数据
	for {
		// 从channel 中获取一个数据
		data := <-c
		//将 0 视为数据结束
		if data == "" {
			break
		}

		//打印数据
		fmt.Println(data)
	}

	//通知main结束循环 （没有数据了）
	c <- ""
}

func main() {
	// 创建一个channel
	c := make(chan string)

	// 并发执行printer, 传入channel
	go printer(c)

	for i := 0; i < 4; i++ {
		c <- "I am working"
	}

	// 通知并发的printer结束循环（没有数据了）
	c <- ""

	//等待printer 结束
	<-c
}
