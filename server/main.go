package main

import (
	"go.uber.org/zap"
	"os"
	
	"github.com/pddzl/kubefish/server/core"
	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/initialize"
)

func main() {
	global.KF_VP = core.Viper() // 初始化viper
	global.KF_LOG = core.Zap()  // 初始化zap日志
	zap.ReplaceGlobals(global.KF_LOG)
	global.KF_DB = initialize.Gorm() // gorm连接数据库
	if global.KF_DB == nil {
		global.KF_LOG.Error("mysql连接失败，退出程序")
		os.Exit(127)
	} else {
		initialize.RegisterTables(global.KF_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.KF_DB.DB()
		defer db.Close()
	}

	global.KF_K8S_Client, global.KF_K8S_CONFIG = initialize.K8sConnect()
	if global.KF_K8S_Client == nil {
		global.KF_LOG.Error("k8s连接失败，退出程序")
		os.Exit(128)
	}

	core.RunServer()
}
