package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

//节点数据信息
type NodeInfo struct {
	//节点id
	NodeId int `json:"nodeID"`
	//ip
	NodeIpAddr string `json:"nodeIpAddr"`
	//端口
	Port string `json:"port"`
}

//节点信息格式化输出
func (node *NodeInfo) String() string {
	return "NodeInfo:{nodeID" + strconv.Itoa(node.NodeId) + ",nodeIpAddr" + node.NodeIpAddr + ",port" + node.Port + "}"
}

//添加节点到集群,请求格式
type AddToClusterMessage struct {
	//源节点
	Source NodeInfo `json:"source"`
	//目的节点
	Dest NodeInfo `json:"dest"`
	//节点链接发送信息
	Message string `json:"message"`
}

//输出信息
func (req AddToClusterMessage) String() string {
	return "AddToClusterMessage:{\n source:" + req.Source.String() + "\n dest:" + req.Dest.String() + ",\n message:" + req.Message + "}"
}

//生成追加到服务器的信息
func genAddToClusterMessage(source NodeInfo, dest NodeInfo, message string) AddToClusterMessage {
	return AddToClusterMessage{
		Source: NodeInfo{
			NodeId:     source.NodeId,
			NodeIpAddr: source.NodeIpAddr,
			Port:       source.Port},
		Dest: NodeInfo{
			NodeId:     dest.NodeId,
			NodeIpAddr: dest.NodeIpAddr,
			Port:       dest.Port},
		Message: message,
	}
}

//链接到集群的信息
func connectToCluster(me NodeInfo, dest NodeInfo) bool {
	//链接socket,超时时间为10秒
	connOut, err := net.DialTimeout("tcp", dest.NodeIpAddr+":"+dest.Port, time.Duration(10)*time.Second)
	if err != nil {
		if _, ok := err.(net.Error); ok {
			fmt.Println("没有链接到集群", me.NodeId)
			return false
		}
	} else {
		fmt.Println("链接到集群,发送消息到节点")
		text := "hi,哥们,请添加我到节点"
		requestMessage := AddToClusterMessage{me, dest, text} //请求消息
		json.NewEncoder(connOut).Encode(&requestMessage)      //编码到connOut中
		decoder := json.NewDecoder(connOut)                   //从connOut中解码收到的信息
		var responseMessage AddToClusterMessage
		decoder.Decode(&responseMessage)
		fmt.Println("得到数据效应\n" + responseMessage.String())
		return true
	}
	return false
}

//接听
func listenOnPort(me NodeInfo) {
	//监听信息
	ln, _ := net.Listen("tcp", fmt.Sprintf(":"+me.Port))
	for {
		connIn, err := ln.Accept() //监听
		if err != nil {
			if _, ok := err.(net.Error); ok {
				fmt.Println("收信息的时候出现网络错误", me.NodeId)
			}
		} else {
			var requestMessage AddToClusterMessage          //处理要加入的节点
			json.NewDecoder(connIn).Decode(&requestMessage) //解码流中信息到requestMessage
			fmt.Println("获得请求\n" + requestMessage.String())
			text := "确定,欢迎加入"
			responseMessage := genAddToClusterMessage(me, requestMessage.Source, text) //发送给节点的信息
			json.NewEncoder(connIn).Encode(responseMessage)	//发送信息,通过connIn
			connIn.Close() //关闭网络
		}
	}
}

func main() {
	fmt.Println("game start")
	makeMasterOnError := flag.Bool("makeMasterOnError", false, "IP没有链接集群,我们将本IP设置为Master节点")
	clusterip := flag.String("clusterip", "127.0.0.1:8001", "任何节点可以链接这个ip")
	myport := flag.String("myport", "8001", "ip addr正在运行这个节点,端口是8001")
	flag.Parse() //解析命令行参数
	fmt.Println(*makeMasterOnError)
	fmt.Println(*clusterip)
	fmt.Println(*myport)

	//节点生成ID
	rand.Seed(time.Now().UTC().UnixNano()) //生成随机数种子
	myid := rand.Intn(99999)               //生成随机数
	fmt.Println(myid)

	//抓取ip地址
	myip, _ := net.InterfaceAddrs()
	fmt.Println(myip[0])

	//创建当前节点信息
	me := NodeInfo{
		NodeId:     myid,
		NodeIpAddr: myip[0].String(),
		Port:       *myport,
	}
	fmt.Println("当前节点", me)

	dest := NodeInfo{
		NodeId:     -1,
		NodeIpAddr: strings.Split(*clusterip, ":")[0],
		Port:       strings.Split(*clusterip, ":")[1],
	}
	fmt.Println("目标节点", dest)

	//尝试链接
	ableToConnect := connectToCluster(me, dest)

	//监听其他节点加入集群
	if ableToConnect || (!ableToConnect && *makeMasterOnError) {
		if *makeMasterOnError {
			fmt.Println("当前节点是主节点")
		}
		listenOnPort(me)
	} else {
		fmt.Println("退出系统,错误")
	}

	//命令行 Node.exe -h 查看帮助
	//Node.exe --makeMasterOnError	//作为主节点启动
	//Node.exe --myport 8002 --clusterip 127.0.0.1:8001	//启动子节点链接到集群
	//Node.exe --myport 8003 --clusterip 127.0.0.1:8001	//启动子节点链接到集群
}
