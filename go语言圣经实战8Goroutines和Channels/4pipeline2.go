package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
		close(naturals)	//主动关闭,当它没有被引用时也会被垃圾自动回收器回收
	}()

	go func() {
		//range循环在channels上迭代
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
