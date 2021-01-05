package main

import "fmt"

func main0101()  {
	//数组  是指一系列同一类型数据的集合
	//数组定义
	//var 数组名 [长度]类型
	var a [10]int
	//fmt.Println(len(a))
	fmt.Println(a)

	//注意以下方式
	//报错err
	//var arr int = 10
	//var a [n]int

	//数组赋值
	//a[0] = 1
	//a[1] = 2
	//a[2] = 3
	//a[3] = 4
	//fmt.Println(a)
	//fmt.Println(a[0])
	//fmt.Println(a[1])
	//下标超范围  数组越界
	//fmt.Println(a[10])
	//for循环 数组赋值 输出
	for i:=0;i<10;i++{
		a[i] = i+1
	}
	fmt.Println(a)

	for i:=0;i<10 ;i++  {
		fmt.Println(a[i])
	}
	//匿名变量
	for _,v := range a{
		fmt.Println(v)
	}
	//数组存储的元素类型可以是其他类型
	//定义完成  直接输出
	//var a [10]float64  //直接输出  默认为0
	//var a [10]string  //直接输出  默认为空字符
	//var a [10]bool  //直接输出  默认为false
}

func main0102()  {
	//数组初始化
	//1.全部初始化
	var a [5]int = [5]int{1,2,3,4,5}
	fmt.Println(a)

	//自动推导
	b := [5]int{1,2,3,4,5}
	fmt.Println(b)

	//部分初始化
	//没有初始化的部分 默认为0
	c := [5]int{1,2,3}
	fmt.Println(c)

	//指定某个元素初始化

	d := [5]int{2:10,4:20}
	fmt.Println(d)
	//...  通过初始化确定长度
	f:=[...]int{1,2,3}
	fmt.Println(len(f))

}

func main()  {

	//常见问题
	//数组长度  常量
	var arr [5]int = [5]int{1,2,3,4,5}
	//数组下标  数组下标越界
	//arr[6] = 123  //err
	//arr[-1] = 123 //err
	//数组名
	//arr = 123 //err

	//两个数组类型相同 个数相同  可以赋值
	//arr1 := arr
	//
	//fmt.Println(arr)
	//fmt.Println(arr1)
	//
	//fmt.Printf("%T\n",arr)

	//数组名表示 整个数组  数组名对应地址 就是数组第一个元素的地址
	fmt.Printf("%p\n",&arr)
	fmt.Printf("%p\n",&arr[0])
	fmt.Printf("%p\n",&arr[1])
	fmt.Printf("%p\n",&arr[2])

}