# 构建说明

## 跨平台编译限制

**重要提示：** Wails 在不同平台上的交叉编译支持有限：

- ✅ **macOS**: 可以构建 macOS (amd64/arm64/universal) 版本
- ⚠️ **Windows**: 从 macOS 交叉编译到 Windows 可能失败（CGO 和 UPX 问题）
- ❌ **Linux**: **不支持**从 macOS 交叉编译到 Linux

### 推荐的构建方式

#### 方式一：使用 GitHub Actions（推荐）✨

项目配置了完整的 CI/CD 工作流，支持自动构建和发布：

##### 📦 完整构建（build.yml）

自动在各平台的原生环境中构建所有架构：

```bash
# 推送 tag 触发自动构建和发布
git tag v1.0.0
git push origin v1.0.0

# 或在 GitHub 网页手动触发
# Actions → Build Multi-Platform → Run workflow
```

**特性：**
- ✅ 所有平台所有架构并行构建
- ✅ 自动 UPX 压缩
- ✅ 生成 SHA256 校验和
- ✅ 自动创建 GitHub Release
- ✅ 依赖缓存加速构建

**构建产物：**
- macOS: amd64、arm64、universal
- Windows: amd64、arm64
- Linux: amd64、arm64
- 所有文件的 SHA256 校验和

##### ⚡ 快速检查（check.yml）

PR 时自动运行代码检查和快速测试：

- 代码 Lint（Go + 前端）
- 单元测试
- 当前平台构建测试

**详细说明**: [.github/workflows/README.md](.github/workflows/README.md)

#### 方式二：使用 Makefile 在本地构建

```bash
# 在 macOS 上只构建 macOS 版本
make mac_amd64
make mac_arm64
make mac_universal

# 查看所有可用目标
make help
```

#### 方式三：使用 Docker 构建 Linux 版本

```bash
# 使用 Linux 容器构建
docker run --rm -v $(pwd):/app -w /app ubuntu:22.04 bash -c "
  apt-get update && \
  apt-get install -y golang nodejs npm build-essential libgtk-3-dev libwebkit2gtk-4.1-dev upx && \
  go install github.com/wailsapp/wails/v2/cmd/wails@latest && \
  wails build -platform linux/amd64 -upx
"
```

## 单平台构建命令

### macOS 平台

```bash
# Apple Silicon (M1/M2/M3)
wails build -platform darwin/arm64 -upx

# Intel
wails build -platform darwin/amd64 -upx

# Universal (支持所有 Mac)
wails build -platform darwin/universal -upx
```

### Windows 平台（需在 Windows 系统上构建）

```bash
# x64
wails build -platform windows/amd64 -upx

# ARM64
wails build -platform windows/arm64 -upx
```

### Linux 平台（需在 Linux 系统上构建）

```bash
# AMD64
wails build -platform linux/amd64 -upx

# ARM64
wails build -platform linux/arm64 -upx
```

## 使用 Makefile 构建

```bash
# 清理构建目录
make clean

# 构建 macOS 所有架构
make mac_amd64
make mac_arm64
make mac_universal

# 构建并打包（生成 .dmg）
make package_mac

# 查看帮助
make help
```

## 输出文件

构建完成后，可执行文件将生成在 `build/bin/` 目录下：

- **macOS**: `jetbrains.app` (免安装，直接打开)
- **Windows**: `jetbrains.exe` (免安装，直接运行)
- **Linux**: `jetbrains` (免安装，添加执行权限后运行)

## 构建选项说明

### UPX 压缩（推荐）
```bash
wails build -platform <platform>/<arch> -upx
```
使用 UPX 压缩可执行文件，大幅减小体积。

### 清理重新构建
```bash
wails build -platform <platform>/<arch> -clean
```
清理缓存后重新构建。

### 跳过前端构建
```bash
wails build -platform <platform>/<arch> -s
```
跳过前端构建，加快编译速度（用于后端修改）。

## 开发模式

```bash
# 运行开发服务器
wails dev

# 或使用 Makefile
make dev
```

## 依赖安装

### macOS
```bash
# 安装 Wails
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# 安装 UPX（可选，用于压缩）
brew install upx
```

### Windows
```bash
# 安装 Wails
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# 安装 UPX
choco install upx
```

### Linux
```bash
# 安装依赖
sudo apt-get update
sudo apt-get install -y build-essential libgtk-3-dev libwebkit2gtk-4.1-dev upx

# 安装 Wails
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

## 常见问题

### Q: 为什么不能从 macOS 构建 Linux 版本？
A: Wails 依赖于系统原生的 WebView 组件（Linux 上是 WebKit2GTK），无法在 macOS 上交叉编译。建议使用 GitHub Actions 或 Docker。

### Q: Windows 构建时出现 "unknown ARM64 relocation type" 错误？
A: 这是从 macOS ARM64 交叉编译到 Windows 的已知问题。解决方案：
1. 使用 GitHub Actions 在 Windows runner 上构建
2. 或在 Windows 系统上直接构建

### Q: 如何减小可执行文件体积？
A: 使用 `-upx` 参数启用 UPX 压缩，可以减小 50-70% 的体积。