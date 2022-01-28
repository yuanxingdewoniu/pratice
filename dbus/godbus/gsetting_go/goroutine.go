package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f(from string)  {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":",i)
	}
}

func say(s string) {
	for i :=0; i < 3; i++  {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main(){
	f("direct")
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going  test 123")

	time.Sleep(time.Second)
	fmt.Println("done")

	go say("hello world")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("over")

	runtime.GOMAXPROCS(3)
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start Goroutines")

	go func() {
		defer wg.Done()

		for count :=0; count < 15;  count++ {
			for char := 'a'; char < 'a'+26; char++ {
				//fmt.Println(char)
				fmt.Printf("%c", char)

			}
			fmt.Println()
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count  < 15; count ++ {
			for char := 'A'; char < 'A' + 26; char ++ {
				fmt.Printf("%c", char)
			}
			fmt.Println()
		}
	}()


	fmt.Println("Waiting to Finish")
	wg.Wait()
	fmt.Println("Terminating Program")

}