package message

import "tranvancu185/vey-pos-ws/global"

func GetMessage(message_code string) string {
	return global.Message[message_code]
}
