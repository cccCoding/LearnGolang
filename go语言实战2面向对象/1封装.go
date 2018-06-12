package main

import "fmt"

//定义类型,包含数据的属性
type People struct {
	name string
	age int
	height float32
}

//函数,数据初始化
func (xiaowang *People) setData(name string, age int, height float32) {
	xiaowang.name = name
	xiaowang.age = age
	xiaowang.height = height
}

func (xiaowang *People) getData() (string,int,float32) {
	return xiaowang.name,xiaowang.age,xiaowang.height
}

func (old *People)compare(new *People) bool {
	return old.age < new.age
}

func main() {
	var 光头强 = new(People)
	光头强.setData("阿强", 18, 188.8)
	fmt.Println(光头强.getData())

	var 熊大 = new(People)
	熊大.setData("熊大大", 8, 288.8)
	fmt.Println(熊大.getData())

	fmt.Println(熊大.compare(光头强))
}
