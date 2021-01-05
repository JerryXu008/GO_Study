package main

import "fmt"

//函数可以返回多个值
func test() (a,b,c int) {

	return 1,2,3

}

func main() {

	//a := 10
	//b := 20
	//c := 30
	//多重赋值
	a, b := 10, 20
	fmt.Println(a, b)

	//a = b
	//b = a
	//fmt.Println(a, b)
	var c int
	c = a
	a = b
	b = c
	fmt.Println(a, b)

	//i := 10
	//j := 20
	i, j := 10, 20
	//多重赋值
	i, j = j, i
	fmt.Println(i, j)

	//_匿名变量  丢弃数据不处理

	tmp,_:=7,8
	fmt.Println(tmp)

	//var d,e,f int

	_,e,f := test()
	fmt.Println(e,f)


}


/*
总结：
一 多重赋值   一次对多个变量进行赋值
	例如 a,b,c = 1,2,3

二 匿名变量
	"_"单个下划线命名的变量就是匿名变量
	特点：丢弃数据不处理
	例如：
	tmp,_ := 10,20
	注意：
	1.不能打印_匿名变量
	2.匿名变量配合函数返回值才有优势


 */