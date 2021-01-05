package main

import "fmt"

//定义函数			//形参
func test10(a int, b string) {
	fmt.Printf("a=%d,b=%s\n", a, b)
}



func main() {

	//函数调用		//参数的传递
	test10(1, "H") //  实参
	test10(3, "E")

}
