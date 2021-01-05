package main

import "fmt"

func main() {
	//给int64起一个别名 bigint
	type bigint int64
	var a bigint  // 等价于  var a int64
	a = 10
	fmt.Println(a)
	fmt.Printf("%T\n",a)

	type (
		long int64
		char byte
	)

	var b long = 11
	var ch char = 'a'
	fmt.Println(b,ch)
	fmt.Printf("%T,%T\n",b,ch)


}
