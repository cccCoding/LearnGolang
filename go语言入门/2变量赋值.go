package main

import "fmt"

func main2() {
	//go语言会自动推理数据类型,但是不指明数据则数据类型必须赋值
	var a = "朱建涛"	//变量定义,自适应 变量名不能用标识符if之类的
	fmt.Println(a)
	var name string = "莽子"
	fmt.Println(name)
	var key string
	key = "啦啦啦"		//指定类型可以晚些赋值
	fmt.Println(key)

	var num int = 10	//指定int类型变量
	fmt.Println("朱建涛有",num,"套房子")
	var num1 int
	num1 = 20
	fmt.Println("朱建涛有",num1,"辆保时捷")
	var length float32 = 1.23
	fmt.Println(length)
}

func main()  {
	//var c = go语言入门
	c:=10		// := 表示新建一个变量并赋值,如果有重名变量,会报错
	fmt.Println(c)

	//多个变量赋值
	var a,b int
	a,b=1,2
	fmt.Println(a,b)
	var x,y,z=4,5,6
	fmt.Println(x,y,z)
	k,v:=100,200
	fmt.Println(k,v)

	//一般用于全局变量
	var (
		age int
		name string
	)
	age = 10
	name = "酥饼"
	fmt.Println(age,name)
}