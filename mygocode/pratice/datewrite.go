package main

import "fmt"

//定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{ }) error
}

//定义文件结构，用于实现DataWrite
type file struct {

}

// 实现DataWriter 接口的WriteData()方法
func (f *file) WriteData(data interface{}) error {
	//模拟写入数据
	fmt.Println("WriteData:",data)
	return nil
}

func main()  {
	//实例化file
	f := new(file)

	// 声明一个 DataWriterd的接口
	var writer DataWriter

	// 将接口赋值给f, 也就是*file的类型

	writer = f
	// 使用DataWriter接口进行数据的写入
	writer.WriteData("data")
}
