package main

import "fmt"

func main() {

	//一.语法:
	//for 初始化条件;判断条件 ; 条件变化 {
	//	代码体
	//
	//}
	//二.计算1+2+3+...+10的和
	//1.初始化条件 i:=1
	//2.判断条件是否满足  满足则执行循环体  不满足  跳出循环
	//3.条件变化  i++
	//4.重复 2， 3 ，4
	//sum := 0
	//for i:=1;i<11 ; i++ {
	//	sum=sum+i
	//	//fmt.Println(sum)
	//}
	//
	//fmt.Println(sum)


	//三.range
	//通过for打印每个字符
	str := "abc"
	//for i:=0;i<len(str) ;i++  {
	//	fmt.Println(str[i])
	//	fmt.Printf("str[%d]=%c\n",i,str[i])
	//}

	//迭代打印每个元素  默认返回两个值 第一个元素位置  一个元素本身

	for i ,data :=range str{
		fmt.Println(i,data)

	}

	for _,data:=range str{
		fmt.Println(data)

	}

}
