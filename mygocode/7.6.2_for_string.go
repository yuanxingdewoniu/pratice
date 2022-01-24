package main

import "fmt"

func main() {
	 s:= "\u02ff\u752c"
	for i, c := range s {
		fmt.Printf("%d :%c \n", i, c)
	}
	s2 := "hello"
	c := []byte(s2)
	c[0] = 'c'O
	s3 := string(c)
	println(s3)

}
