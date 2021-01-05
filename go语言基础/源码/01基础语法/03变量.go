package main

import "fmt"

func main() {

	//一变量的声明
	// 1. 声明var 变量名  类型   变量声明之后，必须使用
	// 2. 只是声明变量没有初始化，默认为0
	// 3. 在同一个{}里，声明变量是唯一的

	var a int
	a = 10   //变量的赋值   先声明  再赋值

	fmt.Println(a)

	//4.可以同时声明多个变量
	//var b, c int
	//
	//b,c = 20,30
	//fmt.Println(b,c)


	//二变量初始化  声明变量的同时进行赋值

	var b int = 20  //初始化

	b = 30
	fmt.Println(&b)

	var value float64 =2
	fmt.Println(value)

	//=后可以跟表达式

	var sum float64 = 1*9
	fmt.Println(sum)



}
