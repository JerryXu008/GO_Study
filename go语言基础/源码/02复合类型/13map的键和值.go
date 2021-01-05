package main

import (
	"fmt"
)

func main1301() {

	//map中key类型必须支持 == ！= 一般建议写 基本类型
	//切片 函数 切片的结构类型  不能作为字典的key
	//m:= make(map[[]string]int)  //err  invalid map key type []string
	//map中 数据是无序的
	m := make(map[string][3]int)

	m["小明"] = [3]int{100, 99, 98}
	m["小亮"] = [3]int{6, 9, 8}
	m["小红"] = [3]int{101, 100, 102}

	fmt.Println(m)

	fmt.Println(m)
	//for循环取值
	for k, v := range m {
		fmt.Println(k, v)
	}

}

func main() {
	m := make(map[int]string)
	m[1] = "刘备"
	m[2] = "曹操"
	m[3] = "孙权"

	//通过key取值
	//fmt.Println(m[1])
	//fmt.Println(m["刘备"])  //err
	//在map中如果没有提供key找到具体的值  打印value类型 的默认值
	fmt.Println(m[4])
	//判断键值对是否存在
	value, ok := m[4]
	fmt.Println(value, ok)

	//删除 map中的元素  根据key进行删除
	delete(m,2)
	//如果key不存在  不会报错
	delete(m,4)
	fmt.Println(m)

}
