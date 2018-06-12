package main

import "fmt"

type Person struct {
	name string
	age int
}
type Man struct {
	Person		//实现了继承,拥有Person的属性和方法
	hobby string	//子类独有的属性
}

func (person Person)getNameAndAge() (string, int) {
	return person.name,person.age
}

func (m Man) getHobby() string {	//子类独有的方法
	return m.hobby
}

func main() {
	var p = new(Person)
	p.name = "大头"
	p.age = 19
	fmt.Println(p.getNameAndAge())

	var m = new(Man)
	m.name = "朱建涛"
	m.age = 25
	m.hobby = "code"
	fmt.Println(m.getNameAndAge())
	fmt.Println(m.getHobby())
}