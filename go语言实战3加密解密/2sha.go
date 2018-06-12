package main

import (
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
)

func main() {
	var mystr = "sha256 hahahaha"
	mysha256 := sha256.New()	//创建加密算法对象
	mysha256.Write([]byte(mystr))	//加入要处理散列的数据
	result := mysha256.Sum(nil)	//结果计算
	fmt.Println(result)
	fmt.Printf("%x",result)

	fmt.Println()

	mysha512 := sha512.New()
	mysha512.Write([]byte(mystr))
	result2 := mysha512.Sum(nil)
	fmt.Printf("%x",result2)
}
