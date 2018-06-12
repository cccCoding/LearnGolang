package main

import (
	"encoding/pem"
	"errors"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"fmt"
	"encoding/base64"
)

var privatekey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDYdrBJazD5XHHSUmTtg4DL7kN4745GDIOewCpM1KnUomrHOTLj
DdRHq/yXU+C8+Xmb6hp5CMwmdkxeYWfFiOhOS4j+OVKdz/6kwEloGyhmGfvU3Fjo
bifECwDyOAoUJP/Rbk++G7KVT0kRo8P5WM46i8wDPq4i5AiRivV1SEe+ywIDAQAB
AoGBALTFsgUuYpDtFhUqQCVmnAoy6eA2Vx/C1ayfGPRe6ZGtLfVAnHGoG+7a/7A1
GUtYIKoHwKHxqeQ5CSAMwofiBPXMyF6dsA3xk890CE14FzgzYsASv7beFUyC0WlR
Sw6NVRDgnCsVUHyH6ncw5la+Q2strPDUElpLzRa0Y7auYM/ZAkEA6hU+zaIT2+qx
4idqSbi2y4CSbzy+vXUUpRtTC2+e277Ii6g95n13hA3W50OBE5V5KHiWSyvNPR9d
SoFe9EOpDQJBAOy7H5L8dbh6t1kmbsREH2JWXdL/0cigfLjoVFFyFonYqplyn72i
yLDjXRJTTAO1ftscA7WKsjLhikWiU/BS4TcCQAvbt6rDIy5o3UoPpRrG+Lumb8Si
1ybR35HdqH9T0EBhkddBVqFuibdu3AwrJ0bOs6yRL0vvlB4ckKVNmHnXU6ECQE0p
yWbt56lURsNGZcCPu/Mf18FZJZZRyZW0FRffKj2QZDtUQ4FauDQASGwavqJO3KSr
AwJ/zQoMvjdobBffzA0CQBeb6uY1GCnI3I70lOI1jdk5L4Ubl4lSxHaGUnQ6YcPR
UkeJ3kMQJAF0/c5wsPWTEusYYTz0xeh+rUcVQ+uAZkg=
-----END RSA PRIVATE KEY-----
`)
var publickey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDYdrBJazD5XHHSUmTtg4DL7kN4
745GDIOewCpM1KnUomrHOTLjDdRHq/yXU+C8+Xmb6hp5CMwmdkxeYWfFiOhOS4j+
OVKdz/6kwEloGyhmGfvU3FjobifECwDyOAoUJP/Rbk++G7KVT0kRo8P5WM46i8wD
Pq4i5AiRivV1SEe+ywIDAQAB
-----END PUBLIC KEY-----
`)

func main() {
	data,err := RSAEncrypt([]byte("mima"))
	if err != nil {
		fmt.Println("错误",err)
	}
	fmt.Println("加密",base64.StdEncoding.EncodeToString(data))

	origData,err := RSADecrypt(data)
	if err != nil {
		fmt.Println("错误",err)
	}
	fmt.Println("解密",string(origData))
}

func RSAEncrypt(origData []byte) ([]byte, error) {
	block,_ := pem.Decode(publickey)
	if block == nil {
		return nil,errors.New("public key is bad")
	}
	publicInterface,err := x509.ParsePKIXPublicKey(block.Bytes)//获取公钥
	if err != nil {
		return nil,err
	}
	pub := publicInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader,pub,origData)	//加密
}

func RSADecrypt(cipherText []byte) ([]byte, error) {
	block,_ := pem.Decode(privatekey)
	if block == nil {
		return nil,errors.New("private key is bad")
	}
	priv,err := x509.ParsePKCS1PrivateKey(block.Bytes)//获取私钥
	if err != nil {
		return nil,err
	}
	return rsa.DecryptPKCS1v15(rand.Reader,priv,cipherText)
}
