# gin-svc


## Description


## Installation

```bash
$ cd gin-svc
## 下载依赖
$ go mod tidy
## 启动
$ go run main.go
```


## Usage

```bash
$ go run main.go
```

```shell
swag init
## swagger 文档地址
http://127.0.0.1:8122/na/swagger/index.html
```


如何部署
```go
docker build .  --file=./Dockerfile  --platform=linux/amd64 -t ybook:v1.0.1
```

## License

