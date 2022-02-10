package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

func main() {
	fmt.Printf(" %f \n", math.Pi)
	fmt.Printf("%2.4f", math.Pi)

	// 图片大小
	const size = 300

	// 根据给定大小创建灰色图
	pic := image.NewGray(image.Rect(0, 0, size, size))

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})

		}
	}

}
