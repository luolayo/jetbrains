//go:build windows

package util

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

func getBIOSSerialNumber() (string, error) {
	// 优先获取 UUID，如果失败则尝试主板序列号
	cmd := exec.Command("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe", "-Command",
		"$cs = Get-CimInstance Win32_ComputerSystemProduct; "+
			"if ($cs.UUID -and $cs.UUID -ne '00000000-0000-0000-0000-000000000000') { $cs.UUID } "+
			"else { (Get-CimInstance Win32_BaseBoard | Select-Object -First 1).SerialNumber }")
	// 隐藏 PowerShell 窗口
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000, // CREATE_NO_WINDOW
	}
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("执行 PowerShell 命令失败: %v", err)
	}
	result := strings.TrimSpace(string(output))

	// 清理可能的多余空白和换行
	result = strings.ReplaceAll(result, "\r", "")
	result = strings.ReplaceAll(result, "\n", "")

	return result, nil
}

func getCPUSerialNumber() (string, error) {
	cmd := exec.Command("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe", "-Command",
		"Get-CimInstance Win32_Processor | Select-Object -First 1 -ExpandProperty ProcessorId")
	// 隐藏 PowerShell 窗口
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000, // CREATE_NO_WINDOW
	}
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("执行 PowerShell 命令失败: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}

func getDiskSerialNumber() (string, error) {
	cmd := exec.Command("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe", "-Command",
		"Get-CimInstance Win32_DiskDrive | Select-Object -First 1 -ExpandProperty SerialNumber")
	// 隐藏 PowerShell 窗口
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000, // CREATE_NO_WINDOW
	}
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("执行 PowerShell 命令失败: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}

func getBoardSerialNumber() (string, error) {
	cmd := exec.Command("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe", "-Command",
		"Get-CimInstance Win32_BaseBoard | Select-Object -First 1 -ExpandProperty SerialNumber")
	// 隐藏 PowerShell 窗口
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000, // CREATE_NO_WINDOW
	}
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("执行 PowerShell 命令失败: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}
