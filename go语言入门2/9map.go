package main

import "fmt"

func main() {
	var a map[string]string		//定义类型
	a = make(map[string]string)	//开辟内存
	a["一号"] = "我是一号"		//插入数据
	a["二号"] = "我是二号"
	a["三号"] = "我是三号"
	a["一号"] = "我是真正的一号"	//key重复时替换value
	for adc := range a {		//循环
		fmt.Println(adc, a[adc])
	}

	delete(a, "一号")	//删除

	//判断key是否存在
	b,ok := a["一号"]
	if ok {
		fmt.Println("一号存在",b)
	} else {
		fmt.Println("一号不存在",b)
	}
}
