package main

import "fmt"

func main1() {
	//基本数据类型,赋值时各自独立,不共享内存
	var num = 100
	fmt.Println(num,&num)	//变量及其地址0xc0420080d8
	var num2 = num
	fmt.Println(num2,&num2)	//值相同,地址不同

	var str = "朱建涛"
	var str2 = str
	//var str3				//声明了变量必须使用,否则报错
	fmt.Println(str,&str)
	fmt.Println(str2,&str2)	//值相同,地址也相同,一个字符串指向一个内存

}

func main() {
	var _,x = 100,200	// _ 代表占位,无法显示,读取和写入
	fmt.Println(x)
}