package util

import (
	"archive/zip"
	"fmt"
	"io"
	"jetbrains/global"
	"net/http"
	"os"
	"os/exec"
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

// ProgressCallback 下载进度回调函数类型
type ProgressCallback func(downloaded int64, total int64, percent float64)

// InstallProgressCallback 安装进度回调函数类型
type InstallProgressCallback func(percent float64)

func DownloadFile(fileUrl string, filePath string, progressCallback ProgressCallback) error {
	// 打印接收到的参数
	fmt.Printf("DownloadFile 接收到的参数 - fileUrl: %s, filePath: %s\n", fileUrl, filePath)

	// 检测filePath的目录是否存在，不存在则创建
	dir := filepath.Dir(filePath)
	fmt.Printf("提取的目录: %s\n", dir)
	if _, err := EnsureDir(dir, true); err != nil {
		return err
	}

	// 创建HTTP请求
	resp, err := http.Get(fileUrl)
	if err != nil {
		return fmt.Errorf("下载文件失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查HTTP响应状态
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败，HTTP状态码: %d", resp.StatusCode)
	}

	// 获取文件总大小
	totalSize := resp.ContentLength

	// 创建目标文件
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer out.Close()

	// 创建一个带进度的 Reader
	var downloaded int64
	buffer := make([]byte, 32*1024) // 32KB 缓冲区

	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			// 写入文件
			_, writeErr := out.Write(buffer[:n])
			if writeErr != nil {
				return fmt.Errorf("写入文件失败: %v", writeErr)
			}

			// 更新已下载大小
			downloaded += int64(n)

			// 计算进度百分比
			var percent float64
			if totalSize > 0 {
				percent = float64(downloaded) / float64(totalSize) * 100
			}

			// 调用进度回调
			if progressCallback != nil {
				progressCallback(downloaded, totalSize, percent)
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("读取数据失败: %v", err)
		}
	}

	return nil
}

// InstallWindows 在 Windows 上执行安装
func InstallWindows(filePath string, installDir string, progressCallback InstallProgressCallback) error {
	// 发送进度更新
	if progressCallback != nil {
		progressCallback(10.0)
	}

	// 使用 /S 静默安装，/D= 指定安装目录
	// 注意：/D= 参数必须是最后一个参数，且路径不能用引号
	cmd := exec.Command(filePath, "/S", fmt.Sprintf("/D=%s", installDir))

	// 执行安装命令
	fmt.Printf("执行安装命令: %s /S /D=%s\n", filePath, installDir)

	if progressCallback != nil {
		progressCallback(30.0)
	}

	// 运行命令并等待完成
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("安装失败: %v", err)
	}

	if progressCallback != nil {
		progressCallback(90.0)
	}

	// 安装完成后删除安装包
	fmt.Printf("正在删除安装包: %s\n", filePath)
	if err := os.Remove(filePath); err != nil {
		// 如果删除失败，打印警告但不返回错误
		fmt.Printf("警告：删除安装包失败: %v\n", err)
	} else {
		fmt.Printf("已删除安装包: %s\n", filePath)
	}

	if progressCallback != nil {
		progressCallback(100.0)
	}

	return nil
}

