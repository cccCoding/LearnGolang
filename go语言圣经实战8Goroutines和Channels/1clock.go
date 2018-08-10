package main

import (
	"net"
	"log"
	"io"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		//不支持并发,第二个客户端必须等待第一个客户端完成工作,都在main goroutine中运行
		//handleConn(conn)

		//支持并发,在一个单独的goroutine中运行
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("11:22:00\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
