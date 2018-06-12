package main

import "fmt"

func main() {
	var c = add(3,5)
	fmt.Println(c)
	var x,y = addsub(1,3)
	fmt.Println(x,y)
}

//函数的参数,返回值都要指定类型
func add(a int, b int) int {
	return a + b
}

func addsub(a int, b int) (int, int) { //返回多个数据
	return a+b,a-b
}
