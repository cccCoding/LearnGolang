package main

import (
	"fmt"
	"encoding/base64"
)

const base64Table string = "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm0123456789*/"
var coder = base64.NewEncoding(base64Table)	//编码解码的表格

func main() {
	fmt.Println(len(base64Table))
	var str = "你好我我"
	debyte := base64encode([]byte(str))
	enbyte,err := base64decode([]byte(debyte))
	if err != nil {
		fmt.Println("错误",err.Error())
	}
	fmt.Println(str)
	fmt.Println(string(debyte))
	fmt.Println(string(enbyte))
}

func base64encode(src []byte) []byte {		//编码
	return []byte(coder.EncodeToString(src))
}

func base64decode(src []byte) ([]byte, error) {	//解码
	return coder.DecodeString(string(src))
}
