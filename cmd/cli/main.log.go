package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("Start logger!")
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

func getWriterSync() zapcore.WriteSyncer {
	file, err := os.OpenFile("./log/server.log", os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
