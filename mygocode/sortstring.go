package main

import (
	"fmt"
	"sort"
)

// 将[] string 定义为 MyStringList
type MyStringList []string

//实现sort. Interface接口的获取元素数量方法
func (m MyStringList) Len() int  {

	fmt.Println(len(m))
	return len(m)
}

// 实现sort InInterface 接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	fmt.Println(m[i] < m[j])
	return  m[i] < m[j]
}

//实现 sort.Interface 接口的交换元素的方法
func (m MyStringList) Swap(i, j int) {

	fmt.Printf("m[i] = %d m[j] = %d",m[i], m[j])
	m[i], m[j] = m[j], m[i]
}


func main() {


	//准备一个元素打乱的字符串切片
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. QUadra Kill",
		"1. First Blood",
	}

    //fmt.Println(sort.Len(names))
	//sort.Less(names)
	//for _, t := range names {
	//	fmt.Println("%s\n",t)
	//}

	// 使用sort 包进行排序
	sort.Sort(names)

	for _, v := range names {
			fmt.Println("%s\n",v)
	}
}
