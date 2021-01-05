package main

import "fmt"

func main() {

	//i:=0
	//for {  //for后面不写任何的东西  那么就是一个死循环  永远为真
	//	i++
	//	time.Sleep(time.Second)
	//
	//	if i == 5{
	//		//break
	//		continue
	//	}
	//	fmt.Println(i)
	//}


	//1.break  跳出循环   如果多层for循环  跳出最近的内循环

	//2.continue  跳出本次循环  进入下次循环
	//注意：continue只能在for循环使用

	//3.goto 跳转  可以用在任何地方，但是不能跨函数使用

	fmt.Println("aaaaaaaa")
	   //goto关键字  End 是标签  用户起的名字
	fmt.Println("bbbbbbbb")
End:
	fmt.Println("cccccccc")
	goto End


}
