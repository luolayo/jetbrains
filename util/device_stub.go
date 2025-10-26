//go:build !windows

package util

import "fmt"

// 这些函数在非 Windows 平台上不会被调用
// 仅作为编译占位符存在

func getBIOSSerialNumber() (string, error) {
	return "", fmt.Errorf("此功能仅支持 Windows 平台")
}

func getCPUSerialNumber() (string, error) {
	return "", fmt.Errorf("此功能仅支持 Windows 平台")
}

func getDiskSerialNumber() (string, error) {
	return "", fmt.Errorf("此功能仅支持 Windows 平台")
}

func getBoardSerialNumber() (string, error) {
	return "", fmt.Errorf("此功能仅支持 Windows 平台")
}
