package main

import "fmt"

func swap(a, b *int) {

	t := *a
	*a = *b
	*b = t

}

func swap2(a *int, b *int) {

	t := *a
	*a = *b
	*b = t

}

func main() {
	x, y := 10, 20
	swap(&x, &y)

	fmt.Println(x, y)

	u, v := 101, 202
	swap2(&u, &v)
	fmt.Println(u, v)
}
