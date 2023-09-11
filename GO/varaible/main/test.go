package main //声明当前文件所属的包，每个go文件必须有归属的包
import "fmt" // 引入程序中需要用的包，为了使用包下的函数
// 全局变量
var ng  = 23
// 同时生成多个变量
var (
	ng1 = 23
	ng2 = "测试2"
)

func main()  { // main  主函数 程序中的入口
	// fmt.Println("Hello World!")
	// 局部变量
	// 1.变量的声明
	var num int
	// 2.变量的赋值
	num = 4
	num = 5
	// 3.变量的使用
	fmt.Println(num)

	// var num1 int = 2
	// var num2 int = 3
	// var total int = num1 + num2
	// fmt.Println(num2+num1)
	
	// 变量的使用方式
	// 1.直接使用
	var num2 int = 3
	fmt.Println(num2)

	// 2.不同的变量存在不同的默认值
	var num3 int
	fmt.Println(num3)

	// 3.变量存在类型推断
	var num4 = 28
	fmt.Println(num4)

	// 4.省略var  :=
	test := "内容测试"
	fmt.Println(test)

	fmt.Println("---------------------------------------------")

	// 声明多个变量
	var n1,n2,n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	var n4,n5,n6  = "我说","：",5
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)

	n7,n8 := 12.23,34.123
	fmt.Println(n7)
	fmt.Println(n8)

	fmt.Println(ng1)
	fmt.Println(ng2)
}