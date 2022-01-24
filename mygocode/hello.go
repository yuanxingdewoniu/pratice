// hello_world1.go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 3
	y := 4
	println("  \n", x+y)
	fmt.Printf("%s\n", runtime.Version())
	println("Hello world")
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
}
