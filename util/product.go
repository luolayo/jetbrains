package util

import (
	"fmt"
	"jetbrains/global"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// HandleJetbrainsDir 处理单个 Jetbrains 产品目录
func HandleJetbrainsDir(productDir string) error {
	// 进入传过来的目录去找.home文件，如果存在则读取内容
	homeFilePath := filepath.Join(productDir, ".home")
	if _, err := os.Stat(homeFilePath); os.IsNotExist(err) {
		// .home 文件不存在，跳过
		return nil
	}
	// 读取 .home 文件内容
	data, err := os.ReadFile(homeFilePath)
	if err != nil {
		return err
	}
	// installPath 加上bin再去调用HandleJetbrainsProducts
	installPath := string(data)
	binPath := filepath.Join(installPath, "bin")

	// 从productDir获取产品名称和版本号
	_, productDirName := filepath.Split(productDir)
	var productName, productVersion string
	for i, r := range productDirName {
		if r >= '0' && r <= '9' {
			productName = productDirName[:i]
			productVersion = productDirName[i:]
			break
		}
	}
	fmt.Printf("产品名称: %s, 版本号: %s\n", productName, productVersion)
	err = HandleJetbrainsProducts(binPath, productName, productDirName)
	if err != nil {
		return err
	}
	return nil
}

// HandleJetbrainsProducts 进一步处理所有 Jetbrains 产品
func HandleJetbrainsProducts(productDir string, productName string, productDirName string) error {
	var vmoptions string
	switch global.OS {
	case "windows":
		vmoptions = "64.exe.vmoptions"
	case "darwin":
		vmoptions = ".vmoptions"
	case "linux":
		vmoptions = "64.vmoptions"
	default:
		return fmt.Errorf("不支持的操作系统: %s", global.OS)
	}

	// 构建 vmoptions 文件路径
	productCode := ParseProductDirName(productName)
	vmoptionsPath := filepath.Join(productDir, productCode+vmoptions)
	err := RemoveJetbrainsProductVmoptions(vmoptionsPath)
	if err != nil {
		return err
	}
	vmoptionsPath2 := filepath.Join(GetAppDataDir(), productDirName, productCode+vmoptions)
	err = RemoveJetbrainsProductVmoptions(vmoptionsPath2)
	if err != nil {
		return err
	}
	// Windows 不需要处理第二个 vmoptions 文件

	otherVmoptionsPattern := regexp.MustCompile(`^jetbrains_.*\.vmoptions$`)
	entries, err := os.ReadDir(productDir)
	if err != nil {
		return err
	}
	// 如果有就删掉
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if otherVmoptionsPattern.MatchString(entry.Name()) {
			otherVmoptionsPath := filepath.Join(productDir, entry.Name())
			fmt.Printf("准备删除文件: %s\n", otherVmoptionsPath)
			if err := os.Remove(otherVmoptionsPath); err != nil {
				fmt.Printf("删除文件 %s 时出错: %v\n", otherVmoptionsPath, err)
			} else {
				fmt.Printf("已删除: %s\n", otherVmoptionsPath)
			}
		}
	}
	return nil
}

