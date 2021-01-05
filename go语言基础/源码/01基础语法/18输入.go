package main

import "fmt"

func main() {

	var a float64
	fmt.Printf("请用户输入数据：")
	//阻塞 等待用户的输入
	//格式化的输入
	//fmt.Scanf("%f",&a)

	//更简单

	fmt.Scan(&a)

	fmt.Println(a)

}
