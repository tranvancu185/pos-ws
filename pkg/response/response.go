package response

import (
	"net/http"
	"tranvancu185/vey-pos-ws/internal/uconst/messagecode"
	"tranvancu185/vey-pos-ws/pkg/message"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status      int         `json:"status"`       // Status code
	MessageCode string      `json:"message_code"` // Message code
	Message     string      `json:"message"`      // Message
	Data        interface{} `json:"data"`         // Data return
}

type ParamsResponse struct {
	Status      int
	Message     string
	MessageCode string
	Data        interface{}
	Error       error
}

const (
	// Internal Satuts code
	StatusCodeSuccess      = 200 // Success
	SatusInternalError     = 500 // Internal error
	StatusCodeRequireAuth  = 401 // require_auth
	StatusCodeUnauthorized = 402 // Unauthorized
	StatusCodeForbidden    = 403 // Forbidden
	StatusBadRequest       = 400 // Bad request
	// Validate Satuts code
	StatusValidateError = 422 // Validate error
)

func SuccessResponse(c *gin.Context, params ParamsResponse) {
	var messageTmp string
	var status int
	var data interface{}
	var message_code string
	if params.Status != 0 {
		status = params.Status
	} else {
		status = http.StatusOK
	}

	if params.MessageCode != "" {
		message_code = params.MessageCode
		messageTmp = message.GetMessage(params.MessageCode)
	} else {
		messageTmp = message.GetMessage(messagecode.CODE_SUCCESS)
	}

	if params.Message != "" {
		messageTmp = params.Message
	}

	if params.Data != nil {
		data = params.Data
	} else {
		data = nil
	}

	c.JSON(http.StatusOK, Response{
		Status:      status,
		MessageCode: message_code,
		Message:     messageTmp,
		Data:        data,
	})
}

func ErrorResponse(c *gin.Context, params ParamsResponse) {

	var messageTmp string
	var status int
	var message_code string
	var data interface{}

	if params.Status != 0 {
		status = params.Status
	} else {
		status = http.StatusInternalServerError
	}

	if params.MessageCode != "" {
		message_code = params.MessageCode
		messageTmp = message.GetMessage(params.MessageCode)
	} else {
		messageTmp = message.GetMessage(messagecode.CODE_SUCCESS)
	}

	if params.Message != "" {
		messageTmp = params.Message
	}

	if params.Data != nil {
		data = params.Data
	} else {
		data = nil
	}

	c.JSON(http.StatusBadRequest, Response{
		Status:      status,
		MessageCode: message_code,
		Message:     messageTmp,
		Data:        data,
	})
}
