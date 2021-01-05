package main

import "fmt"

func swap(a int, b int) {
	a, b = b, a
}

func main0501() {
	a := 10
	b := 20
	swap(a, b)
	//值传递
	fmt.Println(a)
	fmt.Println(b)
}
//数组作为函数参数  返回值
func BubbleSort(arr [10]int) [10]int {
	//外层控制行
	for i:=0;i<len(arr)-1 ; i++ {
		//内层控制列
		for j:=0;j<len(arr)-1-i ; j++ {
			//满足条件 进行交换  大于 升序  小于  降序
			if arr[j]<arr[j+1] {
				//交换数据
				arr[j],arr[j+1] = arr[j+1],arr[j]

			}
		}
	}
	fmt.Println(arr)
	return arr
}


func main()  {

	arr := [10]int{9,1,5,6,8,4,7,8,10,3}
	//数组作为函数参数是值传递
	//形参和实参 是不同地址

	//内存中有两份独立的数组  储存不同的数据
	//在函数调用结束后  内存回收  不会影响主函数实参的值


	//如果想通过函数计算结果改变实参的值  需要使用数组作为函数的返回值
	arr1 := BubbleSort(arr)

	fmt.Println(arr)
	fmt.Println(arr1)
}