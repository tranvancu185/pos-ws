package initializes

import (
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
