# 使用hub.atomgit.com/library/node:20-buster-slim 作为基础镜像
# 由于国内Docker被墙了，所以使用了国内的镜像源 
# 拉取镜像用 docker pull hub.atomgit.com/library/node:20-buster-slim
FROM hub.atomgit.com/library/node:20-buster-slim 

# 设置工作目录
WORKDIR /app

# 复制package.json和package-lock.json
COPY package*.json ./

# 安装依赖
RUN npm install

# 复制项目文件
COPY . .

# 暴露应用运行端口
EXPOSE 8080

# 启动应用
CMD ["npm", "run", "dev"]