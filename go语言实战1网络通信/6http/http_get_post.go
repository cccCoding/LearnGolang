package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var url = "http://www.baidu.com"
	resp,err := http.Get(url)
	//resp,err := http.Post(url,
	//	"application/x-www-form-urlencoded",
	//		strings.NewReader("id=1"))
	if err != nil {
		fmt.Println("错误")
	}
	defer resp.Body.Close()
	body,_ := ioutil.ReadAll(resp.Body)	//读取信息
	fmt.Println(string(body))
}


