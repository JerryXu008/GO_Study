package main

import "fmt"

func main() {

	//变量 程序可运行期间 可以改变的量  声明关键字var
	//var a,b int = 10,20
	//fmt.Println(a,b)

	//快捷键  ctrl+/ 注释 解注释

	//不同类型变量的定义
	//var a int =1
	//var b float64 = 2.0
	//var (
	//	a int =1
	//	b float64 = 2.0
	//)

	//自动推导类型
	//var (
	//	a = 1
	//	b = 2.0
	//)

	a,b := 1,2.0
	fmt.Println(a,b)


	//常量 ：程序运行期间 不可以更改量   声明关键字 const
	//常量的定义

	//const i int = 10
	//const j float64 = 3.14
	//fmt.Println(i,j)

	//const (
	//	i int = 10
	//	j float64 = 3.14
	//)



	//自动推导

	const (
		i = 10
		j = 3.14
	)
	fmt.Println(i,j)
	fmt.Printf("%T",j)
	
}
