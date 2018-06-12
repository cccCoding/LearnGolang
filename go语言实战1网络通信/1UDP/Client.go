package main

import (
	"net"
	"fmt"
	"os"
)

func main() {

	//链接网络
	conn,err := net.Dial("udp","127.0.0.1:8848")
	//关闭网络	延迟编程
	defer conn.Close()
	if err != nil {
		fmt.Println("网络链接出错")
		os.Exit(1)
	}
	conn.Write([]byte("hello"))		//发送消息
	fmt.Println("发送消息","hello")

	var msg[30] byte
	conn.Read(msg[:])				//收取消息
	fmt.Println("收到消息",string(msg[:]))
}
