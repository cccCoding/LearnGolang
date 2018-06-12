package main

import (
	"os"
	"io"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

//秒上传,服务器检测要上传文件与已有的hash是否一样

func main() {
	filepath := "haha.txt"
	if hash, err := hashsha256file(filepath); err == nil {
		fmt.Printf("%s sha256hash is %s", filepath, hash)
	} else {
		fmt.Printf("%s error is %s", filepath, err)
	}
}

func hashsha256file(filepath string) (string, error) {
	var hashvalue string	//要返回的hash字符串
	file,err := os.Open(filepath)
	if err != nil {
		return hashvalue,err	//返回错误,文件打开失败,hash为空
	}
	defer file.Close()
	var myhash = sha256.New()	//创建hash算法对象
	if _,err := io.Copy(myhash,file);err!=nil {		//从文件读取,写入到hash算法对象myhash中
		return hashvalue,err	//处理拷贝错误
	}
	hashinbytes := myhash.Sum(nil)
	hashvalue = hex.EncodeToString(hashinbytes)	//hash转换字符串
	return hashvalue,nil
}
