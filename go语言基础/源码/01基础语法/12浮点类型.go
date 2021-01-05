package main

import "fmt"

func main() {
	//浮点型数据分为单精度 float32和双精度 float 64
	var a float64 = 123.45655555
	fmt.Println(a)
	var b float32 = 123.45655555
	fmt.Println(b)
	//通过自动推导类型创建的浮点型变量默认为float64
	c := 3.14
	fmt.Println(c)
	fmt.Printf("%T",c)

}
