# JetBrains 工具

一个跨平台的 JetBrains IDE 管理工具，支持 Windows、macOS 和 Linux。

## 功能特性

- 自动扫描已安装的 JetBrains 产品
- 支持下载 JetBrains 产品安装包
- 清理第三方工具残留的环境变量配置
- 清理第三方工具残留的vmoption配置
- 跨平台支持（Windows / macOS / Linux）

## 环境要求

- Go 1.23+
- Node.js 18+
- pnpm
- Wails CLI

## 安装 Wails CLI

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

## 本地开发

1. 修改配置
    编辑 `wails.json`，将 `obfuscated` 设置为 `false`：
    
    ```json
    {
      "obfuscated": false
    }
    ```
2. 修改前端文件去除授权验证
3. 安装依赖并启动

```bash
# 安装前端依赖
cd frontend && pnpm install && cd ..

# 启动开发模式
wails dev
```

## 构建

```bash
# 构建当前平台
wails build

# 构建指定平台
wails build -platform windows/amd64
wails build -platform darwin/amd64
wails build -platform darwin/arm64
wails build -platform linux/amd64
```

构建产物位于 `build/bin/` 目录。

## 项目结构

```
.
├── main.go              # 应用入口
├── app.go               # 主应用逻辑 (Wails 绑定)
├── util/                # 工具库
│   ├── device.go        # 设备信息获取
│   ├── product.go       # 产品管理逻辑
│   ├── file.go          # 文件操作
│   └── work.go          # 工作流程
├── global/              # 全局配置
├── frontend/            # 前端应用
│   ├── src/
│   │   ├── views/       # 页面组件
│   │   ├── store/       # 状态管理
│   │   ├── api/         # API 接口
│   │   └── router/      # 路由配置
│   └── package.json
├── build/               # 构建配置和产物
└── wails.json           # Wails 配置
```

## 注意事项

- macOS 用户需要授予磁盘访问权限
- 操作前请确保 JetBrains IDE 已关闭

## 许可证

本项目采用 [CC BY-NC 4.0](LICENSE) 许可证，禁止商业使用。
- 本项目并不支持开箱即用，需自行下载jar包等资源
- 本仓库开源是为了学习和交流目的，请勿二次开发并且用于商业用途