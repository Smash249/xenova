# Xenova - 工业自动化设备企业官网

<p align="center">
  <img src="public/vite.svg" alt="Xenova Logo" width="120" height="120">
</p>

<p align="center">
  基于 Vue 3 + TypeScript + Vite 构建的现代化企业官网
</p>

<p align="center">
  <a href="#功能特性">功能特性</a> •
  <a href="#技术栈">技术栈</a> •
  <a href="#快速开始">快速开始</a> •
  <a href="#项目结构">项目结构</a> •
  <a href="#脚本命令">脚本命令</a>
</p>

---

## 项目简介

Xenova 是一个专业的工业自动化设备企业官网项目，展示公司的产品中心、企业动态、下载中心等核心业务内容。网站采用现代化的前端技术栈，提供流畅的用户体验和精美的视觉效果。

## 功能特性

- 🏠 **首页展示** - 产品展示、企业优势、服务流程、选择我们等模块
- 📦 **产品中心** - 展示视觉激光打标机、点胶机器人、精密焊接平台等产品
- 📰 **企业动态** - 公司新闻和行业资讯
- 📥 **下载中心** - 产品资料、技术文档下载
- ℹ️ **关于我们** - 企业简介、荣誉资质等

## 技术栈

| 技术 | 说明 |
|------|------|
| [Vue 3](https://vuejs.org/) | 渐进式 JavaScript 框架 |
| [TypeScript](https://www.typescriptlang.org/) | JavaScript 的超集，提供类型支持 |
| [Vite](https://vitejs.dev/) | 下一代前端构建工具 |
| [Vue Router](https://router.vuejs.org/) | Vue.js 官方路由管理器 |
| [Pinia](https://pinia.vuejs.org/) | Vue.js 状态管理库 |
| [Element Plus](https://element-plus.org/) | 基于 Vue 3 的组件库 |
| [TailwindCSS](https://tailwindcss.com/) | 实用优先的 CSS 框架 |
| [GSAP](https://greensock.com/gsap/) | 专业级动画库 |
| [VueUse](https://vueuse.org/) | Vue 组合式 API 工具集 |
| [Lucide Icons](https://lucide.dev/) | 精美的开源图标库 |

## 快速开始

### 环境要求

- Node.js >= 18.0.0
- pnpm >= 8.0.0

### 安装依赖

```bash
# 克隆项目
git clone https://github.com/Smash249/xenova.git

# 进入项目目录
cd xenova

# 安装依赖
pnpm install
```

### 启动开发服务器

```bash
pnpm dev
```

启动后访问 http://localhost:5173

### 构建生产版本

```bash
pnpm build
```

### 预览生产版本

```bash
pnpm preview
```

## 项目结构

```
xenova/
├── public/                 # 静态资源
├── src/
│   ├── assets/            # 项目资源文件
│   ├── components/        # 公共组件
│   │   ├── customBanner/  # 自定义横幅组件
│   │   └── layout/        # 布局组件
│   ├── config/            # 配置文件
│   │   └── header.ts      # 头部导航配置
│   ├── layout/            # 页面布局
│   │   ├── header/        # 头部组件
│   │   ├── footer/        # 底部组件
│   │   └── index.vue      # 布局入口
│   ├── pages/             # 页面组件
│   │   ├── dashboard/     # 首页
│   │   ├── product/       # 产品中心
│   │   ├── news/          # 企业动态
│   │   ├── download/      # 下载中心
│   │   ├── about/         # 关于我们
│   │   └── honor/         # 荣誉资质
│   ├── router/            # 路由配置
│   ├── store/             # 状态管理
│   ├── styles/            # 全局样式
│   ├── types/             # TypeScript 类型定义
│   ├── App.vue            # 根组件
│   └── main.ts            # 入口文件
├── .env.development       # 开发环境变量
├── .env.production        # 生产环境变量
├── index.html             # HTML 入口
├── package.json           # 项目配置
├── tsconfig.json          # TypeScript 配置
├── vite.config.ts         # Vite 配置
└── README.md              # 项目说明
```

## 脚本命令

| 命令 | 说明 |
|------|------|
| `pnpm dev` | 启动开发服务器 |
| `pnpm build` | 构建生产版本 |
| `pnpm preview` | 预览生产构建 |

## 浏览器支持

| Chrome | Firefox | Safari | Edge |
|--------|---------|--------|------|
| ✅ 最新版 | ✅ 最新版 | ✅ 最新版 | ✅ 最新版 |

## 开发建议

### 推荐的 IDE 设置

- [VS Code](https://code.visualstudio.com/)
- [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) - Vue 3 官方插件
- [TypeScript Vue Plugin](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin)

### 代码规范

项目使用 Prettier 进行代码格式化，配置文件为 `prettier.config.cjs`。

## 许可证

本项目遵循 MIT 许可证

---

<p align="center">
  Made with ❤️ by Xenova Team
</p>
