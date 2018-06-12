package main

import "fmt"

func main() {
	var num1 = 10 //赋值
	var num2 = 2
	num3 := 3	// := 表示新建一个变量并赋值,如果有重名变量,会报错
	fmt.Println(num1,num2,num3)

	num1 += num2
	num3<<=1	//左移1位并赋值
	num2 &= num1	//相与并赋值
	fmt.Println(num1,num2,num3)
}
