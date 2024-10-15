package electron

import (
	"encoding/json"
	"fmt"
)

const (
	ELEC_KEY_RECEIVE = "SERVER_RECEIVE"
	ELEC_KEY_SEND    = "SERVER_SEND"

	ELEC_TYPE_EVENT    = "event"
	ELEC_TYPE_LOG      = "log"
	ELEC_TYPE_ERROR    = "error"
	ELEC_TYPE_RESPONSE = "response"

	ELEC_EVENT_GET_VERSION  = "get-version"
	ELEC_EVENT_QUIT_APP     = "quit-app"
	ELEC_EVENT_MINIMIZE     = "minimize"
	ELEC_EVENT_MAXIMIZE     = "maximize"
	ELEC_EVENT_CHECK_UPDATE = "check-update"
	ELEC_EVENT_UPDATE_APP   = "update-app"
)

type DataMessage struct {
	Key     string  `json:"key"`
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func SendElecEvent(event string, data interface{}) {
	// Tạo message
	message := DataMessage{
		Key:  ELEC_KEY_SEND,
		Type: ELEC_TYPE_EVENT,
		Payload: Payload{
			Event: event,
			Data:  data,
		},
	}

	// Encode message thành JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		// Xử lý lỗi encode
		return
	}

	// Gửi JSON đã encode
	fmt.Println(string(jsonData))
}

func SendElecWithType(xtype string, event string, data interface{}) {
	// Tạo message
	message := DataMessage{
		Key:  ELEC_KEY_SEND,
		Type: xtype,
		Payload: Payload{
			Event: event,
			Data:  data,
		},
	}

	// Encode message thành JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		// Xử lý lỗi encode
		return
	}

	// Gửi JSON đã encode
	fmt.Println(string(jsonData))
}
