package repo

import (
	"context"
	"tranvancu185/vey-pos-ws/global"

	"go.uber.org/zap"
)

const (
	// DEFAULT_LIMIT default limit
	DEFAULT_LIMIT = 10
	// DEFAULT_OFFSET default offset
	DEFAULT_OFFSET = 0
	// DEFAULT_SORT default sort
	DEFAULT_SORT = "desc"
	// DEFAULT_PAGE default page
	DEFAULT_PAGE = 1
	// DEFAULT_PAGE_SIZE default page size
	DEFAULT_PAGE_SIZE = 20

	// TYPE_LOG_INFO type log info
	TYPE_LOG_INFO = "info"
	// TYPE_LOG_ERROR type log error
	TYPE_LOG_ERROR = "error"
	// TYPE_LOG_WARN type log warn
	TYPE_LOG_WARN = "warn"
)

var (
	ctx = context.Background()
)

func Logger(message string, typeLog string, err error) {
	switch typeLog {
	case "info":
		global.Logger.Info(message)
	case "error":
		if err != nil {
			global.Logger.Error(message, zap.Error(err))
			return
		}
		global.Logger.Error(message)
	case "warn":
		global.Logger.Warn(message)
	}
}
