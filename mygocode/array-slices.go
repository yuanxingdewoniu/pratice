package main

import "fmt"

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func main() {

	var arr = [5]int{0, 1, 2, 3, 4}
	sum(arr[:])

	var arr1 [6]int
	var slice1 []int = arr1[2:5]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d \n", i, slice1[1])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	//grow the slice
	slice1 = slice1[2:4]
	println(len(slice1))
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d \n", i, slice1[i])
	}
	var slice2 []int = make([]int, 50, 100)
	for i := 0; i < len(slice2); i++ {
		slice2[i] = i * i

	}
	fmt.Println(slice1)

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

}
