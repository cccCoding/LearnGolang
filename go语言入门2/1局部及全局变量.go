package main

import "fmt"

var b string = "莽子"	//全局变量,函数外部,多个函数都可引用

func main() {
	var x,y,z int
	x = 10		//局部变量,作用于函数内部
	y = 20
	z = 30
	fmt.Println(x,y,z,b)
	hello()
}

func hello() {
	var a = "朱建涛"
	fmt.Println(a, b)
	var b string = "酥酥饼"		//局部变量优先于同名的全局变量
	fmt.Println(a, b)
}
