FROM golang:alpine as builder
# 需要go环境
MAINTAINER hujie

WORKDIR /smoke

# 源
RUN go env -w GOPROXY=https://goproxy.cn,direct && go env -w CGO_ENABLED=0
COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o main main.go


FROM alpine:latest
# 设置时区
RUN apk add --no-cache tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" >  /etc/timezone

WORKDIR /server
# 复制到工作区
COPY --from=builder /smoke/ ./
# COPY --from=builder /work/config ./config
# 对外端口
EXPOSE 8080
# 执行
CMD ["./main"]