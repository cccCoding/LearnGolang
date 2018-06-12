package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	//tcp网络链接
	conn,err := net.Dial("tcp","127.0.0.1:8898")
	defer conn.Close()		//延迟关闭网络链接	避免内存泄漏	逆序执行
	checkError(err)

	//开启一个协成,处理消息发送
	go MessageSend(conn)

	//协成,负责收取消息
	buf := make([]byte,1024)
	for {
		n,err := conn.Read(buf)
		checkError(err)
		if err != nil {
			fmt.Println("退出客户端")
			os.Exit(0)
		}
		fmt.Println("收到服务器消息",string(buf[:n]))
	}

	fmt.Println("客户端关闭")
}

func MessageSend(conn net.Conn) {
	var input string
	for {
		reader := bufio.NewReader(os.Stdin)	//读取键盘输入
		data,_,_ := reader.ReadLine()		//读取一行
		input = string(data)				//键盘输入转化为字符串
		if input == "exit" {
			conn.Close()
			fmt.Println("客户端关闭")
			break
		}
		addr := fmt.Sprintf("%s",conn.LocalAddr())
		input = addr + "#" + input
		_,err := conn.Write([]byte(input))	//输入写入字符串
		//fmt.Println("发送消息",input)
		if err != nil {
			conn.Close()
			fmt.Println("客户端关闭")
			break
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error",err.Error())
	}
}