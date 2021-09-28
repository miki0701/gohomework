package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("start server...")

	// 1.接收客户端 request，并将 request 中带的 header 写入 response header
	http.HandleFunc("/cloneheader", cloneheader)
	// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	http.HandleFunc("/readenv", readenv)
	// 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	http.HandleFunc("/logreq", logreq)
	// 4.当访问 localhost/healthz 时，应返回200
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		fmt.Println("error...")
	}
}

// curl -i -H "user: 123" -H "x-forwarded-for:192.168.100.100" http://127.0.0.1/cloneheader
func cloneheader(w http.ResponseWriter, r *http.Request) {
	// fowarded := r.Header.Get("x-forwarded-for")
	// w.Header().Set("x-forwarded-for", fowarded)
	for name, value := range r.Header {
		fmt.Printf("header:%s value:%s\n", name, value)
		w.Header().Set(name, strings.Join(value, ","))
	}
	io.WriteString(w, "cloneheader done!")
}

// curl -i http://127.0.0.1/readenv
func readenv(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	// WriteHeader 必須在 Header().Set 之後
	w.WriteHeader(http.StatusOK)
}

// curl -i http://127.0.0.1/logreq
func logreq(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	user_agent := r.Header.Get("User-Agent")
	//TODO: 改用 glog/klog
	fmt.Printf("ip:%s user_agent:%s  \n", ip, user_agent)
	io.WriteString(w, "logreq done!")
}

// curl -i http://127.0.0.1/healthz
func healthz(w http.ResponseWriter, r *http.Request) {
	// 只回應 http status code 200
	w.WriteHeader(http.StatusOK)
}