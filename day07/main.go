package main

import (
	"fmt"
	"time"
)

func main() {
	//pointer
	//&:取地址
	//*:根据地址取值
	// n := 18
	// fmt.Println(&n)
	// p := &n
	// fmt.Printf("%T\n", p)
	// fmt.Println(p, *p)
	// fmt.Println(*n)

	// //new函数申请内存地址
	// var a1 = new(int)
	// fmt.Printf("%T\n", a1)
	// //make只作用于slice, map以及chan的内存创建，且直接返回这三个类型，而不是它们的指针类型

	//map：提供映射关系的容器，无序的基于key-value的数据结构，其是引用类型,，必须进行初始化操作
	// var m1 map[string]int
	// m1 = make(map[string]int, 10)
	// m1["test"] = 1290
	// m1["test2"] = 679
	// fmt.Println(m1)
	// _, ok := m1["test5"]
	// if !ok {
	// 	fmt.Println("No key!")
	// }
	// //map的遍历
	// for k, v := range m1 {
	// 	fmt.Println(k, v)
	// }
	// delete(m1, "test")
	// fmt.Println(m1)
	// delete(m1, "test6")
	// fmt.Println(m1)
	//map sort
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
}
