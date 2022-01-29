package main

import (
	"fmt"
	"time"
)

func main()  {
	// 构建一个通道
	ch := make(chan int)

	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知main的goroutine
		ch <-0
		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait Goroutine")

	// 等待匿名goroutine
	<- ch
	fmt.Println("all done")

	//构建一个通道
	ch1 :=make(chan int)

	// 开启一个并发匿名函数
	go func() {
		for i := 1	3; i >= 0; i--{
			// 发送3到0之间的数值
			ch1 <- i
			// 每次发送完时等待
			time.Sleep(time.Second)
		}
	}()

		for data := range ch1 {
			fmt.Println(data)
			if data == 0 {
				break
		}
	 }
   }

