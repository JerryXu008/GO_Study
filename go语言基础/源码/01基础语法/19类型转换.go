package main

import "fmt"

func main() {

	a, b,c := 54,30,20

	sum:= a+b+c
	//类型转换 用来在不同但相互兼容的类型之间相互转换的方式
	//1.数据类型（变量）  int（a）
	//2.数据类型（表达式）
	fmt.Println(float64(sum)/3)
	//不兼容
	//var flag bool
	//flag = true
	//fmt.Println(int(flag))
	//
	//flag = bool(1)


	var d float32 = 3.1
	var f float64 = 3.5

	//在类型转换时建议低类型转成高类型  保证数据的精度
	//int8---int16---int32---int64
	//float32---float64
	//int64 ----float64()
	//建议整型转成浮点型
	num := float64(d)+f
	fmt.Println(num)

	//高类型----低类型  会丢失精度  数据溢出  符号发生改变

	var g int = 1234
	fmt.Println(int8(g))







}
