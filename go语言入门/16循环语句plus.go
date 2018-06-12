package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i == 4 {
			break	//跳出循环,终止
		}
		fmt.Println(i)
	}

	fmt.Println("--------")

	for j := 0; j < 6; j++ {
		if j == 4 {
			continue	//结束本次循环
		}
		fmt.Println(j)
	}

	fmt.Println("--------")

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(i,j,"      ")
		}
		fmt.Println()
	}

	fmt.Println("--------")

	//LOOP:fmt.Println("哈哈")		//goto实现死循环
	//goto LOOP				//跳转

	var i = 0
	AK:if i < 10 {		//goto实现循环
		i += 1
		fmt.Println(i)
		goto AK
	}

}
