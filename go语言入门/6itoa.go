package main

import "fmt"

func main() {
	//iota,特殊常量,可以认为是一个可以被编译器修改的常量
	//在每一个const关键字出现时,被重置为0,然后下一个const出现之前,
	// 每出现一次iota,其所代表的数字会自动加一
	const (
		a = iota
		b = iota
		c = iota
		d = iota
	)
	const e int = 10
	const (
		f int = iota
		g int = iota
	)

	fmt.Println(a,d)
	fmt.Println(a,b,c,d,e,f,g)
}
