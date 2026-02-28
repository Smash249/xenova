# ============================================
# 阶段1: 安装依赖 + 构建
# ============================================
FROM node:20-alpine AS builder
WORKDIR /app

RUN npm install -g pnpm

# 先复制依赖文件，利用 Docker 缓存
COPY package*.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

# 复制源代码并构建
COPY . .
RUN pnpm run build

# ============================================
# 阶段2: 运行（Nginx）
# ============================================
FROM nginx:1.27-alpine

RUN apk add --no-cache curl tzdata && \
    rm -rf /var/cache/apk/*
RUN ln -snf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo Asia/Shanghai > /etc/timezone
RUN rm -rf /etc/nginx/conf.d/*
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]