package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "Hello, world !")
}

func main() {
	//分发请求,即针对某一路径请求将其映射到指定的业务逻辑处理方法中
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe：", err.Error())
	}
}
