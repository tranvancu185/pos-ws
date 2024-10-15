package initializes

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/pkg/electron"
)

type ElectronMessage struct {
	Key     string `json:"key"`
	Type    string `json:"type"`
	Payload struct {
		Event string      `json:"event"`
		Data  interface{} `json:"data"`
	} `json:"payload"`
}

func InitElectron() {
	// Initialize the electron
	go func() {
		// Đọc dữ liệu từ stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			message := scanner.Text()
			global.Logger.Info(message)
			if strings.Contains(message, electron.ELEC_KEY_RECEIVE) {
				// Parse JSON message
				var msgData ElectronMessage

				err := json.Unmarshal(scanner.Bytes(), &msgData)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Lỗi khi parse JSON:", err)
					continue
				}
				// Xử lý message
				handleEvent(msgData)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Lỗi khi đọc stdin:", err)
		}
	}()
}

func handleEvent(msg ElectronMessage) {
	switch msg.Payload.Event {
	case electron.ELEC_EVENT_GET_VERSION:
		version := msg.Payload.Data.(string)
		fmt.Println("Version: ", version)
		electron.SendElecWithType(electron.ELEC_TYPE_RESPONSE, electron.ELEC_EVENT_GET_VERSION, "1.0.0")
	case electron.ELEC_EVENT_QUIT_APP:
		os.Exit(0)
	case electron.ELEC_EVENT_MINIMIZE:
		break
	case electron.ELEC_EVENT_MAXIMIZE:
		break
	case electron.ELEC_EVENT_CHECK_UPDATE:
		break
	case electron.ELEC_EVENT_UPDATE_APP:
		break
	default:
		fmt.Println("Unknown event")
	}
}
