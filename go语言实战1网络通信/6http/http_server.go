package main

import "net/http"

func main() {
	//设置服务器返回信息
	http.HandleFunc("/hellohttp",
		func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("hello http"))
	})
	//开启服务器
	http.ListenAndServe("127.0.0.1:8080", nil)
}
