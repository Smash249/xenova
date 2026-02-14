# Docker 使用指南 / Docker Usage Guide

本项目已为后端和前端添加 Dockerfile，可以使用 Docker 容器化部署应用。

This project includes Dockerfiles for both backend and frontend for containerized deployment.

## 后端 / Backend

### 构建镜像 / Build Image

```bash
cd backend
docker build -t xenova-backend .
```

### 运行容器 / Run Container

```bash
docker run -d \
  -p 8099:8099 \
  --name xenova-backend \
  xenova-backend
```

### 配置说明 / Configuration

后端配置文件位于 `config/system.yaml`。在生产环境中，建议：
- 使用环境变量或 Docker secrets 管理敏感信息
- 将配置文件作为 volume 挂载
- 更新数据库和 Redis 连接为容器化服务

Backend configuration is in `config/system.yaml`. For production:
- Use environment variables or Docker secrets for sensitive data
- Mount config file as a volume
- Update database and Redis connections for containerized services

示例 / Example:
```bash
docker run -d \
  -p 8099:8099 \
  -v /path/to/your/config.yaml:/app/config/system.yaml \
  --name xenova-backend \
  xenova-backend
```

## 前端 / Frontend

### 构建镜像 / Build Image

```bash
cd frontend
docker build -t xenova-frontend .
```

### 运行容器 / Run Container

```bash
docker run -d \
  -p 80:80 \
  --name xenova-frontend \
  xenova-frontend
```

或指定其他端口 / Or specify a different port:
```bash
docker run -d \
  -p 3000:80 \
  --name xenova-frontend \
  xenova-frontend
```

## Docker Compose (推荐 / Recommended)

创建 `docker-compose.yml` 文件可以同时运行后端和前端：

Create a `docker-compose.yml` file to run both backend and frontend together:

```yaml
version: '3.8'

services:
  backend:
    build: ./backend
    ports:
      - "8099:8099"
    environment:
      # 根据需要添加环境变量 / Add environment variables as needed
      - DATABASE_HOST=mysql
      - REDIS_HOST=redis
    depends_on:
      - mysql
      - redis
    restart: unless-stopped

  frontend:
    build: ./frontend
    ports:
      - "80:80"
    depends_on:
      - backend
    restart: unless-stopped

  mysql:
    image: mysql:8
    environment:
      MYSQL_ROOT_PASSWORD: your_password
      MYSQL_DATABASE: xenova
    volumes:
      - mysql_data:/var/lib/mysql
    restart: unless-stopped

  redis:
    image: redis:alpine
    restart: unless-stopped

volumes:
  mysql_data:
```

使用 Docker Compose 启动 / Start with Docker Compose:
```bash
docker-compose up -d
```

停止服务 / Stop services:
```bash
docker-compose down
```

## 注意事项 / Notes

1. **后端 Dockerfile** 使用多阶段构建，生成的镜像体积小，仅包含运行必需的二进制文件
   
   **Backend Dockerfile** uses multi-stage builds for a smaller image with only the necessary binary

2. **前端 Dockerfile** 使用 Nginx 作为 Web 服务器，适合生产环境部署
   
   **Frontend Dockerfile** uses Nginx as the web server, suitable for production deployment

3. 确保在生产环境中更新所有默认密码和密钥
   
   Make sure to update all default passwords and keys in production

4. 建议使用 `.dockerignore` 文件排除不必要的文件以加快构建速度
   
   Use `.dockerignore` files to exclude unnecessary files for faster builds
