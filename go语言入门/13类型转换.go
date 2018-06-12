package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num = 18.5
	fmt.Println(num)
	fmt.Println(int(num))	//实数转int数

	var str = "12345"
	var strNum,error = strconv.Atoi(str)	//字符串转int数
	strNum -= 5
	fmt.Println(str,strNum,error)

	var a = 2345
	mystr := strconv.Itoa(a)	//数字转字符串
	fmt.Println(mystr)

}
