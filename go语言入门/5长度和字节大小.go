package main

import (
	"fmt"
	"unsafe"
)

func main() {
	const str = "adc"
	//len 数组或字符串长度
	//字符串的内存地址一般用16字节来保存
	fmt.Println(str,len(str),unsafe.Sizeof(str))

	var num uint32 = 10		//32位无符号整数
	fmt.Println(num,unsafe.Sizeof(num))	//num占内存字节大小

	var num2 uint64 = 10		//64位无符号整数
	fmt.Println(num2,unsafe.Sizeof(num2))	//num2占内存字节大小

	var num3 uint = 10
	fmt.Println(num3,unsafe.Sizeof(num3))	//num3占内存字节大小
}