package main

import "fmt"

func main() {
	//位运算转为二进制处理

	// & 且
	// | 或
	// ^ 异或 不同则为1
	var num1 = 5			//101
	var num2 = 6			//110
	fmt.Println(num1&num2)	//100=4
	fmt.Println(num1|num2)	//111=7
	fmt.Println(num1^num2)	//011=3

	//左移乘以2,右移除以2
	var num3 = 8
	fmt.Println(num3<<2)	//左移2位,乘以4
	fmt.Println(num3>>1)	//右移1位,除以2
}
