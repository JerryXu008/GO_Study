package main

import "fmt"

func test11() func()int{
	var x int

	return func() int {
		x++
		return x
	}
}
func main()  {

	//匿名函数  没有名字的函数
	//var num int
	//num = 9
	//f:=func(){
	//	num++
	//	fmt.Println(num)
	//} //花括号  {}() 函数调用
	//////函数调用
	//f()
	////fmt.Println(num)
	//
	//type FuncType func()
	//
	//var f1 FuncType
	//f1 = f
	//f1()
	//fmt.Println(num)




	//匿名函数参数传递
	//1.
	//func(a,b int){  //匿名函数带参数
	//	var sum int
	//	sum = a+b
	//	fmt.Println(sum)
	//}(3,6)  //调用时传值
	//2.
	//f:=func(a,b int){  //匿名函数带参数
	//	var sum int
	//	sum = a+b
	//	fmt.Println(sum)
	//} //调用时传值
	//f(3,6)


	//匿名函数返回值

	//x,y := func(i,j int)(max,min int){
	//	if i >j {
	//		max = i
	//		min = j
	//	}else {
	//		max = j
	//		min = i
	//	}
	//	return
	//}(10,20)
	//fmt.Println(x,y)


	//闭包
	f:=test11()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}