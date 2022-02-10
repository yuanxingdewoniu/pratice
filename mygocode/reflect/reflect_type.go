package main

import (
	"fmt"
	"reflect"
)

type Enum int

const (
	Zero Enum = 0
)

func main() {
	var a int
	typeOfA := reflect.TypeOf(a)

	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	type cat struct {
	}

	typeOfCat := reflect.TypeOf(cat{})
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	typeOfA2 := reflect.TypeOf(Zero)

	fmt.Println(typeOfA2.Name(), typeOfA2.Kind())

}
