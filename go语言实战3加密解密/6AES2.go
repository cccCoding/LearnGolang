package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"encoding/base64"
)

func main() {
	fmt.Println("start")
	//AES128 ,16,24,32,AES128,192,256
	key := []byte("hahahahahahahaha")	//16位密钥
	result,err := AESEncrypt([]byte("wodemima"),key)	//加密
	if err != nil {
		//panic(err)		//处理异常
		fmt.Println("加密出错",err)
	}
	fmt.Println("加密后的数据", base64.StdEncoding.EncodeToString(result))	//加密后的数据
	origdata,err := AESDecrypt(result,key)	//解密
	if err != nil {
		//panic(err)
		fmt.Println("解密出错",err)
	}
	fmt.Println("解密后的数据", string(origdata))
}

//AES 对称加密，  128,不足补充0
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

//加密
func AESEncrypt(origdata, key []byte) ([]byte, error) {
	block,err := aes.NewCipher(key)		//创建密码根据密码加密
	if err != nil {
		return nil, err
	}
	blocksize := block.BlockSize()	//定义大小
	//origdata=PKCS5Padding(origdata,blocksize)跨语言
	origdata = zeroPadding(origdata, blocksize)	//补充位数
	blockmode := cipher.NewCBCEncrypter(block, key[:blocksize])	//创建加密算法
	cyphed := make([]byte,len(origdata))	//加密
	blockmode.CryptBlocks(cyphed, origdata)	//按照区块加密
	return cyphed,nil
}

//解密
func AESDecrypt(cryped, key []byte) ([]byte, error) {
	block,err := aes.NewCipher(key)
	if err != nil {
		return nil,err
	}
	blocksize := block.BlockSize()
	blockmode := cipher.NewCBCDecrypter(block,key[:blocksize])
	origData := make([]byte,len(cryped))
	blockmode.CryptBlocks(origData,cryped)	//解密
	origData = zerounPadding(origData)		//去掉尾部补充的位数
	return origData,nil
}