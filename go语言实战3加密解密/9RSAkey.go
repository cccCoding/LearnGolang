package main

import (
	"flag"
	"fmt"
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func main() {
	//生成密钥
	var bits int
	flag.IntVar(&bits,"b",1024,"密码长度默认1024")
	if err := GenRsaKey(bits); err != nil {
		fmt.Println("密钥文件生成失败")
	}
	fmt.Println("密钥文件生成成功")
}

func GenRsaKey(bits int) error {
	//生成私钥
	privatekey,err := rsa.GenerateKey(rand.Reader,bits)
	if err != nil {
		fmt.Println("私钥生成失败")
		return err
	}
	fmt.Println("私钥",privatekey)
	derStream := x509.MarshalPKCS1PrivateKey(privatekey)//处理私钥
	block := &pem.Block{		//按照块处理
		Type:"RSA PRIVATE KEY",
		Bytes:derStream,
	}
	privatefile,err := os.Create("myprivate.pem")//创建私钥文件
	if err != nil {
		fmt.Println("私钥文件生成失败")
		return err
	}
	defer privatefile.Close()	//延迟关闭文件
	err = pem.Encode(privatefile, block)	//写入文件
	if err != nil {
		fmt.Println("私钥文件写入失败")
		return err
	}
	//获取公钥
	publickey := &privatekey.PublicKey
	fmt.Println("公钥",publickey)
	derpkix,err := x509.MarshalPKIXPublicKey(publickey)//处理公钥
	if err != nil {
		fmt.Println("公钥处理出错")
		return err
	}
	block = &pem.Block{
		Type:"PUBLIC KEY",
		Bytes:derpkix,	//模块处理key
	}
	publicfile,err := os.Create("mypublic.pem")//创建公钥文件
	if err != nil {
		fmt.Println("公钥文件生成失败")
		return err
	}
	defer publicfile.Close()
	err = pem.Encode(publicfile,block)	//编码写入
	if err != nil {
		fmt.Println("公钥文件写入失败")
		return err
	}
	return nil
}
