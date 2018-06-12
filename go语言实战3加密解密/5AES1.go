package main

import (
	"bytes"
	"fmt"
)

//AES 对称加密，  128,不足补充0
//129  32  129=32*4+1
func  zeroPadding(ciphertext []byte,blocksize int )[]byte{
	padding:= blocksize- len(ciphertext)%blocksize //填充的长度
	fmt.Println("padding", padding)
	padtext:=bytes.Repeat([]byte{0},padding)//填充0
	return append(ciphertext,padtext...)//字节的叠加
}

func  zerounPadding(origdata []byte)[]byte{ //删除填充数据
	return  bytes.TrimRightFunc(origdata,func (r rune)bool{
		return r==rune(0)
	})
}

func main(){
	mystr:="1234567899"
	newstr:=string(zeroPadding([]byte(mystr),8))
	fmt.Println(newstr)
	fmt.Println(string(zerounPadding([]byte(mystr))))
}