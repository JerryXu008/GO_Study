package main

import "fmt"
//func 函数名(函数参数列表)(函数返回值类型)
func Test(a,b int) (c,d,sum int) {
	//var a int = 5
	//var b int = 1
	//var sum int
	//c,d = 1,2
	c = 1
	d = 2
	sum = a + b
	//return 表示函数的结束  如果函数有返回值return 可以将返回值返回
	return

}

func main() {
	//var result int
	//var result1 int
	var result2 int
	//函数有多个返回值 要一一对应接收数据
	//result,result1,result2 = Test(5,1)
	//匿名变量   丢弃接受数据
	_,_,result2 = Test(5,1)
	fmt.Println(result2)
	var sum int
	fmt.Println(sum)
}
