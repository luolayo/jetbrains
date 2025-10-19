# 应用程序名称和版本信息
APP_NAME := jetbrains
VERSION := 1.0.0
GIT_HASH := $(shell git rev-parse --short HEAD 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date +%Y-%m-%dT%H:%M:%S)
IDENTIFIER := jb.luola.me
PKG_OUTPUT := build/$(APP_NAME)-$(VERSION)

# 编译标记
LDFLAGS := -ldflags="-X 'main.AppName=$(APP_NAME)' -X 'main.Version=$(VERSION)' -X 'main.GitHash=$(GIT_HASH)' -X 'main.BuildTime=$(BUILD_TIME)'"

# Wails 命令
WAILS := wails

# 默认目标
.PHONY: all
all: clean build-native package_mac

# 清理旧的构建文件
.PHONY: clean
clean:
	@echo "Cleaning build directory..."
	rm -rf build

# 生成静态文件（如需要）
.PHONY: static-generated
static-generated:
	statik -src=./configs

# 运行开发模式
.PHONY: dev
dev:
	$(WAILS) dev -s

# 各个平台和架构的构建目标
.PHONY: build
build: build-native

# macOS 原生构建（推荐在 macOS 上使用）
.PHONY: build-native
build-native: mac_amd64 mac_arm64 mac_universal

# 全平台构建（仅用于 CI/CD，本地跨平台编译会失败）
.PHONY: build-all
build-all: mac_amd64 mac_arm64 mac_universal windows_amd64 windows_arm64 linux_amd64 linux_arm64
	@echo ""
	@echo "警告: 从 macOS 跨平台构建 Windows/Linux 可能失败"
	@echo "推荐使用 GitHub Actions 进行多平台构建"
	@echo "详见 BUILD.md 文档"

# 构建 macOS 版本
.PHONY: mac_amd64
mac_amd64:
	@echo "Building macOS AMD64 version with Git Hash: $(GIT_HASH)"
	GOOS=darwin GOARCH=amd64 $(WAILS) build -platform darwin/amd64  $(LDFLAGS) -o build/$(APP_NAME)-macOS-AMD64

.PHONY: mac_arm64
mac_arm64:
	@echo "Building macOS ARM64 version with Git Hash: $(GIT_HASH)"
	GOOS=darwin GOARCH=arm64 $(WAILS) build -platform darwin/arm64  $(LDFLAGS) -o build/$(APP_NAME)-macOS-ARM64

.PHONY: mac_universal
mac_universal:
	@echo "Building macOS Universal version (AMD64 + ARM64) with Git Hash: $(GIT_HASH)"
	GOOS=darwin GOARCH=arm64 $(WAILS) build -platform darwin/universal  $(LDFLAGS) -o build/$(APP_NAME)-macOS-Universal

# 构建 Windows 版本 (禁用 CGO 以支持交叉编译)
.PHONY: windows_amd64
windows_amd64:
	@echo "Building Windows AMD64 version with Git Hash: $(GIT_HASH)"
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(WAILS) build -platform windows/amd64  $(LDFLAGS) -o build/$(APP_NAME)-Windows-AMD64.exe

.PHONY: windows_arm64
windows_arm64:
	@echo "Building Windows ARM64 version with Git Hash: $(GIT_HASH)"
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 $(WAILS) build -platform windows/arm64  $(LDFLAGS) -o build/$(APP_NAME)-Windows-ARM64.exe

# 构建 Linux 版本
.PHONY: linux_amd64
linux_amd64:
	@echo "Building Linux AMD64 version with Git Hash: $(GIT_HASH)"
	GOOS=linux GOARCH=amd64 $(WAILS) build -platform linux/amd64 -upx $(LDFLAGS) -o build/$(APP_NAME)-Linux-AMD64

.PHONY: linux_arm64
linux_arm64:
	@echo "Building Linux ARM64 version with Git Hash: $(GIT_HASH)"
	GOOS=linux GOARCH=arm64 $(WAILS) build -platform linux/arm64 -upx $(LDFLAGS) -o build/$(APP_NAME)-Linux-ARM64

# 打包目标：打包所有平台和架构的应用
.PHONY: package
package: package_mac package_windows package_linux

# macOS 打包
.PHONY: package_mac
package_mac: mac_amd64 mac_arm64 mac_universal
	@echo "Packaging macOS versions..."
	hdiutil create -volname "$(APP_NAME) AMD64" -srcfolder build/$(APP_NAME)-macOS-AMD64 -ov -format UDZO $(PKG_OUTPUT)-macOS-AMD64.dmg
	hdiutil create -volname "$(APP_NAME) ARM64" -srcfolder build/$(APP_NAME)-macOS-ARM64 -ov -format UDZO $(PKG_OUTPUT)-macOS-ARM64.dmg
	hdiutil create -volname "$(APP_NAME) Universal" -srcfolder build/$(APP_NAME)-macOS-Universal -ov -format UDZO $(PKG_OUTPUT)-macOS-Universal.dmg

# Windows 打包
.PHONY: package_windows
package_windows: windows_amd64 windows_arm64
	@echo "Packaging Windows versions..."
	zip -j $(PKG_OUTPUT)-Windows-AMD64.zip build/$(APP_NAME)-Windows-AMD64.exe
	zip -j $(PKG_OUTPUT)-Windows-ARM64.zip build/$(APP_NAME)-Windows-ARM64.exe

# Linux 打包
.PHONY: package_linux
package_linux: linux_amd64 linux_arm64
	@echo "Packaging Linux versions..."
	tar -czvf $(PKG_OUTPUT)-Linux-AMD64.tar.gz -C build $(APP_NAME)-Linux-AMD64
	tar -czvf $(PKG_OUTPUT)-Linux-ARM64.tar.gz -C build $(APP_NAME)-Linux-ARM64

# 帮助信息
.PHONY: help
help:
	@echo "Available targets:"
	@echo ""
	@echo "  主要目标:"
	@echo "    build            : 构建当前平台的原生版本 (macOS)"
	@echo "    build-native     : 构建 macOS 所有架构版本"
	@echo "    build-all        : 构建所有平台 (跨平台编译会失败，仅用于 CI)"
	@echo "    clean            : 清理构建目录"
	@echo "    dev              : 运行开发模式"
	@echo ""
	@echo "  macOS 构建:"
	@echo "    mac_amd64        : 构建 macOS AMD64 (Intel) 版本"
	@echo "    mac_arm64        : 构建 macOS ARM64 (Apple Silicon) 版本"
	@echo "    mac_universal    : 构建 macOS Universal 版本"
	@echo "    package_mac      : 打包 macOS 应用为 .dmg"
	@echo ""
	@echo "  Windows 构建 (需在 Windows 上或使用 GitHub Actions):"
	@echo "    windows_amd64    : 构建 Windows AMD64 版本"
	@echo "    windows_arm64    : 构建 Windows ARM64 版本"
	@echo "    package_windows  : 打包 Windows 应用"
	@echo ""
	@echo "  Linux 构建 (需在 Linux 上或使用 GitHub Actions):"
	@echo "    linux_amd64      : 构建 Linux AMD64 版本"
	@echo "    linux_arm64      : 构建 Linux ARM64 版本"
	@echo "    package_linux    : 打包 Linux 应用"
	@echo ""
	@echo "  其他:"
	@echo "    package          : 打包所有平台的应用"
	@echo "    help             : 显示此帮助信息"
	@echo ""
	@echo "注意: 跨平台编译有限制，详见 BUILD.md"