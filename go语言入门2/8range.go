package main

import "fmt"

func main() {
	//数组求和
	var nums = []int{1,2,3,4,5,6}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println(sum)

	//用range实现
	sum2 := 0
	for i, num := range nums {	//range nums 只能用于循环
		sum2 += num
		fmt.Println(sum2, num, i)
	}
	fmt.Println(sum2)

	fmt.Println("--------")

	var str = "你好光头强hahaha"
	for i, c := range str {
		fmt.Printf("%d %c\n",i,c)
	}
}