// InstallMacOS 在 macOS 上挂载 DMG 并复制应用到 /Applications
func InstallMacOS(filePath string, installDir string, progressCallback InstallProgressCallback) error {
	// 在 macOS 上，忽略 installDir 参数，始终安装到 /Applications
	applicationsDir := "/Applications"
	fmt.Printf("macOS 安装: %s 到 %s\n", filePath, applicationsDir)

	if progressCallback != nil {
		progressCallback(30.0)
	}

	// 1. 挂载 DMG 文件
	mountCmd := exec.Command("hdiutil", "attach", "-nobrowse", filePath)
	mountOutput, err := mountCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("挂载 DMG 失败: %v, 输出: %s", err, string(mountOutput))
	}

	// 从挂载输出中提取挂载点路径
	// 输出格式类似: /dev/disk2s1    Apple_HFS    /Volumes/AppName With Spaces
	lines := strings.Split(string(mountOutput), "\n")
	var mountPoint string
	for _, line := range lines {
		// 查找包含 /Volumes/ 的行
		if idx := strings.Index(line, "/Volumes/"); idx >= 0 {
			// 从 /Volumes/ 开始提取到行尾（支持路径中包含空格）
			mountPoint = strings.TrimSpace(line[idx:])
			break
		}
	}

	if mountPoint == "" {
		return fmt.Errorf("无法从挂载输出中提取挂载点路径")
	}

	fmt.Printf("挂载点: %s\n", mountPoint)

	// 确保卸载（使用 defer）
	defer func() {
		fmt.Printf("正在卸载: %s\n", mountPoint)
		detachCmd := exec.Command("hdiutil", "detach", mountPoint)
		if err := detachCmd.Run(); err != nil {
			fmt.Printf("卸载 DMG 失败: %v\n", err)
		}
	}()

	if progressCallback != nil {
		progressCallback(60.0)
	}

	// 2. 查找 .app 文件
	var appPath string
	entries, err := os.ReadDir(mountPoint)
	if err != nil {
		return fmt.Errorf("读取挂载点失败: %v", err)
	}

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".app" {
			appPath = filepath.Join(mountPoint, entry.Name())
			break
		}
	}

	if appPath == "" {
		return fmt.Errorf("在 DMG 中未找到 .app 文件")
	}

	fmt.Printf("找到应用: %s\n", appPath)

	if progressCallback != nil {
		progressCallback(80.0)
	}

	// 3. 复制 .app 到 /Applications 目录
	targetPath := filepath.Join(applicationsDir, filepath.Base(appPath))

	// 如果目标已存在，先删除
	if _, err := os.Stat(targetPath); err == nil {
		fmt.Printf("删除已存在的应用: %s\n", targetPath)
		if err := os.RemoveAll(targetPath); err != nil {
			return fmt.Errorf("删除已存在应用失败: %v", err)
		}
	}

	fmt.Printf("复制 %s 到 %s\n", appPath, targetPath)
	copyCmd := exec.Command("cp", "-R", appPath, targetPath)
	if output, err := copyCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("复制应用失败: %v, 输出: %s", err, string(output))
	}

	if progressCallback != nil {
		progressCallback(90.0)
	}

	// 安装完成后删除 DMG 文件
	fmt.Printf("正在删除安装包: %s\n", filePath)
	if err := os.Remove(filePath); err != nil {
		// 如果删除失败，打印警告但不返回错误
		fmt.Printf("警告：删除安装包失败: %v\n", err)
	} else {
		fmt.Printf("已删除安装包: %s\n", filePath)
	}

	if progressCallback != nil {
		progressCallback(100.0)
	}

	fmt.Printf("macOS 安装完成: %s\n", targetPath)
	return nil
}

// InstallLinux 在 Linux 上解压 tar.gz
func InstallLinux(filePath string, installDir string, progressCallback InstallProgressCallback) error {
	fmt.Printf("Linux 安装: %s 到 %s\n", filePath, installDir)

	if progressCallback != nil {
		progressCallback(30.0)
	}

	// 确保安装目录存在
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return fmt.Errorf("创建安装目录失败: %v", err)
	}

	if progressCallback != nil {
		progressCallback(50.0)
	}

	// 解压 tar.gz 文件到目标目录
	// 使用 tar -xzf 命令解压
	cmd := exec.Command("tar", "-xzf", filePath, "-C", installDir)

	// 执行解压命令
	fmt.Printf("执行解压命令: tar -xzf %s -C %s\n", filePath, installDir)

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("解压失败: %v, 输出: %s", err, string(output))
	}

	if progressCallback != nil {
		progressCallback(90.0)
	}

	// 安装完成后删除安装包
	fmt.Printf("正在删除安装包: %s\n", filePath)
	if err := os.Remove(filePath); err != nil {
		// 如果删除失败，打印警告但不返回错误
		fmt.Printf("警告：删除安装包失败: %v\n", err)
	} else {
		fmt.Printf("已删除安装包: %s\n", filePath)
	}

	if progressCallback != nil {
		progressCallback(100.0)
	}

	fmt.Printf("Linux 安装完成: %s\n", installDir)
	return nil
}
