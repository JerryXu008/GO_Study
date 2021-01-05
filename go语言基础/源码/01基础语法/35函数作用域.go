package main

import "fmt"

//全局变量
//1.定义在函数外部的变量  全局变量
//	全局变量作用域 整个文件
//2.全局变量名不能其他文件中的额全局变量名重名
	//全局变量名和局部变量名可以重名  尽量避免

var a int =20

func test8()  {
	a= 5
	a+=1
}

func demo6()  {
	fmt.Println(a)
}


func main02() {

	//作用域  作用范围
	//局部变量
	//1.定义在函数内部的变量  局部变量
	//	局部变量的作用域  函数内部有效
	//2.不同的函数  可以定义相同的局部变量名  互不影响
	//3.变量 先声明  在使用  在函数内部  变量名是唯一的
	//var i int = 10
	//for i:=0;i<5 ;i++  {
	//	fmt.Println(i)
	//}
	//fmt.Println(i)


	//如果全局变量名和局部变量名相同 那么会使用 局部变量 就近原则
	a:=9
	//test8()
	fmt.Println(a)

	//修改全局变量  会影响到其他位置使用全局变量
	demo6()

}

func main()  {
	//打印函数地址       代码区
	fmt.Println(demo6)
	//打印全局变量地址    数据区
	fmt.Println(&a)
	//打印局部变量地址   栈区
	a:=10
	fmt.Println(&a)
}