// RemoveJetbrainsProductVmoptions 清理 Jetbrains 产品的 vmoptions 文件
func RemoveJetbrainsProductVmoptions(vmoptionsPath string) error {
	// 检测 vmoptions 文件是否存在
	if _, err := os.Stat(vmoptionsPath); os.IsNotExist(err) {
		return fmt.Errorf("vmoptions 文件不存在: %s", vmoptionsPath)
	}

	// 读取 vmoptions 内容并打印调试信息
	data, err := os.ReadFile(vmoptionsPath)
	if err != nil {
		return err
	}

	// 使用正则表达式匹配 -javaagent: *.jar 和 --add-opens=java.base/jdk.internal.org.objectweb.asm(.tree)?=ALL-UNNAMED
	reJavaAgent := regexp.MustCompile(`(?i)-javaagent:.*\.jar`)
	reAddOpens := regexp.MustCompile(`^--add-opens=java\.base/jdk\.internal\.org\.objectweb\.asm(?:\.tree)?=ALL-UNNAMED`)

	lines := strings.Split(string(data), "\n")
	var out []string
	changed := false
	for _, line := range lines {
		trim := strings.TrimSpace(line)
		if trim == "" {
			out = append(out, line)
			continue
		}
		if reJavaAgent.MatchString(trim) || reAddOpens.MatchString(trim) {
			changed = true
			fmt.Printf("移除匹配行: %s\n", line)
			continue
		}
		out = append(out, line)
	}

	if !changed {
		fmt.Println("vmoptions 未修改:", vmoptionsPath)
		return nil
	}

	newContent := strings.Join(out, "\n")
	if err := os.WriteFile(vmoptionsPath, []byte(newContent), 0644); err != nil {
		return err
	}
	fmt.Printf("已清理 vmoptions: %s （原始字节=%d， 新字节=%d）\n", vmoptionsPath, len(data), len(newContent))
	return nil
}

// ReadAllJetbrainsProducts 从配置目录下读取所有 Jetbrains 产品目录
func ReadAllJetbrainsProducts() error {
	configDir := GetConfigDir()
	entries, err := os.ReadDir(configDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			productDir := filepath.Join(configDir, entry.Name())
			if err := HandleJetbrainsDir(productDir); err != nil {
				fmt.Printf("处理目录 %s 时出错: %v\n", productDir, err)
			}
		}
	}
	return nil
}

// ParseProductDirName 产品解析函数
func ParseProductDirName(productName string) string {
	// 将产品名称解析成产品code
	switch {
	case strings.HasPrefix(productName, "IntelliJIdea"):
		return "idea"
	case strings.HasPrefix(productName, "PyCharm"):
		return "pycharm"
	case strings.HasPrefix(productName, "WebStorm"):
		return "webstorm"
	case strings.HasPrefix(productName, "PhpStorm"):
		return "phpstorm"
	case strings.HasPrefix(productName, "CLion"):
		return "clion"
	case strings.HasPrefix(productName, "GoLand"):
		return "goland"
	case strings.HasPrefix(productName, "Rider"):
		return "rider"
	case strings.HasPrefix(productName, "DataGrip"):
		return "datagrip"
	case strings.HasPrefix(productName, "RubyMine"):
		return "rubymine"
	case strings.HasPrefix(productName, "AppCode"):
		return "appcode"
	default:
		return ""
	}
}

