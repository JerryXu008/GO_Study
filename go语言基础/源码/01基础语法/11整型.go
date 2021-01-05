package main

import (
	"fmt"
	"unsafe"
)

func main() {


	//uint无符号整型  int有符号整型
	//var a uint = -10
	//fmt.Println(a)


	b := 20
	fmt.Println(b)
	//%T打印变量所属类型
	fmt.Printf("%T\n",b)
	fmt.Println(unsafe.Sizeof(b))

	var c int32 = 10

	fmt.Printf("%T\n",c)
	fmt.Println(unsafe.Sizeof(c))


}
