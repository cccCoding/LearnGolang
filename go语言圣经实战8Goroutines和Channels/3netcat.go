package main

import (
	"net"
	"github.com/labstack/gommon/log"
	"io"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Print("done\n")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)

	conn.Close()
	//TODO 只关闭写的部分,后台goroutine可以在标准输入被关闭后继续打印从reverb1服务器传回的数据

	//等待后台goroutine的done,再结束main goroutine
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
