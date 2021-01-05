package main

import "fmt"

func demo1(m map[int]string)  {

	m[4] = "黄月英"
	fmt.Printf("%p\n",m)
	fmt.Println(m)
}


func main() {
	m := map[int]string{1:"貂蝉",2:"大乔",3:"小乔"}
	//地址传递  引用传递  形参和实参指向相同内存地址  修改形参  会影响实参的值
	demo1(m)
	fmt.Printf("%p\n",m)
	fmt.Println(m)


}
