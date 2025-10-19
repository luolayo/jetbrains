package main

import (
	"context"
	"fmt"
	"jetbrains/util"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// 初始化全局路径变量
	if err := util.InitGlobalPaths(); err != nil {
		fmt.Printf("初始化全局路径失败: %v\n", err)
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetDeviceID 获取设备唯一标识符
func (a *App) GetDeviceID() util.DeviceInfo {
	deviceUtil := &util.App{}
	return deviceUtil.GetDeviceID()
}

// Actions 自动激活
func (a *App) Actions() util.ActionsType {
	err := util.RemoveEnvOther()
	if err != nil {
		return util.ActionsType{
			Error: []string{err.Error()},
		}
	}
	err = util.ReadAllJetbrainsProducts()
	if err != nil {
		return util.ActionsType{
			Error: []string{err.Error()},
		}
	}
	var actions util.ActionsType
	return actions.Actions()
}

// Clean 激活清理
func (a *App) Clean() string {
	_ = util.RemoveEnvOther()
	_ = util.ReadAllJetbrainsProducts()
	return "清理完成"
}
