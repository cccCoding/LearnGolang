package main

import (
	"net"
	"time"
	"fmt"
	"strings"
	"bufio"
	"log"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn ,err := l.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprint(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprint(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		//上一个shout被处理才会处理下一个
		//echo(c, input.Text(), 1*time.Second)

		//并发处理
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
