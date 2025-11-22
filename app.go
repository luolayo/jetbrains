package main

import (
	"context"
	"fmt"
	"jetbrains/global"
	"jetbrains/util"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// 初始化全局路径变量
	if err := util.InitGlobalPaths(); err != nil {
		fmt.Printf("初始化全局路径失败: %v\n", err)
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetDeviceID 获取设备唯一标识符
func (a *App) GetDeviceID() util.DeviceInfo {
	deviceUtil := &util.App{}
	return deviceUtil.GetDeviceID()
}

// Actions 自动激活
func (a *App) Actions() util.ActionsType {
	err := util.RemoveEnvOther()
	if err != nil {
		return util.ActionsType{
			Error: []string{err.Error()},
		}
	}
	err = util.ReadAllJetbrainsProducts()
	if err != nil {
		return util.ActionsType{
			Error: []string{err.Error()},
		}
	}
	var actions util.ActionsType
	return actions.Actions()
}

// Clean 激活清理
func (a *App) Clean() string {
	switch global.OS {
	case "darwin", "linux":
		_ = util.RemoveEnvOther()
	case "windows":
		_ = util.RemoveEnvOtherWindows()

	default:
	}

	_ = util.ReadAllJetbrainsProducts()
	return "清理完成"
}

// DownloadAndExtract 下载并解压文件到 WorkDir
// fileData: 文件的字节数据
// filename: 文件名（例如: "jetbrains-activation.zip"）
func (a *App) DownloadAndExtract(fileData []byte, filename string) error {
	// 清空workdir 里面的文件
	err := util.ClearWorkDir()
	if err != nil {
		return fmt.Errorf("清空工作目录失败: %v", err)
	}
	// 构建保存路径
	zipPath := filepath.Join(global.WorkDir, filename)

	// 保存文件到 WorkDir
	if err := os.WriteFile(zipPath, fileData, 0644); err != nil {
		return fmt.Errorf("保存文件失败: %v", err)
	}

	// 解压文件到 WorkDir
	if err := util.UnzipFile(zipPath, global.WorkDir); err != nil {
		return fmt.Errorf("解压文件失败: %v", err)
	}

	// 删除 zip 文件
	if err := os.Remove(zipPath); err != nil {
		fmt.Printf("删除 zip 文件失败: %v\n", err)
	}

	return nil
}

// CopyText 从工作目录的激活码.txt 复制文本内容
func (a *App) CopyText() string {
	return util.CopyTextFromWorkDir()
}

// SelectDirectory 打开目录选择对话框
func (a *App) SelectDirectory() (string, error) {
	return runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择保存目录",
	})
}

// SelectFile 打开文件选择对话框
func (a *App) SelectFile() (string, error) {
	return runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择安装文件",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Windows 安装文件 (*.exe)",
				Pattern:     "*.exe",
			},
			{
				DisplayName: "macOS 安装文件 (*.dmg)",
				Pattern:     "*.dmg",
			},
			{
				DisplayName: "Linux 安装文件 (*.tar.gz)",
				Pattern:     "*.tar.gz",
			},
			{
				DisplayName: "所有文件 (*.*)",
				Pattern:     "*.*",
			},
		},
	})
}

// DownloadFile 从给定的URL下载文件到指定路径，并实时返回下载进度
// fileUrl: 文件下载链接
// filePath: 保存文件的完整路径
func (a *App) DownloadFile(fileUrl string, filePath string) error {
	// 创建进度回调函数，通过事件将进度发送到前端
	progressCallback := func(downloaded int64, total int64, percent float64) {
		// 发送下载进度事件到前端
		runtime.EventsEmit(a.ctx, "download:progress", map[string]interface{}{
			"filePath":   filePath,
			"downloaded": downloaded,
			"total":      total,
			"percent":    percent,
		})
	}

	// 调用下载函数并传入进度回调
	err := util.DownloadFile(fileUrl, filePath, progressCallback)

	// 下载完成或失败后发送完成事件
	if err != nil {
		runtime.EventsEmit(a.ctx, "download:error", map[string]interface{}{
			"filePath": filePath,
			"error":    err.Error(),
		})
		return err
	}

	runtime.EventsEmit(a.ctx, "download:complete", map[string]interface{}{
		"filePath": filePath,
	})

	return nil
}

// DownloadAndInstall 下载文件并自动安装
// fileUrl: 文件下载链接
// downloadPath: 下载保存路径（临时目录）
// installDir: 安装目标目录
func (a *App) DownloadAndInstall(fileUrl string, downloadPath string, installDir string) error {

	// 先下载文件
	err := a.DownloadFile(fileUrl, downloadPath)
	if err != nil {
		return err
	}

	// 下载完成后自动安装
	return a.InstallFile(downloadPath, installDir)
}

// InstallFile 安装下载的文件到指定目录
// filePath: 下载的文件路径
// installDir: 安装目标目录
func (a *App) InstallFile(filePath string, installDir string) error {
	// 发送安装开始事件
	runtime.EventsEmit(a.ctx, "install:start", map[string]interface{}{
		"filePath": filePath,
	})

	// 创建进度回调函数
	progressCallback := func(percent float64) {
		runtime.EventsEmit(a.ctx, "install:progress", map[string]interface{}{
			"filePath": filePath,
			"percent":  percent,
		})
	}

	// 根据文件扩展名和操作系统判断安装方式
	var err error

	if filepath.Ext(filePath) == ".exe" {
		// Windows 执行安装
		err = util.InstallWindows(filePath, installDir, progressCallback)
	} else if filepath.Ext(filePath) == ".dmg" {
		// macOS 挂载并安装
		err = util.InstallMacOS(filePath, installDir, progressCallback)
	} else if filepath.Ext(filePath) == ".gz" {
		// Linux 解压安装
		err = util.InstallLinux(filePath, installDir, progressCallback)
	} else {
		err = fmt.Errorf("不支持的文件格式: %s", filepath.Ext(filePath))
	}

	if err != nil {
		runtime.EventsEmit(a.ctx, "install:error", map[string]interface{}{
			"filePath": filePath,
			"error":    err.Error(),
		})
		return err
	}

	// 发送安装完成事件
	runtime.EventsEmit(a.ctx, "install:complete", map[string]interface{}{
		"filePath": filePath,
	})

	return nil
}

// CheckPermissions 检查并请求必要的权限（macOS）
func (a *App) CheckPermissions() bool {
	_, err := os.ReadDir("~/Library/Containers")
	// 如果没有权限就请求权限
	if err != nil {
	}
	return err == nil
}
