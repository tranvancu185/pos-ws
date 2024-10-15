package initializes

import (
	"log"
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/pkg/electron"
)

func Run() {
	// Load the configuration file
	LoadConfig()
	// Create App Config
	InitApp()
	// Load the message file
	LoadMessage()
	// Initialize the logger
	InitLogger()
	// Initialize the mysql
	InitMysqlC()
	global.Logger.Info("Loading server configuration succeeded!!!!!")
	// Initialize the router
	r := InitRouter()
	// // Initialize the electron
	// InitElectron()
	// Run the server
	electron.SendElecEvent("SERVER_STARTED", "Server running in: http://localhost"+global.Config.Server.Port)
	err := r.Run(global.Config.Server.Port)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}

}
