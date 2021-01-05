package main

import "fmt"

func main0801() {

	//切片名本身就是一个地址
	//空切片 指向内存地址吧编号为0的空间
	var slice []int
	fmt.Printf("%p\n",slice)
	//当使用append进行追加数据时，切片地址可能会发生改变
	slice = append(slice,1,2,3)
	fmt.Printf("%p\n",slice)
	fmt.Printf("%p\n",&slice[1])

	slice = append(slice,4,5,6)
	fmt.Printf("%p\n",slice)
}

func main()  {
	//容量
	//s:=make([]int,5,8)
	//s = append(s, 1)
	//s = append(s, 2)
	//s = append(s, 3)
	//s = append(s, 4)
	////s[0] = 1
	//fmt.Println(s)
	//fmt.Printf("len = %d,cap = %d",len(s),cap(s))


	s:=make([]int,0,1)
	oldcap := cap(s)
	for i:=0;i<200000 ;i++  {
		s = append(s, i)
		newcap:=cap(s)
		if oldcap<cap(s) {
			fmt.Printf("cap:%d  ===   %d\n",oldcap,newcap)
			oldcap = newcap
		}
	}


	//在使用append追加数据时，长度超过容量 容量自动扩容
	//一般  容量*2   如果超过 1024字节  每次扩容上一次1/4
	//容量扩容每次偶数
}