package main

import "fmt"

func main() {
	var times = 3
	for i := 0; i < times; i++ {
		fmt.Println("i的值为",i)
	}
	//for ; ; {		//死循环
	//	fmt.Println("hahaha")
	//}
	var a = 0
	b := 5
	for a < b {			//相当于while
		a++
		fmt.Println("a的值为",a)
	}

	//数组循环
	numbers := [5]int{1,2,3,4,5}	//数组
	fmt.Println(numbers)
	for i, number := range numbers {	//range取得数组索引和元素
		fmt.Println(i, number)	//打印索引和对应索引数组元素
	}
	for i := range numbers {
		fmt.Println(i)			//打印索引
	}
}