package global

import (
	"database/sql"
	"tranvancu185/vey-pos-ws/pkg/logger"
	"tranvancu185/vey-pos-ws/pkg/setting"

	"go.uber.org/zap"
)

var (
	Config  setting.Config
	Mdbc    *sql.DB
	Logger  *logger.LoggerZap
	Message map[string]string
)

func SendLog(message string, typeLog string, err error) {
	switch typeLog {
	case "info":
		Logger.Info(message)
	case "error":
		if err != nil {
			Logger.Error(message, zap.Error(err))
			return
		}
		Logger.Error(message)
	case "warn":
		Logger.Warn(message)
	}
}
