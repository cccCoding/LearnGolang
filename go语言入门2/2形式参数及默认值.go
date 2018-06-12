package main

import "fmt"

var num int
var money float32
var pnum *int	//指针,默认为 nil

func main() {
	var x, y int
	x = 10
	y = 20
	add(x,y)
	fmt.Println(x,y)				//值并没有被改变
	fmt.Println(num,money,pnum)		//打印默认值 0 0 <nil>
}

//参数有副本机制,形式参数a,b当作局部变量
func add(a int, b int) {
	a = 20
	b = 40
	fmt.Println(a,b)
	fmt.Println("go add",a + b)
}
