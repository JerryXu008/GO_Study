package main

import "fmt"

type Student struct {
	id    int
	name  string
	score int
}

//结构体数组作为函数参数
//结构体切片作为函数参数
func Sort(arr []Student) {
	//外层控制行
	for i := 0; i < len(arr)-1; i++ {
		//内层控制列
		for j := 0; j < len(arr)-1-i; j++ {
			//满足条件 进行交换  大于 升序  小于  降序
			if arr[j].score < arr[j+1].score {
				//交换数据
				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
		}
	}
	fmt.Println(arr)
}

func main() {

	//结构体数组 | 切片
	var arr []Student = []Student{
		Student{1, "李白", 100},
		Student{2, "王维", 100},
		Student{3, "杜甫", 100}}

	//fmt.Println(arr)
	//循环打印结构体信息
	//for i:=0;i<len(arr) ; i++ {
	//	fmt.Println(arr[i])
	//}
	//修改结构体成员指定信息
	arr[1].score = 97
	arr[2].score = 99

	//结构体切片 添加数据
	arr = append(arr, Student{4, "袁华", 100})

	//结构体数组作为函数参数  值传递

	//结构体切片作为函数参数  地址传递 （引用传递）
	Sort(arr)
	fmt.Println(arr)
}
