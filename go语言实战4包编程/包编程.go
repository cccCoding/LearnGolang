package main //执行模块运行必须是main

import "fmt"
import "./move"
import "./move/go"

func main()  {
	fmt.Println("hello world")
	move.Isgo()
	_go.IsGoTest()
}

//命令行 go get ... 获取包,然后导入
