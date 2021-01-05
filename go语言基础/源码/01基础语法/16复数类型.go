package main

import "fmt"

func main() {

	var t complex64
	t = 2.1 + 3.14i
	fmt.Println(t)

	//自动推导类型
	//默认为complex128
	t1 := 3.3 + 4.4i
	fmt.Printf("%T\n",t1)

	//通过内建函数  取实部   虚部
	fmt.Println(real(t1),imag(t1))

}
