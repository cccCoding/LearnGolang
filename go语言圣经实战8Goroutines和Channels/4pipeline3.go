package main

import (
	"time"
	"fmt"
)

//tip:不能将一个单项型的channel转换为双向型channel

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {		//chan<- int 表示单方向channel类型,只用于发送
	for x := 0; x < 10; x++ {
		out <- x
		time.Sleep(1 * time.Second)
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {	//<-chan int 表示单方向channel类型,只用于接收
	for x := range in {
		fmt.Println(x)
	}
}
