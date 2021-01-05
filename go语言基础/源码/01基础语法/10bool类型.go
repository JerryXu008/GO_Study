package main

import "fmt"

func main() {
	//1.声明变量 没有初始化  零值 为false
	var a bool
	fmt.Println(a)

	a = false
	fmt.Println(a)
	//2.布尔类型不接受其他类型的赋值，不支持自动或强制的类型转换
	//a = 1
	//a = bool(1)
	//fmt.Println(a)

	//3.自动推导类型

	var b = true
	fmt.Println(b)

	c := false
	fmt.Println(c)


	v2 := (1==2)
	fmt.Println(v2)
	fmt.Printf("%T",v2)

}
