package main

import (
	"fmt"
	"net"
)

func main() {
	//tcp网络链接
	conn,err := net.Dial("tcp","127.0.0.1:8898")
	defer conn.Close()		//延迟关闭网络链接	避免内存泄漏	逆序执行
	checkError(err)

	conn.Write([]byte("hello tcp"))
	fmt.Println("发送消息","hello tcp")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error",err.Error())
	}
}