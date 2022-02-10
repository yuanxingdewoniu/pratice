package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	var str string = "Hi I'm Marc, Hi."
	fmt.Println("The position of \"Marc\" is :", str)

	s := t.Format("20010102")
	fmt.Println(t, "=>", s)

}
