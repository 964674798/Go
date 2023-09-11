package main //声明当前文件所属的包，每个go文件必须有归属的包
import (
	"fmt"
	"unsafe"
	)// 引入程序中需要用的包，为了使用包下的函数

func main()  { // main  主函数 程序中的入口
	var a1 = "2323"
	fmt.Printf("这是变量类型: %T",a1)
	fmt.Println() //换行
	fmt.Println(unsafe.Sizeof(a1))

	// 数据采用最小原则
	// 表示学生的年龄
	var a2 uint8 = 23
	fmt.Printf("%T",a2)

	// 浮点类型 float32：4  float64：8  和操作系统无关
	// 存储分为 符号位 指数位 尾数位   会存在精度损失
	fmt.Println()

	var f1 float32 = 3.14
	fmt.Println(f1)

	var f2 float32 = -3.14
	fmt.Println(f2)

	// 科学技术法
	var f3 float32 = 314E-2
	fmt.Println(f3)

	var f4 float32 = 314E+2
	fmt.Println(f4)

	var f5 float32 = 314e+2
	fmt.Println(f5)

	var f6 float64 = 314e+2
	fmt.Println(f6)

	// 浮点数存在精度丢失的问题 float64表示的会比float32多 。开发中推荐使用的是float64
	var f7 float32 = 3.14159265748
	fmt.Println(f7)
	var f8 float64 = 3.14159265748
	fmt.Println(f8)

	var f9 = 3.14159265748
	fmt.Printf("float推断值的默认类型: %T",f9)



}