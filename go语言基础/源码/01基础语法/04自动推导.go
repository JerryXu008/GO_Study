package main

import "fmt"

func main(){
	//自动推导
	var a = 10
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	//常用的自动推导
	//自动推导  := 左边变量名没使用过
	b := 10
	b = 20
	b = 30
	fmt.Println(b)

	c :=3.14
	fmt.Println(c)
	fmt.Printf("%T\n",c)

	d,f,e := 20,3.14,30
	fmt.Println(d,f,e)
	fmt.Printf("%T",e)
}