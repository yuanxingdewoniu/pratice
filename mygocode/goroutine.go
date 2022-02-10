package main

import (
	"fmt"
	"time"
)

func goroutine1(t int) {

	for t > 0 {
		t--
		fmt.Println("Hello goroutine ", t)
	}

}

func main() {

	go goroutine1(4)

	go func() {
		var times int
		for {
			times++
			fmt.Println("num = ", times)
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("Hello main")
}
