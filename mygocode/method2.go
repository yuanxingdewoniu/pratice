package main

import (
	"fmt"
	"runtime"
)

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}
func main() {
	fmt.Println(IntVector{1, 2, 3}.Sum())
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)


}
