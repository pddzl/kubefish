package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"

	"github.com/pddzl/kubefish/server/core/internal"
	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/utils"
)

func Zap() (logger *zap.Logger) {
	// 判断是否有Director文件夹
	if ok, _ := utils.PathExists(global.KF_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.KF_CONFIG.Zap.Director)
		_ = os.Mkdir(global.KF_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.KF_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
