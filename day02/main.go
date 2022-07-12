package main

import (
	"fmt"
	"strings"
)

//整形
func main() {
	// //十进制
	// var i1 = 101
	// fmt.Printf("%d\n", i1)
	// fmt.Printf("%o\n", i1) //把十进制数转化为八进制数
	// //八进制
	// i2 := 077
	// fmt.Printf("%d\n", i2)
	// //十六进制
	// i3 := 0x1234567
	// fmt.Printf("%T\n", i3)
	//浮点数
	//math.MaxFloat32 //最大浮点数
	// f1 := 1.23456
	// fmt.Printf("%T/n", f1) //默认Go语言中的小数都是float64类型
	//fmt study
	//查看类型
	// var n = 100
	// fmt.Printf("%T", n)
	// fmt.Printf("%#v", n)
	// fmt.Printf("%d", n)
	// fmt.Printf("%b", n)
	// fmt.Printf("%o", n)
	// fmt.Printf("%x", n)
	//字符串
	s := "D:\\Go\\Go-study"
	ret := strings.Split(s, "\\")
	fmt.Printf("%T", ret)
	fmt.Println(ret)
	fmt.Println(strings.Contains(s, "D:"))
	fmt.Println(strings.HasPrefix(s, "D:"))
	fmt.Println(strings.Index(s, "study"))
	//拼接字符串
	fmt.Println(strings.Join(ret, "+"))
}
