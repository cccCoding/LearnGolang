package main

import (
	"crypto/md5"
	"fmt"
)

//密码一般md5加密,可以实现密码保密和密码校验
func main() {
	mymd5 := md5.New()
	mymd5.Write([]byte("hello md5"))
	result := mymd5.Sum([]byte(""))
	fmt.Println(result)
	fmt.Printf("%x", result)
}
