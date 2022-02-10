package main

import "fmt"

type DataWriter interface {
	WriteData(data interface{}) error
}

type file struct {
}

func (d *file) WriteData(data interface{}) error {
	//模拟写入数据
	fmt.Println("WriteData :", data)
	return nil
}

func main() {

	// 实例化 file
	f := new(file)

	// 声明一个DataWriter的接口
	var writer DataWriter

	// 将接口赋值f， 也就是file 类型
	writer = f

	writer.WriteData("date test 2022-02-08")

}


