package main

import "fmt"

func main()  {

	var arr [10]int = [10]int{9,1,4,2,3,6,5,7,8,10}
	fmt.Println(arr)
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

}