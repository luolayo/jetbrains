package main

import (
	"context"
	"fmt"
	"jetbrains/global"
	"jetbrains/util"
	"os"
	"path/filepath"
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
