package main

import "fmt"

// 首字母大写可以被其他包使用，首字母小写不能被其他包使用。命名的时候需要使用驼峰
func calc (num1 int,num2 int) (int) { //返回值就一个的时候（）是可以省略不写
	var sum int = 0 
	sum += num1
	sum += num2
	return sum
}

// 函数可以有多个返回值
// 函数有多个返回值的情况，如果不接受可以用返回值_代替
func MixCalc (num1 int, num2 int)(int,int,int){
	var sum int = 0 
	sum += num1
	sum += num2

	sub := num1 - num2
	return sum,sum,sub
}

// 内存的分析、 堆，栈，执行空间  

// 支持可变参数数量
func moreParams (args...int)(){
	for i:=0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}
// 不支持重复声明

// 数组在函数中是按照基础类型计算的，不是按照引用传递


// 函数内改变函数外的方法。传入变量地址&   函数内通过指针*的方式操作变量。从效果来看，类似引用传递
func modifyParams (arg1 *int){
	*arg1 = 23
}

// 类似js  函数也是一种类型
func typeFunction (){}

// 类似js  函数也可以作为一个形参

// GO支持自定义数据类型  类似ts ，新自定类型和原来类型不一致 、 给一个函数起别名会场方便
type myInt int

// 支持对返回值进行命名
func MixCalcA (num1 int, num2 int)(sum int, sub int){
	sum += num1
	sum += num2

	sub = num1 - num2
	return 
}


func main()  {
	// if 语句
	// 1. 可以省略()  2.if后面可以同步声明变量
	
	// sum := calc(20,30)
	// fmt.Println("求和方法的值",sum)

	var result1,result2,_ = MixCalc(10,8)
	fmt.Println(result1,result2)
	var _,_,result3 = MixCalc(5,8)
	fmt.Println(result3)

	moreParams(12,13,14,15,17)

	result4 := 20
	// &result4 表示地址
	modifyParams(&result4)
	fmt.Println(result4)

	// 类似js  函数也是一种类型
	a := typeFunction
	fmt.Printf("%T,%T",typeFunction,a) //func(),func()
	fmt.Println("-----------------------------")

	test1,test2 := MixCalcA(1,2)
	fmt.Println(test1,test2)

}