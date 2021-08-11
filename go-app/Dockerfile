# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.14 as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .

RUN GOOS=linux GOARCH=amd64 go build .

FROM alpine

WORKDIR /app

RUN go build .

EXPOSE 8888

ENTRYPOINT ["./main"]