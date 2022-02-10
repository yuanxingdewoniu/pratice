package main

import "fmt"

// 定义接口
type tank interface {
	// 方法
	Tarea() float64
	Volume() float64
	clean() string
}

//定义接口体
type myvalue struct {
	radius float64
	height float64
}

// tank interface
func (m myvalue) Tarea() float64 {
	return 2*m.radius*m.height + 2*3.14*m.height*m.radius
}

func (m myvalue) Volume() float64 {
	return 3.14 * m.radius * m.radius * m.height
}

func (m myvalue) clean() string {
	if m.radius > 200 {
		return "big tank"
	} else {
		return "small tank"
	}

}

func main() {
	var t tank
	t = myvalue{10, 14}
	fmt.Println("水桶的面积", t.Tarea())
	fmt.Println("水桶的容量", t.Volume())
	fmt.Printf("tank have some water : %s", t.clean())

}
