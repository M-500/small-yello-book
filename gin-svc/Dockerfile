FROM golang:1.22 AS builder

ENV APP_HOME /go/pkg/src/github.com/ybook/gin-svc
WORKDIR "$APP_HOME"

COPY go.mod ./
COPY go.sum ./
# RUN go env -w GOPROXY=https://goproxy.cn,direct \
#     && go mod download
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -v -o server main.go && chmod +x server

FROM alpine:latest

COPY --from=builder /go/pkg/src/github.com/ybook/gin-svc /app/bbs-go
COPY --from=builder /go/pkg/src/github.com/ybook/gin-svc/bbs-go.docker.yaml /app/bbs-go.docker.yaml

WORKDIR /app

EXPOSE 8082
ENV ENV=docker
CMD ["./bbs-go"]