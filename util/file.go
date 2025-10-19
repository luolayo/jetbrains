package util

import (
	"jetbrains/global"
	"os"
	"path/filepath"
	"runtime"
)

// InitGlobalPaths 初始化全局路径变量
func InitGlobalPaths() error {
	// 获取操作系统
	global.OS = runtime.GOOS

	// 获取用户主目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	global.UserHome = homeDir
	getWorkDir()
	return nil
}

// GetAppDataDir 获取应用数据目录
func GetAppDataDir() string {
	var appDataDir string
	switch global.OS {
	case "windows":
		appDataDir = filepath.Join(global.UserHome, "AppData", "Local", "Jetbrains")
	case "darwin":
		appDataDir = filepath.Join(global.UserHome, "Library", "Application Support", "Jetbrains")
	case "linux":
		appDataDir = filepath.Join(global.UserHome, ".config", "share", "Jetbrains")
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
		configDir = filepath.Join(global.UserHome, "AppData", "Roaming", "Jetbrains")
	case "darwin":
		configDir = filepath.Join(global.UserHome, "Library", "Caches", "Jetbrains")
	case "linux":
		configDir = filepath.Join(global.UserHome, ".cache", "Jetbrains")
	default:
		configDir = filepath.Join(global.UserHome, ".Jetbrains")
	}
	return configDir
}

// getWorkDir 获取工作目录
func getWorkDir() {
	switch global.OS {
	case "windows":
		global.WorkDir = filepath.Join(global.UserHome, "Public", ".jb_run")
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
