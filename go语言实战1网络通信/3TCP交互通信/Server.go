package main

import (
	"fmt"
	"net"
)

func main() {
	//建立TCP服务器
	listen_socket,err := net.Listen("tcp","127.0.0.1:8898")
	defer listen_socket.Close()		//关闭网络
	checkError(err)
	fmt.Println("服务器正在运行")
	for ; ; {
		conn,err := listen_socket.Accept()	//新的客户端链接
		checkError(err)
		//处理每一个客户端
		go processInfo(conn)
	}
}

func processInfo(conn net.Conn) {
	buf := make([]byte, 1024)	//开创缓冲区
	defer conn.Close()
	for {
		n,err := conn.Read(buf)	//读取数据
		if err != nil {		//出错
			break
		}
		if n != 0 {
			fmt.Println("收到消息",string(buf))
			//反向链接
			remoteAddr := conn.RemoteAddr()	//获取远程地址
			fmt.Println("收到消息来自",remoteAddr)

		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error",err.Error())
	}
}
