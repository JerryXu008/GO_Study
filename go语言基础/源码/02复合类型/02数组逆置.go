package main

import "fmt"

func main0201() {
	//练习：从以下一个整数数组中取出最大的整数,最小整数,总和,平均值。
	//var a [6]int = [6]int{0,1,2,3,4,5}
	//1.全部初始化
	var a [5]int = [5]int{-1, -2, -3, -4, -5}

	//2.声明变量  储存最大值 最小值
	var max int = a[0]
	var min int = a[0]
	var sum int

	//3.循环
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]

		}
		sum += a[i]
	}
	fmt.Printf("最大值：%d,最小值：%d，和：%d，平均值：%d", max, min, sum, sum/len(a))

}

func main()  {

	var arr [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
	//数组逆置
	i :=0    //最小下标
	j := len(arr)-1  //最大下标

	//for 表达式1;表达式2;表达式3{}
	//for 返回值 := range 集合{}
	//for 条件{}
	for i < j{
		//if i>=j {
		//	//跳出循环
		//	break
		//}
		//交换数据
		arr[i],arr[j] = arr[j],arr[i]
		//改变下标
		i++
		j--
	}
	fmt.Println(arr)


}