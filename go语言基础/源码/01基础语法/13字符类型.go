package main

import "fmt"

func main() {
	//byte字符类型  同时也uint8的别名
	//1.单引号 字符 双引号  字符串
	var a byte = 'a'
	fmt.Println(a)
	//%c是一个占位符  表示打印输出一个字符
	fmt.Printf("%c\n",97)
	fmt.Printf("%T\n",a)

	var b byte = '0'
	fmt.Println(b)


	//\n \t  \0 字符串的结束的标志

	fmt.Println('\n')
	
}
