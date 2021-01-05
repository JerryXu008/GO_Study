package main

import "fmt"

func main() {

	var str1 string  //声明变量
	str1 = "abc"   //赋值
	fmt.Println(str1)

	//自动推导类型
	str2 := "知了课堂"
	fmt.Println(str2)
	fmt.Printf("%T\n",str2)

	//ch := 'a'
	//str := "a"   //'a''\0'  字符串结束标志

	//len函数  计算字符串中字符的个数  不包含\0
	//在go语言中 一个汉字占3个字符
	fmt.Println(len(str2))


	//字符串拼接  +

	str3 := "hello"
	str4 := "知了课堂"
	str5 := str4+str3
	fmt.Println(str5)

	
}
