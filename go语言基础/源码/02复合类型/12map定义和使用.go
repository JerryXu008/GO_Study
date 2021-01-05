package main

import "fmt"

func main() {
	//names := []string{"张三","李四","王五"}
	//fmt.Println(names[2])
	//map[键类型]值类型
	//定义
	//1.
	var m map[int]string
	fmt.Println(m)
	//内存地址编号为0的空间 系统占用  不允许进行读写操作
	fmt.Printf("%p\n",m)
	////在字典中不能使用cap函数 只能使用len函数
	////fmt.Println(cap(m))   //err
	//fmt.Println(len(m))

	//2.
	m2 := make(map[int]string)
	fmt.Println(m2)
	fmt.Printf("%p\n",m2)

	//赋值
	//键是唯一的  值可以重复
	//map 是无序的
	//map长度是自动扩容
	m2[1] = "张三"
	fmt.Println(len(m2))
	m2[2] = "李四"
	fmt.Println(len(m2))
	m2[3] = "王五"
	fmt.Println(len(m2))
	m2[4] = "张三"
	fmt.Println(len(m2))

	fmt.Println(m2)

	//初始化
	m3 := map[int]string{1: "张三", 2: "李四", 3: "王五"}
	fmt.Println(m3[3])



}