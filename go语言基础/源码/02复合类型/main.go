package main

//同级目录  包名一致
import (
	"fmt"
	"userinfo"
)

//go build main.go test.go
func main() {
	fmt.Println("main")
	test()

	//包名.函数名  (函数名首字母大写)
	userinfo.Userlogin()

	arr := [3]int{1,2,3}
	fmt.Println(arr)



}
