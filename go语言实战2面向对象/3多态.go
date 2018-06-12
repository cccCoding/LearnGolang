package main

import "fmt"

//接口固定,行为可以扩展
type qinshou interface {
	run()
	eat()
}

type yiguanqinshou interface {
	run()
}

type person struct {
}

func (p person) run() {
	fmt.Println("p run")
}

func (p person) eat() {
	fmt.Println("p eat")
}

type tiger struct {
}

func (p tiger) run() {
	fmt.Println("tiger run")
}

func (p tiger) eat() {
	fmt.Println("tiger eat")
}

func main() {
	var qs qinshou
	var ygqs yiguanqinshou

	qs = new(person)
	qs.run()
	qs.eat()

	ygqs = qs		//可以赋值
	ygqs.run()
	//ygqs.eat()	//但是没得这个方法

	qs = new(tiger)
	qs.run()
	qs.eat()
}
