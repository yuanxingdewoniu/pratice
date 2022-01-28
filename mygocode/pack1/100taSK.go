package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)

	var cat  int = 1
	var str string = "banana"

	fmt.Println("%p %p",&cat, &str)


	var house = "Malibu Point 10880, 90265"
	ptr := &house

	fmt.Printf("ptr type : %t\n",ptr)

	fmt.Printf("address : %p \n", ptr)

	value := *ptr
	fmt.Printf("value type: %T \n",value)

	fmt.Printf("Value : %s\n", value)

}
