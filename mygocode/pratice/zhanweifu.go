package main

import "fmt"

func main() {

	// 这里我们使用range来计算一个切片的所有元素和
	// 这种方法对数组也适用
	nums := []int{2, 3, 4, 5, 6}
	sum := 0
	for t, num1 := range nums {
		sum += num1
		fmt.Printf("t = %d num = %d\n", t, num1)
	}
	fmt.Println("sum:", sum)

	nums2 := nums[:3]
	fmt.Println("nums2 = ", nums2)
	sum2 := 0
	for num2 := range nums2 {
		sum2 += num2
		fmt.Printf(" num2 = %d\n", num2)
	}
	fmt.Println("sum:", sum2)

	nums3 := nums[:3]
	fmt.Println("nums3 = ", nums3)
	sum3 := 0
	for _, num3 := range nums3 {
		sum3 += num3
		fmt.Printf("num3 = %d\n", num3)
	}
	fmt.Println("sum:", sum3)

	// range 用来遍历数组和切片的时候返回索引和元素值
	// 如果我们不要关心索引可以使用一个下划线(_)来忽略这个返回值
	// 当然我们有的时候也需要这个索引
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 使用range来遍历字典的时候，返回键值对。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k1, _ := range kvs {
		fmt.Printf("k1 -> %s\n", k1)
	}

	// range函数用来遍历字符串时，返回Unicode代码点。
	// 第一个返回值是每个字符的起始字节的索引，第二个是字符代码点，
	// 因为Go的字符串是由字节组成的，多个字节组成一个rune类型字符。
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
