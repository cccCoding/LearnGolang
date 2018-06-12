package main

import (
	"os"
	"fmt"
	"net"
)

func main() {
	//创造UDP服务器
	udp_addr,err := net.ResolveUDPAddr("udp",":8848")
	checkErr(err)

	conn,err := net.ListenUDP("udp",udp_addr)	//开始监听
	defer conn.Close()	//关闭链接
	checkErr(err)

	recvUDPMsg(conn)	//收消息
}

//处理错误信息
func checkErr(err error) {
	if err != nil {		//指针不为空
		fmt.Println("Error",err.Error())
		os.Exit(1)
	}
}

//接收消息
func recvUDPMsg(conn*net.UDPConn) {
	var buf [30] byte                         //缓冲数组
	n, raddr, err := conn.ReadFromUDP(buf[:]) //从UDP接受数据
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	fmt.Println("msg is",string(buf[0:n]))	//n代表读取了多少字节
	fmt.Println("addr is", raddr)
	_,err = conn.WriteToUDP([]byte("hahahaha"), raddr)	//写入UDP,根据地址发送
	checkErr(err)
}