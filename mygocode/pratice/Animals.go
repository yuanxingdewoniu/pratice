package main

import (
	"fmt"
	"time"
)

func goroutine1() {
	fmt.Println("Hello goroutine")
}

func main() {
	go goroutine1()
	time.Sleep(1 * time.Second) //避免main 方法 所在的 goroutine 就销毁了，其他的 goroutine 都没有机会执行完。
	fmt.Println("Hello main")
}
