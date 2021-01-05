package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main0401() {

	//1.导入头文件  math/rand
	//2.随机数种子
	//3.创建随机数


	//创建随机数种子
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Int())   //生成比较大随机数
	fmt.Println(rand.Intn(10))  //常用   取模10   0-9


}
//练习 双色球
func main()  {

	//随机数
	//红球  1-33 选择6个   不重复   蓝球  1-16  选择1个   可以和红球重复


	rand.Seed(time.Now().UnixNano())

	var red [6]int

	for i:=0; i<len(red);i++  {
		v := rand.Intn(33)+1

		for j:=0;j<i ;j++  {
			//数据重复
			if v == red[j] {
				//重新随机
				v = rand.Intn(33)+1

				j = -1
			}

		}
		red[i] = v
	}
	fmt.Println("红球",red,"蓝球",rand.Intn(16)+1)

}
