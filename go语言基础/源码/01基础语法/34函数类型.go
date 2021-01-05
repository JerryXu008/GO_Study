package main

import "fmt"

func demo1(a,b int)  {
	fmt.Println(a+b)
}
//定义函数类型  为已存在的数据类型起别名
type FUNCDEMO func(int,int)

func main() {
	//demo1(10,20)
	//函数名表示一个地址 函数在代码区的地址

	fmt.Println(demo1)
	//自动推导类型
	//f:=demo1

	//f是func(int,int)函数类型定义的变量
	//var f func(int ,int)
	var f FUNCDEMO
	f = demo1
	fmt.Println(f)
	//通过f调用了函数
	f(10,20)  //  相当于 deme1(10,20)
	//f类型
	fmt.Printf("%T\n",f)

}
