package main

import (
	"fmt"
	"reflect"
)

func main()  {

	// *int 的空指针
	var a *int
	fmt.Println("var a *int", reflect.ValueOf(a).IsNil())

	// nil 值
	fmt.Println("nil:", reflect.ValueOf(nil).IsValid())

	// * int 类型的空指针
	fmt.Println("()")


}
