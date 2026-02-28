# Xenova

[中文](#中文) | [English](#english)

---

## 中文

### 项目简介

Xenova 是一个前后端分离的全栈 Web 应用，面向企业官网场景。项目包含三个子模块：

- **backend**：Go 语言编写的 RESTful API 服务
- **frontend**：面向用户的 Vue 3 官网前端
- **admin**：面向管理员的 Vue 3 后台管理系统

提供产品展示、新闻资讯、软件下载、荣誉资质、联系我们等功能模块，并支持用户注册与登录。

### 技术栈

| 层级 | 技术 |
|------|------|
| 后端框架 | [Go](https://go.dev/) + [Echo v5](https://echo.labstack.com/) |
| ORM | [GORM](https://gorm.io/) + MySQL 8 |
| 缓存 | Redis |
| 认证 | JWT (golang-jwt/jwt v5) |
| 配置管理 | Viper (YAML) |
| 前端框架 | [Vue 3](https://vuejs.org/) + TypeScript + Vite |
| UI 组件库（前端） | [Element Plus](https://element-plus.org/) |
| UI 组件库（后台） | [Naive UI](https://www.naiveui.com/) |
| CSS 框架 | [Tailwind CSS v4](https://tailwindcss.com/) |
| 状态管理 | [Pinia](https://pinia.vuejs.org/) |
| HTTP 客户端 | Axios |
| 图表 | ECharts |

### 功能模块

#### 官网前端 (frontend)

- **首页 (Dashboard)**：英雄区、关于我们、业务介绍、产品预览、合作伙伴等
- **产品 (Product)**：产品系列与产品列表、产品详情
- **新闻 (News)**：新闻系列与文章列表、文章详情、浏览量统计
- **下载 (Download)**：软件系列与软件列表
- **荣誉 (Honor)**：公司荣誉、公益活动、专利证书
- **联系我们 (Contact)**：联系信息展示
- **用户认证**：注册、登录、JWT 令牌刷新

#### 后台管理系统 (admin)

- **数据看板**：业务数据概览与图表展示
- **内容管理**：产品、新闻、软件、荣誉等内容的增删改查
- **系统管理**：用户管理与权限控制

### 环境要求

- Go 1.25+
- Node.js 20.19.0+ 或 22.12.0+
- MySQL 8
- Redis 6+

### 快速开始

#### 1. 克隆仓库

```bash
git clone https://github.com/Smash249/xenova.git
cd xenova
```

#### 2. 配置后端

编辑 `backend/config/system.yaml`，填写数据库、Redis、JWT 和邮件服务配置：

```yaml
Router:
  port: 8099

Database:
  mysqlConfig:
    dsn: "root:YOUR_PASSWORD@tcp(127.0.0.1:3306)/xenova?charset=utf8mb4&parseTime=True&loc=Local"
    defaultStringSize: 256
    disableDatetimePrecision: true

Redis:
  host: 127.0.0.1
  port: 6379
  db: 0
  password:

Jwt:
  signingKey: "your-secret-key"
  refresh_exp: 72
  token_exp: 24
  issuer: "xenova"

Email:
  stmp: "smtp.example.com"
  port: "465"
  host: "smtp.example.com"
  pwd: "your-smtp-password"
```

#### 3. 启动后端

```bash
cd backend
go mod download
go run main.go
```

后端默认监听 `http://localhost:8099`

#### 4. 启动官网前端

```bash
cd frontend
npm install
npm run dev
```

官网前端默认运行在 `http://localhost:5173`，API 请求通过 `/api` 前缀代理到后端。

#### 5. 启动后台管理系统

```bash
cd admin
npm install
npm run dev
```

后台管理系统默认运行在 `http://localhost:5799`，API 请求通过 `/api` 前缀代理到后端。

### API 路由概览

#### 公开接口（无需认证）

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/login` | 用户登录 |
| POST | `/api/register` | 用户注册 |
| GET | `/api/health` | 健康检查 |
| GET | `/api/product_series` | 获取产品系列列表 |
| GET | `/api/products` | 获取产品列表 |
| GET | `/api/journalism_series` | 获取新闻系列列表 |
| GET | `/api/journalisms` | 获取新闻列表 |
| GET | `/api/software_series` | 获取软件系列列表 |
| GET | `/api/softwares` | 获取软件列表 |
| GET | `/api/softwares/:software_id` | 获取软件详情 |
| GET | `/api/company_honor` | 获取公司荣誉列表 |
| GET | `/api/love_activity` | 获取公益活动列表 |
| GET | `/api/company_patnet` | 获取公司专利列表 |

#### 私有接口（需要 JWT 认证）

需在请求头中携带 `Authorization: Bearer <token>`。

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/private/products/:product_id` | 获取产品详情 |
| POST/PUT/DELETE | `/api/private/products` | 产品 CRUD |
| POST/PUT/DELETE | `/api/private/product_series` | 产品系列 CRUD |
| GET | `/api/private/journalisms/:journalism_id` | 获取新闻详情 |
| POST/PUT/DELETE | `/api/private/journalisms` | 新闻 CRUD |
| POST/PUT/DELETE | `/api/private/journalism_series` | 新闻系列 CRUD |
| POST/PUT/DELETE | `/api/private/softwares` | 软件 CRUD |
| POST/PUT/DELETE | `/api/private/software_series` | 软件系列 CRUD |
| POST/PUT/DELETE | `/api/private/company_honor` | 公司荣誉 CRUD |
| POST/PUT/DELETE | `/api/private/love_activity` | 公益活动 CRUD |
| POST/PUT/DELETE | `/api/private/company_patent` | 公司专利 CRUD |

### 目录结构

```
xenova/
├── backend/                # Go 后端 API 服务
│   ├── config/             # 配置文件 (system.yaml)
│   ├── initialize/         # 初始化 (Viper, GORM, Redis, Router)
│   ├── internal/
│   │   ├── api/            # 控制器层
│   │   ├── global/         # 全局变量 & 基础模型
│   │   ├── models/         # 数据库模型
│   │   └── service/        # 业务逻辑层
│   ├── pkg/                # 工具包
│   ├── router/             # 路由注册
│   ├── utils/              # 通用工具函数
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
├── frontend/               # Vue 3 官网前端 (端口 5173)
│   ├── src/
│   │   ├── api/            # Axios 请求封装
│   │   ├── components/     # 公共组件
│   │   ├── layout/         # 页面布局
│   │   ├── pages/          # 页面视图
│   │   ├── router/         # 路由配置
│   │   ├── store/          # Pinia 状态
│   │   ├── styles/         # 全局样式
│   │   ├── types/          # TypeScript 类型
│   │   └── utils/          # 工具函数
│   ├── Dockerfile
│   └── package.json
├── admin/                  # Vue 3 后台管理系统 (端口 5799)
│   ├── src/
│   │   ├── api/            # Axios 请求封装
│   │   ├── components/     # 公共组件
│   │   ├── composables/    # 组合式函数
│   │   ├── layout/         # 页面布局
│   │   ├── router/         # 路由配置
│   │   ├── stores/         # Pinia 状态
│   │   ├── theme/          # 主题配置
│   │   ├── types/          # TypeScript 类型
│   │   ├── utils/          # 工具函数
│   │   └── views/          # 页面视图
│   ├── Dockerfile
│   ├── nginx.conf
│   └── package.json
├── DOCKER.md               # Docker 部署指南
└── README.md
```

### Docker 部署

请参阅 [DOCKER.md](./DOCKER.md) 获取完整的 Docker 和 Docker Compose 部署说明。

---

## English

### Overview

Xenova is a full-stack web application for enterprise websites. The project includes three submodules:

- **backend**: RESTful API service written in Go
- **frontend**: User-facing Vue 3 website frontend
- **admin**: Vue 3 admin panel for content management

It provides modules for product showcase, news, software downloads, company honors, contact information, and user authentication.

### Tech Stack

| Layer | Technology |
|-------|------------|
| Backend Framework | [Go](https://go.dev/) + [Echo v5](https://echo.labstack.com/) |
| ORM | [GORM](https://gorm.io/) + MySQL 8 |
| Cache | Redis |
| Authentication | JWT (golang-jwt/jwt v5) |
| Configuration | Viper (YAML) |
| Frontend Framework | [Vue 3](https://vuejs.org/) + TypeScript + Vite |
| UI Components (frontend) | [Element Plus](https://element-plus.org/) |
| UI Components (admin) | [Naive UI](https://www.naiveui.com/) |
| CSS Framework | [Tailwind CSS v4](https://tailwindcss.com/) |
| State Management | [Pinia](https://pinia.vuejs.org/) |
| HTTP Client | Axios |
| Charts | ECharts |

### Features

#### Website Frontend (frontend)

- **Dashboard**: Hero section, About Us, Business overview, Product preview, Partners
- **Products**: Product series & list, product details
- **News**: News series & articles, article detail, view count
- **Downloads**: Software series & software list
- **Honor**: Company honors, public welfare activities, patents
- **Contact**: Contact information
- **Authentication**: Register, login, JWT token refresh

#### Admin Panel (admin)

- **Dashboard**: Business data overview and charts
- **Content Management**: CRUD for products, news, software, honors, etc.
- **System Management**: User management and access control

### Prerequisites

- Go 1.25+
- Node.js 20.19.0+ or 22.12.0+
- MySQL 8
- Redis 6+

### Quick Start

#### 1. Clone the repository

```bash
git clone https://github.com/Smash249/xenova.git
cd xenova
```

#### 2. Configure the backend

Edit `backend/config/system.yaml` with your database, Redis, JWT, and email settings:

```yaml
Router:
  port: 8099

Database:
  mysqlConfig:
    dsn: "root:YOUR_PASSWORD@tcp(127.0.0.1:3306)/xenova?charset=utf8mb4&parseTime=True&loc=Local"
    defaultStringSize: 256
    disableDatetimePrecision: true

Redis:
  host: 127.0.0.1
  port: 6379
  db: 0
  password:

Jwt:
  signingKey: "your-secret-key"
  refresh_exp: 72
  token_exp: 24
  issuer: "xenova"

Email:
  stmp: "smtp.example.com"
  port: "465"
  host: "smtp.example.com"
  pwd: "your-smtp-password"
```

#### 3. Start the backend

```bash
cd backend
go mod download
go run main.go
```

The backend listens on `http://localhost:8099` by default.

#### 4. Start the website frontend

```bash
cd frontend
npm install
npm run dev
```

The frontend runs on `http://localhost:5173` by default. API requests with the `/api` prefix are proxied to the backend.

#### 5. Start the admin panel

```bash
cd admin
npm install
npm run dev
```

The admin panel runs on `http://localhost:5799` by default. API requests with the `/api` prefix are proxied to the backend.

### Project Structure

```
xenova/
├── backend/                # Go backend API service
│   ├── config/             # Configuration (system.yaml)
│   ├── initialize/         # App bootstrap (Viper, GORM, Redis, Router)
│   ├── internal/
│   │   ├── api/            # Handler / controller layer
│   │   ├── global/         # Global variables & base model
│   │   ├── models/         # Database models
│   │   └── service/        # Business logic layer
│   ├── pkg/                # Shared packages
│   ├── router/             # Route registration
│   ├── utils/              # Utility functions
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
├── frontend/               # Vue 3 website frontend (port 5173)
│   ├── src/
│   │   ├── api/            # Axios request wrappers
│   │   ├── components/     # Shared components
│   │   ├── layout/         # Page layout
│   │   ├── pages/          # Page views
│   │   ├── router/         # Vue Router config
│   │   ├── store/          # Pinia stores
│   │   ├── styles/         # Global styles
│   │   ├── types/          # TypeScript types
│   │   └── utils/          # Utility functions
│   ├── Dockerfile
│   └── package.json
├── admin/                  # Vue 3 admin panel (port 5799)
│   ├── src/
│   │   ├── api/            # Axios request wrappers
│   │   ├── components/     # Shared components
│   │   ├── composables/    # Composable functions
│   │   ├── layout/         # Page layout
│   │   ├── router/         # Vue Router config
│   │   ├── stores/         # Pinia stores
│   │   ├── theme/          # Theme configuration
│   │   ├── types/          # TypeScript types
│   │   ├── utils/          # Utility functions
│   │   └── views/          # Page views
│   ├── Dockerfile
│   ├── nginx.conf
│   └── package.json
├── DOCKER.md               # Docker deployment guide
└── README.md
```

### Docker Deployment

See [DOCKER.md](./DOCKER.md) for full Docker and Docker Compose deployment instructions.

### License

This project is for learning and reference purposes. Please refer to the repository for license details.
