package initializes

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"tranvancu185/vey-pos-ws/global"

	"go.uber.org/zap"
)

func InitApp() {
	// Set default time zone "Asia/Ho_Chi_Minh"
	time.Local, _ = time.LoadLocation("Asia/Ho_Chi_Minh")

	// Get config path
	var configPath string
	var appDataDir string
	appDir, _ := os.Getwd()
	if global.Config.Server.Mode != "dev" {
		configPath, _ = os.UserConfigDir()
		appDataDir = filepath.Join(configPath, global.Config.Server.AppName)
		_ = os.MkdirAll(appDataDir, 0755)
		appDir = filepath.Join(appDir, global.Config.Server.Path)
	} else {
		configPath, _ = os.Getwd()
		appDataDir = configPath
	}

	// set path db
	global.Config.Database.PathMigration = filepath.Join(appDir, global.Config.Database.PathMigration)
	global.Config.Database.Path = filepath.Join(appDataDir, global.Config.Database.Path)
	// Create App SQLite Folder
	_ = os.MkdirAll(global.Config.Database.Path, 0755)

	// Set App Data Folder
	global.Config.Path.AppDataDir = appDataDir
	global.Config.Path.AppDir = appDir
	// set path storage
	global.Config.Path.PathStorage = filepath.Join(appDataDir, global.Config.Path.PathStorage)
	global.Config.Path.PathAvatar = filepath.Join(global.Config.Path.PathStorage, global.Config.Path.PathAvatar)
	global.Config.Path.PathImage = filepath.Join(global.Config.Path.PathStorage, global.Config.Path.PathImage)
	global.Config.Path.PathFile = filepath.Join(global.Config.Path.PathStorage, global.Config.Path.PathFile)
	global.Config.Logger.PathLog = filepath.Join(global.Config.Path.PathStorage, global.Config.Logger.PathLog)

	// set path web
	global.Config.Path.PathWeb = filepath.Join(appDir, global.Config.Path.PathWeb)
	fmt.Println(global.Config.Path.PathWeb)
}

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}
