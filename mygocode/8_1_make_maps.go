package main

import "fmt"

func main() {
	var maplist map[string]int
	var mapAssigned map[string]int

	maplist = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)

	mapAssigned = maplist

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159

	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is :%d \n", maplist["one"])
	fmt.Printf("Map created at \"key2\" is : %f \n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n ", maplist["two"])
	fmt.Printf("Map literal at \"ten\" is: %d \n", maplist["ten"])

	v := maplist["one"]

	fmt.Printf("test v information  \n v = %d ", v)

}
