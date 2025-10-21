package util

import (
	"context"
	"fmt"
	"jetbrains/global"
	"os/exec"
	"strings"

	machinecode "github.com/super-l/machine-code"
)

type App struct {
	ctx context.Context
}

// DeviceInfo 设备信息结构
type DeviceInfo struct {
	MachineCode interface{} `json:"machineCode"`
	Error       string      `json:"error,omitempty"`
}

// GetDeviceID 获取设备唯一标识符
func (a *App) GetDeviceID() DeviceInfo {
	if len(machinecode.MachineErr) > 0 {
		if global.OS == "windows" {
			var errs string
			uuid, err := getBIOSSerialNumber()
			if err != nil {
				errs += "获取BIOS序列号失败:" + err.Error()
			}
			boardSN, err := getBoardSerialNumber()
			if err != nil {
				errs += "获取主板序列号失败:" + err.Error()
			}
			cpuSN, err := getCPUSerialNumber()
			if err != nil {
				errs += "获取CPU序列号失败:" + err.Error()
			}
			diskSN, err := getDiskSerialNumber()
			if err != nil {
				errs += "获取硬盘序列号失败:" + err.Error()
			}
			if len(errs) > 0 {
				return DeviceInfo{
					Error: errs,
				}
			}
			machinecode.Machine.UUID = uuid
			machinecode.Machine.BoardSerialNumber = boardSN
			machinecode.Machine.CpuSerialNumber = cpuSN
			machinecode.Machine.DiskSerialNumber = diskSN
			return DeviceInfo{
				MachineCode: machinecode.Machine,
			}
		}
		fmt.Println("获取机器码信息错误:")
		errorMsg := ""
		for _, err := range machinecode.MachineErr {
			fmt.Println(err.Error())
			errorMsg += err.Error() + "; "
		}
		return DeviceInfo{Error: errorMsg}
	}

	return DeviceInfo{
		MachineCode: machinecode.Machine,
	}
}

func getBIOSSerialNumber() (string, error) {
	cmd := exec.Command("powershell", "-Command",
		"Get-WmiObject -Class Win32_Bios | Select-Object -ExpandProperty SerialNumber")

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("执行 PowerShell 命令失败: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}

func getCPUSerialNumber() (string, error) {
	cmd := exec.Command("powershell", "-Command",
		"Get-CimInstance Win32_Processor | Select-Object -First 1 -ExpandProperty ProcessorId")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("执行 PowerShell 命令失败: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}

func getDiskSerialNumber() (string, error) {
	cmd := exec.Command("powershell", "-Command",
		"Get-CimInstance Win32_DiskDrive | Select-Object -First 1 -ExpandProperty SerialNumber")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("执行 PowerShell 命令失败: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}

func getBoardSerialNumber() (string, error) {
	cmd := exec.Command("powershell", "-Command",
		"Get-CimInstance Win32_BaseBoard | Select-Object -First 1 -ExpandProperty SerialNumber")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("执行 PowerShell 命令失败: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}
