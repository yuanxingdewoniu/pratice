package main

import "fmt"

func f(a [3]int)   { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a { // derefencing *a to get back to array is not necessary
		sum += v
	}
	return
}

func main() {
	var ar [3]int
	f(ar)
	fp(&ar)
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array)
	fmt.Printf("The sum of  the array is %f", x)
}
