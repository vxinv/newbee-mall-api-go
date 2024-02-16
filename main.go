package main

import (
	"main/api/v1"
	"main/core"
	"main/global"
	"main/initialize"
)

func main() {

	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库

	v1.RunWindowsServer()
}
