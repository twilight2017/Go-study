package main

//导入其他包
import "fmt"

//函数外只能声明变量、常量、函数、类型
//Go语言中的变量必须先声明再使用
// var name string
// var age int
// var isOk bool

//全局变量声明
// var (
// 	name string
// 	age  int
// 	isOk bool
// )

//定义了常量之后不能修改
//常量是在程序运行期间不会改变的量
const pi = 3.1415926

//批量声明常量
// const (
// 	statusOK = 200
// 	notFound = 404
// )

// //三个值都将被声明为100
// const (
// 	n1 = 100
// 	n2
// 	n3
// )

//iota 实现类似枚举的功能
const (
	n1 = iota //0
	n2 = iota //1
	n3 = iota //2
)

const (
	b1 = iota //0
	b2        //1
	_         //2
	b3        //3
)

const (
	c1 = iota
	c2 = 100
	c3
	c4
)

const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) //1左移10位是2^10
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

//程序的入口
func main() {
	// name = "liuzixuan"
	// age = 18
	// isOk = true
	// //GO语言中变量声明了必须使用
	// fmt.Printf("name:%s", name) //%s:占位符，使用name这个变量的值去替换换行符
	// fmt.Println(age)            //打印完内容后会加上一个换行符
	// fmt.Print(isOk)             //直接打印内容

	// //声明变量的同时进行赋值
	// var s1 string = "ghjh"
	// fmt.Print(s1)

	// var test = 20
	// fmt.Print(test)

	// //简短变量声明，只能在函数中声明
	// s3 := "gdhiahsud"
	// fmt.Print(s3)
	// fmt.Println("n1", n1)
	// fmt.Println("n2", n2)
	// fmt.Println("n3", n3)
	// fmt.Println("b1", b1)
	// fmt.Println("b2", b2)
	// fmt.Println("b3", b3)
	// fmt.Println("c1", c1)
	// fmt.Println("c2", c2)
	// fmt.Println("c3", c3)
	// fmt.Println("c4", c4)
	// fmt.Println("d1", d1)
	// fmt.Println("d2", d2)
	// fmt.Println("d3", d3)
	// fmt.Println("d4", d4)
	var i1 = 101
	fmt.Printf("%d", i1)
	fmt.Printf("%T", i1)
	fmt.Printf("%b\n", i1)
	fmt.Printf("%o\n", i1)
	fmt.Printf("%x\n", i1)
	//八进制
	i2 := 077
	fmt.Printf("%d", i2)
	fmt.Printf("%T", i2)
	//十六进制
	i3 := 0x1234567
	fmt.Printf("%d", i3)
	fmt.Printf("%T", i3)

}
