# Homework-1.golang

## 作業要求
构建本地镜像。
编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
将镜像推送至 Docker 官方镜像仓库。
通过 Docker 命令本地启动 httpserver。
通过 nsenter 进入容器查看 IP 配置。
作业需编写并提交 Dockerfile 及源代码。

作业提交链接： https://jinshuju.net/f/rxeJhn
提交截止时间：10 月 17 日 23:59

## 作業說明
- build.sh : container 打包
- start.sh : 啟動 container  

## survy
1. 申請 docker hub 帳號
https://hub.docker.com/signup

2. pull golang 基礎鏡向
docker pull golang:1.17.2-alpine3.14

3. 開發 Dockerfile 檔案


