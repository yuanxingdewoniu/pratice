// 访问结构体的值


package main

import (
	"fmt"
	"reflect"
)

type dummy struct {
	a int
	b string

	float32
	bool
	next *dummy
}


func main() {

	//值包装结构体
	d :=  reflect.ValueOf(dummy {
		next:&dummy{},
	})

	// 获取字段数量
	fmt.Println("NumFiled", d.NumField())

	// 获取索引为2的字段（float32 字段）
	floatFile := d.Field(2)

	//输出字段类型
	fmt.Println("Filed",floatFile.Type())

	// 根据名字查找字段
	fmt.Println("FiledByName(\"b\").Type", d.FieldByName("b").Type())

	// 根据索引查找值中， next字段的int 字段的值
	fmt.Println("FiledByIndex([]int {4,0}).Type", d.FieldByIndex([] int{4,0}).Type())


}