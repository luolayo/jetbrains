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

type ManualActionsResult struct {
	Product      Product `json:"product"`
	NeedDownload bool    `json:"needDownload"`
	Error        string  `json:"error,omitempty"`
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
		// Windows 不需要移除其他环境变量
	}
	err := ReadAllJetbrainsProducts()
	if err != nil {
		a.Error = append(a.Error, err.Error())
	}
	configDir := GetConfigDir()
	entries, err := os.ReadDir(configDir)
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
				// 对于成功激活的产品，添加到 Product 列表中，如果product为空则不添加

				if product.ProductName != "" && product.ProductVersion != "" {
					a.Product = append(a.Product, product)
				}
			}
		}
	}
	return *a
}
