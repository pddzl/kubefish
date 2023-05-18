package internal

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"

	"github.com/pddzl/kubefish/server/global"
)

type lumberjackLogs struct{}

var LumberjackLogs = new(lumberjackLogs)

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (l *lumberjackLogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
	fileWriter := &lumberjack.Logger{
		Filename:   path.Join(global.KF_CONFIG.Zap.Director, level+".log"),
		MaxSize:    global.KF_CONFIG.RotateLogs.MaxSize,
		MaxBackups: global.KF_CONFIG.RotateLogs.MaxBackups,
		MaxAge:     global.KF_CONFIG.RotateLogs.MaxAge,
		Compress:   global.KF_CONFIG.RotateLogs.Compress,
	}

	if global.KF_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
