package initializes

import (
	"fmt"
	"os"
	"path/filepath"
	"tranvancu185/vey-pos-ws/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	// Load the configuration file
	appDir, errAppDir := os.Getwd()
	if errAppDir != nil {
		panic(fmt.Errorf("unable to get the current directory: %v", errAppDir))
	}
	viper := viper.New()
	viper.AddConfigPath(filepath.Join(appDir, "config"))
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	// Read the configuration file
	isFromServer := true
	isFromElectron := false
	isDebugMode := false
	err := viper.ReadInConfig()
	if err != nil {
		isFromServer = false
		isFromElectron = true
	}

	if !isFromServer && isFromElectron {
		viper.AddConfigPath("./server/config")
		viper.SetConfigName("electron")
		err = viper.ReadInConfig()
		if err != nil {
			isFromElectron = false
		}
	}

	if !isFromElectron && !isFromServer {
		viper.AddConfigPath("./resources/app.asar.unpacked/server/config")
		viper.SetConfigName("production")
		err := viper.ReadInConfig()
		if err != nil {
			isDebugMode = true
		}
	}

	if isDebugMode {
		viper.AddConfigPath("../../config/")
		viper.SetConfigName("debug-dev")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	// Config struct
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("unable to decode configuration: %v", err))
	}
}
