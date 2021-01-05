package main

import "fmt"

func main() {

	//语法
	//switch 变量(表达式) {   和switch关键字同行
	//case 值1:
	//	代码体
	//case 值2:
	//	代码体
	//default:
	//	代码体
	//}
	//1.支持多个条件的匹配
	//2.不同的case之间不需要break
	var score int = 90
	switch score {

	case 90:
		fmt.Println("A")
		//fallthrough
	case 80:
		fmt.Println("B")

	case 50,60,70:
		fmt.Println("C")

	default:
		fmt.Println("D")
	}
	
	
	//1.

	switch s1:=90;s1 { //初始化语句;条件
	case 90:
		fmt.Println("A")
		//fallthrough
	case 80:
		fmt.Println("B")

	case 50,60,70:
		fmt.Println("C")

	default:
		fmt.Println("D")
	}
	
	var s2 int = 90
	switch  { //没有条件
	case s2>=90: //写判断语句
		fmt.Println("A")
	case s2>=80:
		fmt.Println("B")
	default:

		fmt.Println("C")
		
	}

	//3.

	switch s3:=90; { //只有初始化语句  没有条件
	case s3>=90: //写判断语句
		fmt.Println("A")
	case s3>=80:
		fmt.Println("B")
	default:

		fmt.Println("C")
	}
}


//总结

/*
优点
1.if 适合进行区间的判断  嵌套使用
2.switch  固定值    执行效率高  可以将多个满足相同条件的值放在一起

缺点：
1.if 执行效率低
2.switch 不建议使用嵌套
 */
