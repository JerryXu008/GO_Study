package main

import "fmt"

func main1001() {
	//append()
	//copy()

	var s []int = []int{1,2,3,4,5}

	//s1 := make([]int,5)
	s1 := []int{6,7,8,9}
	copy(s,s1)
	fmt.Println(s1)
	//使用copy进行拷贝 在内存中存储两个独立的切片内容 如果任意一个发生修改 不会影响到另一个
	fmt.Printf("%p\n",s)
	fmt.Printf("%p\n",s1)

	s1[2] = 123
	fmt.Println(s)
	fmt.Println(s1)

	s = append(s,6,7,8,9)
	fmt.Println(s)
	fmt.Println(s1)


}

func main()  {
	var arr []int = []int{8,7,5,6,4,2,3,1,9}


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