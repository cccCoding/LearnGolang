package main

import (
	"bytes"
	"fmt"
	"crypto/des"
	"crypto/cipher"
	"encoding/base64"
)

func main() {
	fmt.Println("game start")
	//DES 8位
	key := []byte("hahahaha")
	result,err := DesEncrypt([]byte("mima呀哈哈哈"),key)	//加密
	if err != nil {
		panic(err)
	}
	fmt.Println("加密后数据",base64.StdEncoding.EncodeToString(result))
	origData,err := DesDecrypt(result,key)	//解密
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后数据",string(origData))
}

func DesEncrypt(origData, key []byte) ([]byte, error) {
	block,err := des.NewCipher(key)		//key作为加密密钥
	if err != nil {
		return nil,err
	}
	origData = zeroPadding(origData,block.BlockSize())	//填充位数0
	blockmode := cipher.NewCBCEncrypter(block,key[:block.BlockSize()])	//加密模式
	crypted := make([]byte,len(origData))	//加密存储的空间
	blockmode.CryptBlocks(crypted,origData)	//加密
	return crypted,nil
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block,err := des.NewCipher(key)		//解密密码
	if err != nil {
		return nil,err
	}
	blockmode := cipher.NewCBCDecrypter(block,key[:block.BlockSize()])
	origData := make([]byte,len(crypted))
	blockmode.CryptBlocks(origData,crypted)//解密
	origData = zerounPadding(origData)	//去掉尾部补充的位数
	return origData,nil
}

//129  32  129=32*4+1
func  zeroPadding(ciphertext []byte,blocksize int )[]byte{
	padding:= blocksize- len(ciphertext)%blocksize //填充的长度
	padtext:=bytes.Repeat([]byte{0},padding)//填充0
	return append(ciphertext,padtext...)//字节的叠加
}

func  zerounPadding(origdata []byte)[]byte{ //删除填充数据
	return  bytes.TrimRightFunc(origdata,func (r rune)bool{
		return r==rune(0)
	})
}

func PKCS5Padding(origdata []byte,blocksize int) []byte {
	padding := blocksize - len(origdata)	//剩余的字节
	//处理需要补充位数的数据
	padtext := bytes.Repeat([]byte{byte(padding)},padding)
	return append(origdata, padtext...)	//追加数据
}

func PKCS5UnPadding(ciphertext []byte) []byte {
	length := len(ciphertext)
	unpadding := int(ciphertext[length - 1])	//取出补充的位数
	return ciphertext[:length - unpadding]		//裁剪数据
}
