package main

import "fmt"

func main() {
	var a int = 10
	var b int = 2
	fmt.Println(a+b, a-b)
	a++		//a加一,但是不能引用,不能赋值
	fmt.Println(a)
}
