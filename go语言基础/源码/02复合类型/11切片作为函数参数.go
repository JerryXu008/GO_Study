package main

import "fmt"
//切片作为函数参数  返回值
func demo(s []int) []int {
	//fmt.Println(s)
	//s[0] = 123
	//fmt.Println(s)

	//如果函数中使用append进行数据添加  形参地址发生改变  就会指向实参的地址
	s = append(s,6,7,8,9)
	fmt.Println(s)
	return s
}

func main() {
	//切片名本身就是个地址
	slice := []int{1,2,3,4,5}
	//地址传递  形参可以改变实参
	slice = demo(slice)
	fmt.Println(slice)
}
