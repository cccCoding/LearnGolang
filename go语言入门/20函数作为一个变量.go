package main

import "fmt"

func main() {
	//匿名函数
	var myFunc = func(a int, b int) int {
		return a * b
	}
	fmt.Println(myFunc(1,2))

	//丧心病狂写法,返回值为函数
	var myFunc2 = func() func() int {
		return func() int {
			return 10
		}
	}
	fmt.Println(myFunc2())	//返回的是函数 0x4966e0
	fmt.Println(myFunc2()())	//返回的是返回的函数的返回值 10
}