// RemoveEnvOther 清理三方工具设置的环境变量
func RemoveEnvOther() error {
	osName := global.OS
	jbProducts := []string{
		"idea", "clion", "phpstorm", "goland", "pycharm",
		"webstorm", "webide", "rider", "datagrip", "rubymine",
		"appcode", "dataspell", "gateway", "jetbrains_client", "jetbrainsclient",
	}

	kdeEnvDir := filepath.Join(global.UserHome, ".config", "plasma-workspace", "env")
	profilePath := filepath.Join(global.UserHome, ".profile")
	zshProfilePath := filepath.Join(global.UserHome, ".zshrc")
	plistPath := filepath.Join(global.UserHome, "Library", "LaunchAgents", "jetbrains.vmoptions.plist")

	var bashProfilePath string
	if osName == "darwin" {
		bashProfilePath = filepath.Join(global.UserHome, ".bash_profile")
	} else {
		bashProfilePath = filepath.Join(global.UserHome, ".bashrc")
	}

	// touch 三个文件（存在则无操作，不存在则创建）
	for _, p := range []string{profilePath, bashProfilePath, zshProfilePath} {
		f, err := os.OpenFile(p, os.O_CREATE, 0644)
		if err != nil {
			// 只记录错误并继续，避免因单个不可写文件导致整体失败
			fmt.Printf("warning: 无法 touch %s: %v\n", p, err)
			continue
		}
		_ = f.Close()
	}

	// 删除用户级别的 jetbrains vmoptions shell 文件
	myVmShellName := "jetbrains.vmoptions.sh"
	myVmShellFile := filepath.Join(global.UserHome, "."+myVmShellName)
	if err := os.RemoveAll(myVmShellFile); err != nil {
		fmt.Printf("warning: 删除 %s 失败: %v\n", myVmShellFile, err)
	}

	// 要从 shell 配置中移除的标记（用简单关键字匹配）
	removeMarker := "___MY_VMOPTIONS_SHELL_FILE"

	if osName == "darwin" {
		// 在 macOS 上通过 launchctl 清理环境变量
		for _, prd := range jbProducts {
			envName := strings.ToUpper(prd) + "_VM_OPTIONS"
			_ = exec.Command("launchctl", "unsetenv", envName).Run()
		}
		// 删除 plist
		if err := os.RemoveAll(plistPath); err != nil {
			fmt.Printf("warning: 删除 %s 失败: %v\n", plistPath, err)
		}
		// 从 profile 等文件中删除包含标记的行
		_ = removeLinesContaining(profilePath, removeMarker)
		_ = removeLinesContaining(bashProfilePath, removeMarker)
		_ = removeLinesContaining(zshProfilePath, removeMarker)
	} else {
		// 非 macOS：删除配置中相同的行
		_ = removeLinesContaining(profilePath, removeMarker)
		_ = removeLinesContaining(bashProfilePath, removeMarker)
		_ = removeLinesContaining(zshProfilePath, removeMarker)

		// 删除 KDE env 下的文件
		kdeTarget := filepath.Join(kdeEnvDir, myVmShellName)
		if err := os.RemoveAll(kdeTarget); err != nil {
			fmt.Printf("warning: 删除 %s 失败: %v\n", kdeTarget, err)
		}
	}

	fmt.Println("清理三方工具环境变量完成")
	return nil
}

// RemoveEnvOtherWindows Windows 移除所有以 _VM_OPTIONS 结尾的环境变量
func RemoveEnvOtherWindows() error {
	// 获取所有环境变量
	envVars := os.Environ()
	var vmOptionsVars []string

	for _, env := range envVars {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) != 2 {
			continue
		}
		envName := parts[0]
		if strings.HasSuffix(envName, "_VM_OPTIONS") {
			vmOptionsVars = append(vmOptionsVars, envName)
		}
	}

	if len(vmOptionsVars) == 0 {
		fmt.Println("未找到任何以 _VM_OPTIONS 结尾的环境变量")
		return nil
	}

	fmt.Printf("找到 %d 个 _VM_OPTIONS 环境变量，准备删除...\n", len(vmOptionsVars))

	// 删除用户环境变量
	for _, envName := range vmOptionsVars {
		cmd := exec.Command("reg", "delete", "HKCU\\Environment", "/v", envName, "/f")
		if err := cmd.Run(); err != nil {
			fmt.Printf("warning: 删除用户环境变量 %s 失败: %v\n", envName, err)
		} else {
			fmt.Printf("已删除用户环境变量: %s\n", envName)
		}
	}

	// 删除系统环境变量
	for _, envName := range vmOptionsVars {
		cmd := exec.Command("reg", "delete", "HKLM\\SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment", "/v", envName, "/f")
		if err := cmd.Run(); err != nil {
			fmt.Printf("warning: 删除系统环境变量 %s 失败（需要管理员权限）: %v\n", envName, err)
		} else {
			fmt.Printf("已删除系统环境变量: %s\n", envName)
		}
	}
	fmt.Println("清理 Windows 环境变量完成")
	return nil
}

