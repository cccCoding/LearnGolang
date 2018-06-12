package main

import "fmt"

func main() {
	var a = 100
	var b = 101
	var c = 100
	var d = 101

	if a == b && c > d {
		fmt.Println("阿里阿里")
	} else {
		fmt.Println("河马河马")
	}

	if a != b || c > d {
		fmt.Println("阿里阿里")
	} else {
		fmt.Println("河马河马")
	}
}