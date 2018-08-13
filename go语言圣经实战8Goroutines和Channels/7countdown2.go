package main

import (
	"os"
	"time"
	"fmt"
)

//倒计时10秒,然后发射,可中断
func main() {

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	select {
	case <-time.After(10 * time.Second):
		//Do nothing
	case <-abort:
		fmt.Println("Launch abort")
		return
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
