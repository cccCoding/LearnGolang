package main

import (
	"fmt"
	"os"
	"strings"
)

//os包以跨平台的方式，提供了一些与操作系统交互的函数和变量。
// 程序的命令行参数可从os包的Args变量获取；
// os包外部使用os.Args访问该变量。
//os.Args的第一个元素，os.Args[0], 是命令本身的名字；其它的元素则是程序启动时传给它的参数。
func main() {
	var s, sep string

	//for i := 0; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)	//一直迭代字符串浪费资源，可使用Join方法优化

	// TODO 通过对比时间比较两种方法性能差异

	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println(os.Args[1:])
}
