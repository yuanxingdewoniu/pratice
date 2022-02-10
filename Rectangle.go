package main

import (
	"fmt"
)

//定义接口
type Rectfunc interface {
	area() float64
	permater() float64
}

//定义长方形参数结构体
type Rectangle struct {
	height float64
	wight  float64
}

// 定义正方形参数的结构体
type Square struct {
	length float64
}

// 长方形面积的接口的具体实现
func (R Rectangle) area() float64 {
	return R.wight * R.height
}

// 长方形周长的具体实现
func (R Rectangle) permater() float64 {
	return 2 * R.wight * R.height
}

// 正方形面积的具体实现
func (S *Square) area() float64 {
	return S.length * S.length
}

// 正方形的周长计算公式
func (S *Square) permater() float64 {
	return 4 * S.length
}

//调用相关的方法
func main() {
	//var Rect = new(Rectangle)

	var Rect Rectangle
	Rect = Rectangle{4, 6}
	fmt.Println("长方形的面积 = ", Rect.area())
	fmt.Println("长方形的周长 = ", Rect.permater())

	var s1 Square
	s1 = Square{length: 4}
	fmt.Println("正方形S1的 面积 =:", s1.area())
	fmt.Println("正方形S1的 周长= :", s1.permater())

	var s2 = new(Square)
	*s2 = Square{length: 40}
	fmt.Println(" 正方形S2 的面积 =:", s2.area())
	fmt.Println(" 正方形S2 的周长 = :", s2.permater())

}
