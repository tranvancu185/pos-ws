package initializes

import (
	"fmt"
	"tranvancu185/vey-pos-ws/global"

	"github.com/spf13/viper"
)

func LoadMessage() {
	// Load the message file
	viper := viper.New()
	viper.AddConfigPath("./config/message")
	viper.SetConfigName("message")
	viper.SetConfigType("yaml")

	// Read the message file
	isFromServer := true
	isFromElectron := false
	isDebugMode := false
	err := viper.ReadInConfig()
	if err != nil {
		isFromServer = false
		isFromElectron = true
	}

	if !isFromServer && isFromElectron {
		viper.AddConfigPath("./server/config/message")
		err = viper.ReadInConfig()
		if err != nil {
			isFromElectron = false
		}
	}

	if !isFromElectron && !isFromServer {
		viper.AddConfigPath("./resources/app.asar.unpacked/server/config/message")
		err := viper.ReadInConfig()
		if err != nil {
			isDebugMode = true
			// panic(fmt.Errorf("fatal error message file: %w", err))
		}
	}

	if isDebugMode {
		viper.AddConfigPath("../../config/message")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	// Config struct
	err = viper.Unmarshal(&global.Message)
	if err != nil {
		panic(fmt.Errorf("unable to decode message: %v", err))
	}
}
