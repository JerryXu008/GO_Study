package main

import "fmt"

func test1(num1,num2 int)  {
	fmt.Println(num1+num2)
}

func test(a,b int)  {
	test1(a,b)
}

func test3(arr...int)  {
	fmt.Println(arr)
}


func test2(arr...int)  {

	//不能将不定参的名称传递给另一个不定参
	//test3(arr...int)  //error

	//传递指定个数数据
	//test3(arr[0],arr[1],arr[2],arr[3])
	//传递多个数据
	test3(arr[0:5]...)

	//传递全部参数
	//test3(arr[:]...)
	//test3(arr...)

}

func main() {

	//1.什么是函数的嵌套
		//在一个函数中调用另一个函数
	//2.函数嵌套的执行过程
	//test(10,20)
	//3.函数嵌套调用的应用


	//4.不定参数函数调用
	test2(1,2,3,4,5,6)

}
