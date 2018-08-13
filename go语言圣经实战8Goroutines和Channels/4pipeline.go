package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break		//false,表示通道已被关闭
			}
			squares <- x * x
		}
		close(squares)	//关闭Channel
	}()

	//Printer
	for {
		x, ok := <-squares
		if ok {
			fmt.Println(x)
		} else {
			break
		}
	}
}
