package main

import "fmt"

func main0701() {
	//数组长度固定
	var arr [5]int = [5]int{1,2,3,4,5}
	fmt.Println(arr)
	//函数参数传递


}

func main0702()  {

	//切片
	//定义
	//第一种定义方式
	//var slice []int   //空切片
	//slice = append(slice, 1,2,3)
	//fmt.Println(slice)
	//fmt.Println(len(slice))  //长度
	//fmt.Println(cap(slice))  //容量
	//
	////第二种
	//var s1 []int = []int{1,2,3}
	//fmt.Println(s1)
	//
	////自动推导
	//s2 := []int{1,2,3,4,5}
	//fmt.Println(s2)

	//第三种

	//长度 小于 容量
	//省略容量
	//s := make([]int,10,5)
	s := make([]int,5)
	fmt.Println(s)
	s = append(s, 1,2,3,4,5,6,7,8)
	fmt.Println(s)

}

func main()  {
	//赋值

	//下标赋值
	//s := make([]int,5,10)
	//s[0] = 1
	//s[1] = 2
	//fmt.Println(s)
	//循环赋值

	s1 := make([]int,5,10)
	for i:=0;i<len(s1) ;i++  {
		s1[i] = i
	}
	fmt.Println(s1)

	for _,v := range s1 {
		fmt.Println(v)
	}

}