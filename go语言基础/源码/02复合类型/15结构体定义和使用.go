package main

import "fmt"

//结构体定义在函数外部

//定义

//type 结构体名 struct {
//	//结构体成员列表
//	//成员名 数据类型
//	name string
//	age int
//}

// 结构体名   大小写
type student struct {
	id   int
	name string
	sex  string
	age  int
	addr string
}

func main1501() {
	//顺序初始化 每个成员必须初始化
	var s1 student = student{1, "小明", "男", 18, "市区"}
	fmt.Println(s1)

	//自动推导类型  和  指定成员赋值
	stu := student{name: "小花", age: 18, sex: "女", id: 2, addr: "市区"}

	fmt.Println(stu)
	//定义结构体变量  复合类型
	//var 变量名 结构体名
	var s2 student

	//赋值
	//为结构体成员赋值  .  包名.函数名 结构体.成员   对象.方法
	s2.name = "小亮"
	s2.age = 19
	s2.sex = "男"
	s2.id = 3
	s2.addr = "郊区"

	fmt.Println(s2)

}

func main1502() {
	stu := student{1,"小红","女",19,"市区"}
	//结构体名本身指向第一个成员的地址
	fmt.Printf("%p\n",&stu)
	fmt.Printf("%p\n",&stu.id)

}

func main()  {
	//结构体比较
	//s1 :=student{1,"小明","男",18,"市区"}
	//s2 :=student{1, "小明", "男", 18, "市区"}
	//s3 :=student{2,"小亮","男",18,"市区"}
	//fmt.Println(s1 == s2)
	//fmt.Println(s1 == s3)

	//结构体赋值
	s3 :=student{2,"小亮","男",18,"市区"}

	var s4 student
	s4 = s3
	fmt.Println(s4)

}