package logger

import (
	"fmt"
	"os"
	"time"
	"tranvancu185/vey-pos-ws/pkg/setting"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	// Khởi tạo logger
	logLevel := config.Level
	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "dpanic":
		level = zapcore.DPanicLevel
	case "panic":
		level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	fileName := fmt.Sprintf("%s_%s.log", time.Now().Format("2006-01-02"), config.FileLogName)
	filePath := fmt.Sprintf("%s/%s", config.PathLog, fileName)
	hook := lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,   //days
		Compress:   config.Compress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))}
}

// Format Logger
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()       // Config for production
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // Time format to ISO8601 => 2024-09-10T15:39:28.187+0700
	encodeConfig.TimeKey = "time"                          // Time key ts -> Time
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder // Level format to capital
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder // Caller format to short
	return zapcore.NewJSONEncoder(encodeConfig)
}
