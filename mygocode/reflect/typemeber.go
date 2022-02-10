package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 声明一个空结构体
	type cat struct {
		Name string
		Type int `json："type" id:"100"`
	}
	// 创建cat 实例
	ins := cat{Name: "mini", Type: 1}

	//获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)

	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i)
		//输出成员名和tag
		fmt.Printf("name : %v tag:'%v'\n", fieldType.Name, fieldType.Tag)
	}

	//通过字段名，找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Println(ok)
		// 从tag 中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}

	// 声明整形变量a并赋初始值
	var a int = 1024

	// 获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a)

	// 获取interface{}类型的值，通过类型断言转换
	var getA int = valueOfA.Interface().(int)

	// 获取64位的值，强制类型转换为int 类型
	var getA2 int = int(valueOfA.Int())

	fmt.Println(getA, getA2)

}
