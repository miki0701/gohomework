# Homework-1.golang

## 作業要求

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出 
4. 当访问 localhost/healthz 时，应返回200



## survy

1. 要求一
   - [x] 如何讀取 request header
   - [x] 如何寫入 response header
2. 要求二
   - [x] 如何讀取環境變量
     - [Go by Example: Environment Variables](https://gobyexample.com/environment-variables)
3. 要求三
   - [x] 如何讀取 client ip
   - [ ] 如何 log 訪問日誌
			- [How to Collect, Standardize, and Centralize Golang Logs | Datadog](https://www.datadoghq.com/blog/go-logging/)
      - [An example of how to use golang/glog.](https://gist.github.com/heatxsink/7221ebe499b0767d4784)


4. 要求四
   - [x] 如何回應 http 200
     - [http package - net/http - pkg.go.dev](https://pkg.go.dev/net/http#StatusOK)
     - [Go HTTP Header 的那些坑 | Golang](https://tachingchen.com/tw/blog/pitfall-of-golang-header-operation/)
     - [golang http.ResponseWriter用法及代碼示例 - 純淨天空](https://vimsky.com/zh-tw/examples/usage/golang_net_http_ResponseWriter.html)


## 參考網站
- geektutu/7days-golang: 7 days golang programs from scratch (web framework Gee, distributed cache GeeCache, object relational mapping ORM framework GeeORM, rpc framework GeeRPC etc) 7天用Go动手写/从零实现系列
  https://github.com/geektutu/7days-golang
- Go by Example
  https://gobyexample.com/
- http package - net/http - pkg.go.dev
  https://pkg.go.dev/net/http
- web 工作方式 - 使用 Golang 打造 Web 應用程式
  https://willh.gitbook.io/build-web-application-with-golang-zhtw/03.0/03.1


## 擴展需求
1. 如何透過 goroutine 發動 reqeust
2. 如何透過 goroutine 調用 shell command

