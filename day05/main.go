package main

import "fmt"

//切片slice
func main() {
	var s []int
	var s2 []string
	s = []int{1, 2, 3}
	s2 = []string{"Beijing", "Shanghai", "Shenzheng"}
	fmt.Println(s, s2)
	//切片的容量指底层数组的容量
	//底层数组从切片的第一个元素指向最后一个元素
	fmt.Println("len:%d,cap:%d", len(s), cap(s))
	fmt.Println("len:%d,cap:%d", len(s2), cap(s2))

	//由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] //基于数组进行切割，左包含右不包含
	fmt.Println(s3)
	s3[0] = 9
	fmt.Println(s3, a1)

	//用make()函数申明切片
	a2 := make([]int, 5, 10)
	fmt.Println("len=%d, cap=%d", len(a2), cap(a2))
}
