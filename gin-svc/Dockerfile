FROM golang:1.22 AS builder

# 配置工作目录
ENV APP_HOME /go/src/github.com/ybook/gin-svc
WORKDIR "$APP_HOME"

# 复制依赖文件
COPY go.mod ./
COPY go.sum ./
# RUN go env -w GOPROXY=https://goproxy.cn,direct \
#     && go mod download

# 下载依赖
RUN go mod download

# 复制源代码
COPY . ./

# 编译成可执行文件 文件名为server 编译的是main.go 使用详细输出（-v）
RUN CGO_ENABLED=0 go build -v -o ybook-server main.go && chmod +x ybook-server

FROM alpine:latest

COPY --from=builder /go/src/github.com/ybook/gin-svc/ybook-server /app/ybook-server
COPY --from=builder /go/src/github.com/ybook/gin-svc/etc/prod.yaml /app/etc/prod.yaml

WORKDIR /app

EXPOSE 8082

ENV ENV=docker

CMD ["./ybook-server"]