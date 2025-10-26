package util

import (
	"context"
	"fmt"
	"jetbrains/global"
	"os"
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
			// 打印设备信息
			fmt.Println("========== 设备信息 ==========")
			fmt.Printf("BIOS序列号: %s\n", uuid)
			fmt.Printf("主板序列号: %s\n", boardSN)
			fmt.Printf("CPU序列号: %s\n", cpuSN)
			fmt.Printf("硬盘序列号: %s\n", diskSN)
			fmt.Println("=============================")
			return DeviceInfo{
				MachineCode: machinecode.Machine,
			}
		}

		// Linux系统：使用machine-id作为后备方案（不需要sudo权限）
		if global.OS == "linux" {
			machineID, err := GetLinuxMachineID()
			if err != nil {
				fmt.Println("获取机器码信息错误:")
				errorMsg := ""
				for _, e := range machinecode.MachineErr {
					fmt.Println(e.Error())
					errorMsg += e.Error() + "; "
				}
				errorMsg += "无法获取machine-id: " + err.Error()
				return DeviceInfo{Error: errorMsg}
			}

			// 使用machine-id作为UUID
			machinecode.Machine.UUID = machineID
			fmt.Println("========== 设备信息 ==========")
			fmt.Printf("Machine ID: %s\n", machineID)
			fmt.Println("注意: 使用/etc/machine-id作为设备标识（无需sudo权限）")
			fmt.Println("=============================")
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

//func getBIOSSerialNumber() (string, error) {
//	// 优先获取 UUID，如果失败则尝试主板序列号
//	cmd := exec.Command("powershell", "-WindowStyle", "Hidden", "-Command",
//		"$cs = Get-CimInstance Win32_ComputerSystemProduct; "+
//			"if ($cs.UUID -and $cs.UUID -ne '00000000-0000-0000-0000-000000000000') { $cs.UUID } "+
//			"else { (Get-CimInstance Win32_BaseBoard | Select-Object -First 1).SerialNumber }")
//	output, err := cmd.Output()
//	if err != nil {
//		return "", fmt.Errorf("执行 PowerShell 命令失败: %v", err)
//	}
//	result := strings.TrimSpace(string(output))
//
//	// 清理可能的多余空白和换行
//	result = strings.ReplaceAll(result, "\r", "")
//	result = strings.ReplaceAll(result, "\n", "")
//
//	return result, nil
//}

//func getCPUSerialNumber() (string, error) {
//	cmd := exec.Command("powershell", "-WindowStyle", "Hidden", "-Command",
//		"Get-CimInstance Win32_Processor | Select-Object -First 1 -ExpandProperty ProcessorId")
//	output, err := cmd.Output()
//	if err != nil {
//		return "", fmt.Errorf("执行 PowerShell 命令失败: %v", err)
//	}
//	return strings.TrimSpace(string(output)), nil
//}

//func getDiskSerialNumber() (string, error) {
//	cmd := exec.Command("powershell", "-WindowStyle", "Hidden", "-Command",
//		"Get-CimInstance Win32_DiskDrive | Select-Object -First 1 -ExpandProperty SerialNumber")
//	output, err := cmd.Output()
//	if err != nil {
//		return "", fmt.Errorf("执行 PowerShell 命令失败: %v", err)
//	}
//	return strings.TrimSpace(string(output)), nil
//}

//func getBoardSerialNumber() (string, error) {
//	cmd := exec.Command("powershell", "-WindowStyle", "Hidden", "-Command",
//		"Get-CimInstance Win32_BaseBoard | Select-Object -First 1 -ExpandProperty SerialNumber")
//	output, err := cmd.Output()
//	if err != nil {
//		return "", fmt.Errorf("执行 PowerShell 命令失败: %v", err)
//	}
//	return strings.TrimSpace(string(output)), nil
//}

// GetLinuxMachineID 获取Linux机器ID（不需要root权限）
// 优先读取 /etc/machine-id，如果失败则尝试 /var/lib/dbus/machine-id
func GetLinuxMachineID() (string, error) {
	// 尝试读取 /etc/machine-id
	if data, err := os.ReadFile("/etc/machine-id"); err == nil {
		machineID := strings.TrimSpace(string(data))
		if machineID != "" {
			return machineID, nil
		}
	}

	// 尝试读取 /var/lib/dbus/machine-id
	if data, err := os.ReadFile("/var/lib/dbus/machine-id"); err == nil {
		machineID := strings.TrimSpace(string(data))
		if machineID != "" {
			return machineID, nil
		}
	}

	return "", fmt.Errorf("无法读取machine-id文件")
}
