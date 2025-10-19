# GitHub Actions 构建说明

这个项目使用 GitHub Actions 自动在各平台的原生环境中构建应用程序。

## ✨ 特性

- ✅ **多平台原生构建**：macOS、Windows、Linux
- ✅ **多架构支持**：amd64、arm64、universal (macOS)
- ✅ **自动 UPX 压缩**：减小可执行文件体积
- ✅ **依赖缓存**：加速构建过程
- ✅ **自动发布**：Tag 推送时自动创建 Release
- ✅ **SHA256 校验和**：确保下载文件完整性
- ✅ **构建摘要**：清晰显示所有平台构建状态

## 🚀 使用方法

### 方式 1：推送 Tag 触发构建和发布

```bash
# 创建 tag
git tag v1.0.0

# 推送 tag 到 GitHub
git push origin v1.0.0
```

工作流会自动：
1. 在 macOS、Windows、Linux 上构建所有架构
2. 使用 UPX 压缩所有可执行文件
3. 生成 SHA256 校验和
4. 创建 GitHub Release 并上传所有构建产物

### 方式 2：手动触发构建

1. 访问 GitHub 仓库页面
2. 点击 **Actions** 标签
3. 选择 **Build Multi-Platform** 工作流
4. 点击 **Run workflow**
5. （可选）输入版本号
6. 点击 **Run workflow** 确认

构建完成后，可以在 Actions 页面下载各平台的构建产物。

### 方式 3：推送代码触发测试构建

推送到 `main`、`master` 或 `develop` 分支：

```bash
git push origin main
```

这会触发构建但不会创建 Release，构建产物会保留 7 天供下载。

## 📦 构建产物

每次构建会生成以下文件：

### macOS
- `jetbrains-darwin-amd64.zip` - Intel Mac
- `jetbrains-darwin-arm64.zip` - Apple Silicon (M1/M2/M3)
- `jetbrains-darwin-universal.zip` - Universal (支持所有 Mac)

### Windows
- `jetbrains-windows-amd64.zip` - x64
- `jetbrains-windows-arm64.zip` - ARM64

### Linux
- `jetbrains-linux-amd64.tar.gz` - x64
- `jetbrains-linux-arm64.tar.gz` - ARM64

### 校验和
- `checksums.txt` - 所有文件的 SHA256 校验和

## 🔧 工作流配置

### 触发条件

- **Tag 推送** (`v*`): 构建并创建 Release
- **代码推送** (main/master/develop): 仅构建，不创建 Release
- **Pull Request**: 构建测试
- **手动触发**: 随时构建

### 环境变量

可以在 `.github/workflows/build.yml` 中修改：

```yaml
env:
  GO_VERSION: '1.21'      # Go 版本
  NODE_VERSION: '18'      # Node.js 版本
  APP_NAME: 'jetbrains'   # 应用名称
```

### 构建选项

每个平台的构建都使用以下选项：

- `-upx`: 使用 UPX 压缩
- `-clean`: 清理构建缓存

## 📊 构建状态

每次构建完成后，会在 Actions Summary 中显示所有平台的构建状态：

```
🏗️ Build Summary

| Platform | Status  |
|----------|---------|
| macOS    | success |
| Windows  | success |
| Linux    | success |

Triggered by: push
Ref: refs/tags/v1.0.0
```

## ⚡ 加速构建

工作流使用了多种缓存策略加速构建：

1. **Go 模块缓存**: 缓存 Go 依赖
2. **npm 缓存**: 缓存前端依赖
3. **并行构建**: 所有平台和架构并行构建

## 🛠️ 故障排查

### 构建失败

1. 检查 Actions 日志查看具体错误
2. 确保 `frontend/package-lock.json` 存在
3. 验证 Go 和 Node.js 版本兼容性

### Release 未创建

- 确保推送的是以 `v` 开头的 tag
- 检查仓库的 Actions 权限设置

### 下载文件损坏

- 使用 `checksums.txt` 验证文件完整性：

```bash
sha256sum -c checksums.txt
```

## 📝 版本号规范

推荐使用语义化版本号：

- `v1.0.0` - 正式版本
- `v1.0.0-beta.1` - Beta 版本
- `v1.0.0-rc.1` - Release Candidate

## 🔒 安全性

- 使用官方 GitHub Actions
- 最小化权限（仅 `contents: write`）
- 所有构建在 GitHub 托管的 runner 上执行
- 提供 SHA256 校验和验证

## 📚 相关文档

- [BUILD.md](../BUILD.md) - 本地构建说明
- [Wails 文档](https://wails.io/docs/introduction)
- [GitHub Actions 文档](https://docs.github.com/en/actions)
