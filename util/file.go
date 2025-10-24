package util

import (
	"archive/zip"
	"io"
	"jetbrains/global"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// InitGlobalPaths 初始化全局路径变量
func InitGlobalPaths() error {
	// 获取操作系统
	global.OS = runtime.GOOS

	// 获取用户主目录
	homeDir, err := getRealUserHome()
	if err != nil {
		return err
	}
	global.UserHome = homeDir
	getWorkDir()
	return nil
}

// getRealUserHome 获取真实用户的主目录（处理sudo情况）
func getRealUserHome() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// 在Linux/macOS上，如果通过sudo运行，获取原始用户的主目录
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		// 检查是否通过sudo运行
		if sudoUser := os.Getenv("SUDO_USER"); sudoUser != "" {
			// 确认当前是root用户
			if os.Getuid() == 0 {
				if runtime.GOOS == "darwin" {
					return filepath.Join("/Users", sudoUser), nil
				}
				// Linux
				return filepath.Join("/home", sudoUser), nil
			}
		}
	}

	return homeDir, nil
}

// GetAppDataDir 获取应用数据目录
func GetAppDataDir() string {
	var appDataDir string
	switch global.OS {
	case "windows":
		appDataDir = filepath.Join(global.UserHome, "AppData", "Roaming", "JetBrains")
	case "darwin":
		appDataDir = filepath.Join(global.UserHome, "Library", "Application Support", "JetBrains")
	case "linux":
		appDataDir = filepath.Join(global.UserHome, ".config", "JetBrains")
	default:
		appDataDir = filepath.Join(global.UserHome, ".Jetbrains")
	}
	return appDataDir
}

// GetConfigDir 获取配置目录
func GetConfigDir() string {
	var configDir string
	switch global.OS {
	case "windows":
		configDir = filepath.Join(global.UserHome, "AppData", "Local", "JetBrains")
	case "darwin":
		configDir = filepath.Join(global.UserHome, "Library", "Caches", "JetBrains")
	case "linux":
		configDir = filepath.Join(global.UserHome, ".cache", "JetBrains")
	default:
		configDir = filepath.Join(global.UserHome, ".Jetbrains")
	}
	return configDir
}

// getWorkDir 获取工作目录
func getWorkDir() {
	switch global.OS {
	case "windows":
		global.WorkDir = filepath.Join(global.UserHome, "../", "Public", ".jb_run")
	case "darwin", "linux":
		global.WorkDir = filepath.Join(global.UserHome, ".jb_run")
	}
	// 判断这个目录是否存在，如果不存在则创建
	if _, err := EnsureDir(global.WorkDir, true); err != nil {
		panic(err)
	}
}

// EnsureDir 确保目录存在，如果不存在并且 flag 为 true 则创建该目录
func EnsureDir(dir string, flag bool) (bool, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if flag {
			return false, os.MkdirAll(dir, 0755)
		} else {
			return false, nil
		}
	}
	return true, nil
}

// UnzipFile 解压 zip 文件到指定目录
func UnzipFile(zipPath, destDir string) error {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer reader.Close()

	// 确保目标目录存在
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	// 遍历 zip 文件中的所有文件
	for _, file := range reader.File {
		// 构建目标文件路径
		path := filepath.Join(destDir, file.Name)

		// 检查是否为 Zip Slip 漏洞
		if !strings.HasPrefix(path, filepath.Clean(destDir)+string(os.PathSeparator)) {
			return os.ErrInvalid
		}

		if file.FileInfo().IsDir() {
			// 如果是目录，创建目录
			os.MkdirAll(path, file.Mode())
			continue
		}

		// 创建文件所在的目录
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}

		// 创建目标文件
		destFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		// 打开 zip 中的文件
		srcFile, err := file.Open()
		if err != nil {
			destFile.Close()
			return err
		}

		// 复制文件内容
		_, err = io.Copy(destFile, srcFile)
		srcFile.Close()
		destFile.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

// ClearWorkDir 清空工作目录中的所有文件和子目录
func ClearWorkDir() error {
	entries, err := os.ReadDir(global.WorkDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryPath := filepath.Join(global.WorkDir, entry.Name())
		err := os.RemoveAll(entryPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func CopyTextFromWorkDir() string {
	textFilePath := filepath.Join(global.WorkDir, "激活码.txt")
	data, err := os.ReadFile(textFilePath)
	if err != nil {
		return ""
	}
	return string(data)
}
