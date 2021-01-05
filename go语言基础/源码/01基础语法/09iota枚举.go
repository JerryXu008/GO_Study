package main

import "fmt"

func main() {

	//1.iota常量自动生成器，每个一行，自动加1
	//2.iota给常量赋值使用
	const (
		a = iota //0
		b = iota //1
		c = iota //2
	)

	fmt.Println(a, b, c)

	//3.iota遇到const，重置为0
	const d = iota

	fmt.Println(d)
	//4.可以只写一个iota  在一个括号
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Println(a1,b1,c1)

	//5.如果在同一行 ，值一样
	const(
		i = iota  //0
		j1,j2,j3 = iota,iota,iota  // 1 1 1
		k =iota  //2

	)
	fmt.Println(i)
	fmt.Println(j1,j2,j3)
	fmt.Println(k)
}
