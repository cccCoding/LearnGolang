package main

import (
	"sync"
	"gopl.io/ch8/thumbnail"
	"log"
	"os"
)

func main() {

}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)	//在进入worker goroutine之前计数加一
		//worker
		go func() {
			defer wg.Done()		//保证计数减一
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}()
	}

	go func() {		//不能在main中操作
		wg.Wait()		//减到0时关闭channel
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
