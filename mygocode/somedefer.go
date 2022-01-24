package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}

	f()
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

}
