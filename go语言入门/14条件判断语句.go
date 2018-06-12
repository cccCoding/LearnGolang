package main

import "fmt"

func main() {
	var a = 10
	//var b = 12
	if b := 12; a > b {			//定义一个变量可以放在if内部
		fmt.Println("a大于b")
	} else if a == b {
		fmt.Println("a等于b")
	} else {
		fmt.Println("a小于b")
	}

	var c = 21
	switch c {
	case 20,21,22:						//哇塞可以没有break???
		fmt.Println("喜欢帅哥")
	case 30:
		fmt.Println("喜欢有钱")
	case 40,41,42:
		fmt.Println("喜欢暖男")
	default:
		fmt.Println("啦啦啦")
	}

	var x interface{} = "你好"		//接口语法
	switch y := x.(type){		//获取类型
	case int:
		fmt.Println("整数类型", y)
	case string:
		fmt.Println("字符串", y)
	case bool:
		fmt.Println("布尔类型", y)
	default:
		fmt.Println("这是啥", y)
	}
}
