package main

import "fmt"

func main() {
	//函数参数存在副本机制,地址不一样
	var num = 100
	fmt.Println("main",num,&num)
	changeNum(num)
	changeNumP(&num)
	fmt.Println("main",num)	//数据已通过上面内存地址传递修改
}

//传递内存地址修改数据,外挂原理
func changeNumP(ptr *int) {		//传递内存地址
	*ptr = 10000
	fmt.Println("changeNumP",*ptr,ptr)
}

func changeNum(num int) {
	num = 1000
	fmt.Println("changeNum",num,&num)	//地址不同
}
