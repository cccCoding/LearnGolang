package main

import (
	"fmt"
	"strings"
)

func main() {
	//var cmd = exec.Command("notepad")	//字符串用于执行系统指令
	//fmt.Println(cmd.Run())

	var cmd = "notepad"
	fmt.Println(cmd)		//打印字符串
	fmt.Printf("要执行的命令是\n%s",cmd)	//字符串格式化并打印

	fmt.Println()

	for i:=0;i<len(cmd);i++{
		fmt.Printf("%c",cmd[i])		//打印字符 c表示char
		fmt.Println(cmd[i])	//打印字符编号
	}

	var str = "哼哼" + "哈哈"
	fmt.Println(str)
	var newstr = []string{"嘻嘻","呼呼"}	//字符串数组的归并
	fmt.Println(strings.Join(newstr," 呵呵哒 "))

	//字符串格式化
	fmt.Printf("朱建涛有%d辆车\n",10)
	fmt.Printf("朱建涛有%f辆车\n",10.6)
	fmt.Printf("朱建涛有%s辆车\n","100")
}
