package util

import (
	"fmt"
	"jetbrains/global"
	"os"
	"path/filepath"
)

type Product struct {
	ProductName    string `json:"productName"`
	ProductVersion string `json:"productVersion"`
}

type ActionsType struct {
	Product []Product `json:"product"`
	Error   []string  `json:"error,omitempty"`
}

func (a *ActionsType) Actions() ActionsType {
	flag, _ := EnsureDir(GetAppDataDir(), false)
	if !flag {
		a.Error = append(a.Error, "未检测到 Jetbrains 应用数据目录，请确保您需要激活的软件已安装，并且已运行过一次！")
	}
	// 只有 macOS 和 Linux 需要移除其他环境变量
	switch global.OS {
	case "darwin", "linux":
		_ = RemoveEnvOther()
	case "windows":
		_ = RemoveEnvOtherWindows()
	default:
	}
	err := ReadAllJetbrainsProducts()
	if err != nil {
		a.Error = append(a.Error, err.Error())
	}
	configDir := GetConfigDir()
	entries, err := os.ReadDir(configDir)
	// macOS 清空之前可能残留的管理员任务
	if global.OS == "darwin" {
		ClearAdminTasks()
	}
	if err != nil {
		a.Error = append(a.Error, err.Error())
	}
	for _, entry := range entries {
		if entry.IsDir() {
			productDir := filepath.Join(configDir, entry.Name())
			product, err := ActivateSingleJetbrainsProduct(productDir)
			if err != nil {
				a.Error = append(a.Error, fmt.Sprintf("激活 %s 失败，错误信息：%s", entry.Name(), err.Error()))
			} else {
				a.Product = append(a.Product, product)
			}
		}
	}
	// macOS 一次性执行所有需要管理员权限的任务
	if global.OS == "darwin" {
		if err := ExecuteAllAdminTasks(); err != nil {
			a.Error = append(a.Error, fmt.Sprintf("执行管理员权限任务失败：%s", err.Error()))
		}
	}
	return *a
}
