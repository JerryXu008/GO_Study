package main

import "fmt"

func main() {

	//一.for循环嵌套

	//外层执行一次 内层执行一周
	//count:=0
	//for i:= 0;i<5;i++{
	//	for j:=0;j<5 ;j++  {
	//		fmt.Println(i,j)
	//		count++
	//
	//	}
	//}
	//fmt.Println(count)



	//二.练习
	//百钱百鸡
	/*
	中国古代数学家张丘建在他的《算经》中提出了一个著名的“百钱百鸡问题”：
	一只公鸡值五钱，一只母鸡值三钱，三只小鸡值一钱，
	现在要用百钱买百鸡，请问公鸡、母鸡、小鸡各多少只？
	 */

	 //cock 公鸡个数  hen母鸡个数  chicken 小鸡个数
	 count:=0
	 for cock:=0;cock<=20;cock++{
		for hen:=0;hen<=33 ;hen++  {
			 for chicken:=0;chicken<=100 ;chicken+=3  {
				count++
			 	if cock+hen+chicken ==100 && cock*5+hen*3+chicken/3==100{
			 		fmt.Println(cock,hen,chicken)
				}

			 }
		}

	 }
	 fmt.Println(count)

	// count:=0
	//for cock:=0;cock<=20;cock++{
	//	for hen:=0;hen<=33 ;hen++  {
	//		count++
	//		chicken:=100-cock-hen
	//		if chicken %3 == 0 && cock*5+hen*3+chicken/3==100{
	//			fmt.Println(cock,hen,chicken)
	//
	//		}
	//	}
	//}
	//fmt.Println(count)
}
