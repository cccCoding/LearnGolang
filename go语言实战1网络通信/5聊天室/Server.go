package main

import (
	"fmt"
	"net"
	"os"
	"log"
	"strings"
)

//客户端链接,key是ip端口,value是链接对象
var onlineConns = make(map[string]net.Conn)
//消息队列,缓冲区
var messageQueue = make(chan string, 1000)
//消息,处理程序退出
var quitchan = make(chan bool)

const (
	LOG_PATH = "server.log"	//记录错误日志路径
)
var logFile *os.File	//日志文件
var logger *log.Logger	//日志类

func main() {
	//打开文件
	logFile,err := os.OpenFile(LOG_PATH, os.O_RDWR | os.O_CREATE, 0)
	if err != nil {
		fmt.Println("日志文件打开失败")
		os.Exit(-1)
	}
	defer logFile.Close()		//延迟关闭文件
	//创建一个日志对象
	logger = log.New(logFile,"\r\n",log.Ldate|log.Ltime|log.Llongfile)

	//建立TCP服务器
	listen_socket,err := net.Listen("tcp","127.0.0.1:8898")
	defer listen_socket.Close()		//关闭网络
	checkError(err)
	fmt.Println("服务器正在运行")
	logger.Println("写入日志,服务器正在运行")

	go consumeMessage()	//取出并解析消息

	for {
		conn,err := listen_socket.Accept()	//新的客户端链接
		checkError(err)
		//处理每一个客户端
		addr := fmt.Sprintf("%s",conn.RemoteAddr())	//取出地址
		onlineConns[addr] = conn	//追加链接
		fmt.Println("客户端列表-------")
		logger.Println("写入日志,客户端列表-------")
		for i := range onlineConns {	//循环链接
			fmt.Println(i)
			logger.Println("写入日志,i")
		}

		//处理每一个客户端
		go processInfo(conn)
	}
}

//处理每一个客户端
func processInfo(conn net.Conn) {
	buf := make([]byte, 1024)	//开创缓冲区

	//处理客户端退出	//延迟执行,在for循环停止后执行
	defer func(conn net.Conn) {
		//记录远程链接
		addr := fmt.Sprintf("%s",conn.RemoteAddr())
		delete(onlineConns, addr)	//客户端退出时,从来列表删除
		conn.Close()
		for addr := range onlineConns {
			fmt.Println("当前的网络链接", addr)
		}
	}(conn)

	for {
		n,err := conn.Read(buf)	//读取数据
		if err != nil {		//出错
			break
		}
		if n != 0 {
			//消息处理
			message := string(buf[:n])
			//消息队列存储消息
			messageQueue<-message
			fmt.Println("收到来自",conn.RemoteAddr(),"的消息",message)
		}
	}
}

//消息的协程
func consumeMessage() {
	for {
		select {
		case message := <-messageQueue:		//取出消息
			doProcessMessage(message)		//解析消息
		case <-quitchan:					//处理退出
			break
		}
	}
}

//消息的解析函数
func doProcessMessage(message string) {
	contents := strings.Split(message,"#")	//字符串切割
	if len(contents) > 1 {
		addr := contents[0]	//取出地址
		sendmessage := contents[1]	//取出消息
		addr = strings.Trim(addr,  " ")

		if sendmessage == "list" {
			var listip string = ""
			for addr := range onlineConns {
				fmt.Println("在线客户端",addr)
				listip += "||||" + addr
			}
			//给请求的客户端发当前客户端列表
			for _, conn := range onlineConns {
				remoteAddr := fmt.Sprintf("%s",conn.RemoteAddr())
				if remoteAddr == addr {
					_,err := conn.Write([]byte(listip))
					if err != nil {
						fmt.Println("发送失败",err)
					}
				}
			}
		} else {
			//给所有人发消息
			for addr, conn := range onlineConns {
				_,err := conn.Write([]byte(sendmessage))
				fmt.Println("服务器发送消息",sendmessage,"给",addr)
				if err != nil {
					fmt.Println("在线链接发送失败")
				}
			}
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error",err.Error())
		os.Exit(0)	//出现错误退出
	}
}
