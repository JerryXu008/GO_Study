package main

import "fmt"

func main() {

	//变量声明 var

	//常量声明 const

	const a float32 = 10
	//a = 20  //err 常量不允许修改
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	//自动推导类型
	const  b,c  = 3.14,3.2  // 没有使用:=
	fmt.Println(b,c)
	fmt.Printf("%T",b)


	
}
