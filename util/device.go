package util

import (
	"context"
	"fmt"

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
