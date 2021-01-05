package main

import "fmt"

//函数定义   形参
//参数名  自己起   ... type  不定参
func sum(a int,args ...int)  {

	//fmt.Println(args)
	//sum:= args[0]+args[1]+args[2]+args[3]
	//fmt.Println(sum)

	//fmt.Println(len("hello"))

	//count:=len(args)
	//fmt.Println(count)


	sum := 0
	//i 下标

	//for循环遍历
	//for i:=0;i<len(args) ; i++ {
	//	//fmt.Println(args[i])
	//	sum+=args[i]
	//}
	//fmt.Println(sum)


	//匿名变量  丢弃数据
	for _,v:=range args{
		//fmt.Println(v)
		sum+=v
	}
	fmt.Println(sum,"hello",444,55,5,5,5,5,5,5,5,5)
}

func main() {

	//函数调用
	sum(1)
}


//1.不定参放最后
//2.函数调用  固定参数必须传值  不定参数 根据需要决定是否传值