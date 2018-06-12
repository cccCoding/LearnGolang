package main

import "fmt"

func main() {
	//任何数据都有一个内存地址	//地址又称为指针
	var num = 10
	var ptr = &num//指针,存储地址
	fmt.Println(num, ptr, &num)	//打印值,指针和地址
	*(&num) = 100	//根据地址取出内容  //根据地址修改内容
	fmt.Println(num)
	*ptr = 1000
	fmt.Println(num)
}
