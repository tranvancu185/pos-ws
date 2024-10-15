package middlewares

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/constants/messagecode"
	"tranvancu185/vey-pos-ws/pkg/message"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginErr := range c.Errors {
			// Ghi log lỗi
			global.SendLog("Error middleware found the error!", "error", ginErr)

			// Xử lý lỗi sqlc
			if errors.Is(ginErr.Err, sql.ErrNoRows) {
				response.ErrorResponse(c, response.ParamsResponse{
					Status:      response.StatusBadRequest,
					MessageCode: messagecode.CODE_USER_NOT_FOUND,
				})
				c.Abort()
				return
			} else if sqliteErr, ok := ginErr.Err.(sqlite3.Error); ok {

				switch sqliteErr.Code {

				case sqlite3.ErrConstraint:
					switch sqliteErr.ExtendedCode {
					case sqlite3.ErrConstraintUnique:
						response.ErrorResponse(c, response.ParamsResponse{
							Status:      response.StatusBadRequest,
							MessageCode: messagecode.CODE_UNIQUE_CONSTRAINT,
						})
						c.Abort()
						return
					case sqlite3.ErrConstraintForeignKey:
						response.ErrorResponse(c, response.ParamsResponse{
							Status:      response.StatusBadRequest,
							MessageCode: messagecode.CODE_FOREIGN_KEY_CONSTRAINT,
						})
						c.Abort()
						return
					case sqlite3.ErrConstraintCheck:
						response.ErrorResponse(c, response.ParamsResponse{
							Status:      response.StatusBadRequest,
							MessageCode: messagecode.CODE_CHECK_CONSTRAINT,
						})
						c.Abort()
						return
					default:
						response.ErrorResponse(c, response.ParamsResponse{
							Status:      response.SatusInternalError,
							MessageCode: messagecode.CODE_INTERNAL_ERR,
							Message:     ginErr.Error(),
						})
						c.Abort()
						return
					}

				default:
					response.ErrorResponse(c, response.ParamsResponse{
						Status:      response.SatusInternalError,
						MessageCode: messagecode.CODE_INTERNAL_ERR,
						Message:     ginErr.Error(),
					})
					c.Abort()
					return
				}

			} else {
				if msg := message.GetMessage(ginErr.Error()); msg != "" {
					response.ErrorResponse(c, response.ParamsResponse{
						Status:      response.StatusBadRequest,
						Message:     msg,
						MessageCode: ginErr.Error(),
					})
					c.Abort()
					return
				}

				if jsonErr, ok := ginErr.Err.(*json.UnmarshalTypeError); ok {
					msg := fmt.Sprintf("Invalid value for field %s", jsonErr.Field)
					response.ErrorResponse(c, response.ParamsResponse{
						Status:      response.StatusValidateError,
						Message:     msg,
						MessageCode: messagecode.CODE_PARAM_INVALID,
						Data:        ginErr.Error(),
					})
					c.Abort()
					return
				}

				// Xử lý lỗi validate
				switch ginErr.Type {
				case gin.ErrorTypeBind:
					response.ErrorResponse(c, response.ParamsResponse{
						Status:      response.StatusValidateError,
						MessageCode: messagecode.CODE_PARAM_INVALID,
					})
					c.Abort()
					return
				default:
					response.ErrorResponse(c, response.ParamsResponse{
						Status:      response.StatusValidateError,
						MessageCode: messagecode.CODE_INTERNAL_ERR,
						Message:     ginErr.Error(),
					})
					c.Abort()
					return
				}
			}
		}
	}
}