// removeLinesContaining 读取 `filePath`，删除包含 `substr` 的行并写回，若无变化则不修改文件。
func removeLinesContaining(filePath, substr string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		// 如果文件不存在，视为无须处理
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	lines := strings.Split(string(data), "\n")
	out := make([]string, 0, len(lines))
	changed := false
	for _, line := range lines {
		if strings.Contains(line, substr) {
			changed = true
			continue
		}
		out = append(out, line)
	}
	if !changed {
		return nil
	}
	newData := strings.Join(out, "\n")
	// 写回文件（保留原模式为 0644）
	if err := os.WriteFile(filePath, []byte(newData), 0644); err != nil {
		return err
	}
	return nil
}

// ActivateSingleJetbrainsProduct 激活单个 Jetbrains 产品
func ActivateSingleJetbrainsProduct(productDir string) (Product, error) {
	// 进入传过来的目录去找.home文件，如果存在则读取内容
	homeFilePath := filepath.Join(productDir, ".home")
	if _, err := os.Stat(homeFilePath); os.IsNotExist(err) {
		// .home 文件不存在，跳过
		return Product{}, nil
	}
	// 读取 .home 文件内容
	data, err := os.ReadFile(homeFilePath)
	if err != nil {
		return Product{}, nil
	}
	// installPath 加上bin再去调用HandleJetbrainsProducts
	installPath := string(data)
	binPath := filepath.Join(installPath, "bin")

	// 从productDir获取产品名称和版本号
	_, productDirName := filepath.Split(productDir)
	// 检测产品名称有没有空格有就去掉
	productDirName = strings.ReplaceAll(productDirName, " ", "")
	var productName, productVersion string
	for i, r := range productDirName {
		if r >= '0' && r <= '9' {
			productName = productDirName[:i]
			productVersion = productDirName[i:]
			break
		}
	}
	err = ActivateJetbrainsProduct(binPath, productName, productDirName)
	if err != nil {
		return Product{}, err
	}
	var product Product
	product.ProductName = productName
	product.ProductVersion = productVersion
	return product, nil
}

// ActivateJetbrainsProduct 进一步激活单个产品
func ActivateJetbrainsProduct(productDir string, productName string, productDirName string) error {
	var vmoptions string
	switch global.OS {
	case "windows":
		vmoptions = "64.exe.vmoptions"
	case "darwin":
		vmoptions = ".vmoptions"
	case "linux":
		vmoptions = "64.vmoptions"
	default:
		return fmt.Errorf("不支持的操作系统: %s", global.OS)
	}

	// 构建 vmoptions 文件路径
	productCode := ParseProductDirName(productName)
	vmoptionsPath := filepath.Join(productDir, productCode+vmoptions)
	err := AppendVmoptionsForActivation(vmoptionsPath)
	if err != nil {
		return err
	}
	vmoptionsPath2 := filepath.Join(GetAppDataDir(), productDirName, productCode+vmoptions)
	err = AppendVmoptionsForActivation(vmoptionsPath2)
	if err != nil {
		return err
	}
	return nil
}

// AppendVmoptionsForActivation 给vmoptions文件追加激活参数
func AppendVmoptionsForActivation(vmoptionsPath string) error {
	// 检测 vmoptions 文件是否存在
	if _, err := os.Stat(vmoptionsPath); os.IsNotExist(err) {
		fmt.Println("vmoptions 文件不存在: %s", vmoptionsPath)
		return nil
	}
	// 拼接激活参数
	activationLines := []string{
		"--add-opens=java.base/jdk.internal.org.objectweb.asm=ALL-UNNAMED",
		"--add-opens=java.base/jdk.internal.org.objectweb.asm.tree=ALL-UNNAMED",
		"-javaagent:" + global.WorkDir + "/ja-netfilter.jar=jetbrains",
	}
	// 在文件最后面追加激活参数
	f, err := os.OpenFile(vmoptionsPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("文件打开失败: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("关闭文件失败: %v\n", err)
		}
	}(f)

	for _, line := range activationLines {
		if _, err := f.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("写入文件失败: %v", err)
		}
	}
	return nil
}
