package main

import "fmt"

func main0601() {

	//var arr [10]int  //一维数组
	//fmt.Println(arr)

	//二维数组定义
	var arr [2][3]int
	//fmt.Println(arr)
	////行数
	//fmt.Println(len(arr))
	//
	////列数
	//fmt.Println(len(arr[1]))
	//
	////赋值
	arr[0][1] = 123
	arr[1][2] = 234
	//fmt.Println(arr)
	//for循环取值

	//外层控制行  内层控制列
	for i:=0;i<len(arr) ;i++  {
		for j:=0;j<len(arr[0]) ;j++  {
			fmt.Println(arr[i][j])
		}
	}

}

func main()  {

	//二维数组
	//初始化
	//1.全部初始化
	//自动推导
	b := [3][4]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	fmt.Println(b)
	//部分初始化
	c := [3][4]int{{1,2,3,4},{5,6,7,8}}
	fmt.Println(c)


}