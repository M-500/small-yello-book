
# 打包镜像
docker build .  --file=./Dockerfile  --platform=linux/amd64 -t ybook:v1.0.1

## 登陆harbor仓库
#docker login http://124.223.67.129:8081 -u wulinlin
## 推送镜像到远端
#
docker tag ybook:v1.0.1 124.223.67.129:8081/linlin_dev/ybook:v1.0.1
#
#docker push 124.223.67.129:8081/linlin_dev/ybook:v1.0.1