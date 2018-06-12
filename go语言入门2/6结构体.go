package main

import "fmt"

//定义people结构体
type people struct {
	name string
	age int
	height float32
}

func main() {
	var stan people
	stan.name = "朱建涛"
	stan.age = 25
	stan.height = 1.61
	showPeople(stan)
	var aaa *people = &stan		//结构体指针
	fmt.Println(aaa, aaa.name)		//通过结构体指针访问属性

	var bbb[3] people
	bbb[0].name = "name0"
	bbb[1].name = "name1"
	bbb[2].name = "name2"
	fmt.Println(bbb)
}

func showPeople(p people) {
	fmt.Println(p)
}
