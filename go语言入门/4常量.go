package main

import "fmt"

func main() {
	fmt.Println("朱建涛有",10,"套房子")	//十进制
	fmt.Println("王志豪有",012,"套房子")//八进制
	fmt.Println("酥酥饼有",0x0a,"套房子")//十六进制
	//fmt.Println("朱建涛有",10u,"套房子")	//正数???u
	//fmt.Println("朱建涛有",10l,"套房子")	//long???l

	const length1 float32 = 1.23	//常量,无法修改,重新赋值
	fmt.Println(length1)

	const length2 = 9e+2
	fmt.Println(length2)

	//枚举
	const(
		a=0
		b=1
		c=2
	)
	const aa = a
	fmt.Println(aa)
}