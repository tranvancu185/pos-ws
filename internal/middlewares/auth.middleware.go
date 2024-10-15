package middlewares

import (
	"errors"
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/constants"
	"tranvancu185/vey-pos-ws/internal/constants/messagecode"
	"tranvancu185/vey-pos-ws/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(role_id int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Do something
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.Error(errors.New(messagecode.CODE_REQUIRE_AUTH))
			c.Abort()
			return
		}
		tokenString = tokenString[len("Bearer "):]
		// Parse và xác thực token
		token, err := jwt.ParseWithClaims(tokenString, &auth.JWTDataset{}, func(token *jwt.Token) (interface{}, error) {
			// Kiểm tra phương thức ký
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New(messagecode.CODE_UNAUTHORIZED)
			}

			// Trả về khóa bí mật để xác thực
			return []byte(global.Config.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			c.Error(errors.New(messagecode.CODE_FORBIDDEN))
			c.Abort()
			return
		}

		// Lưu trữ thông tin claims vào context để sử dụng sau này
		claims, ok := token.Claims.(*auth.JWTDataset)
		if !ok {
			c.Error(errors.New(messagecode.CODE_FORBIDDEN))
			c.Abort()
			return
		}

		if role_id > 0 && claims.RoleID != constants.USER_ROLEID_ADMIN {
			if claims.RoleID != role_id {
				c.Error(errors.New(messagecode.CODE_FORBIDDEN))
				c.Abort()
				return
			}
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.RoleID)
		c.Next()
	}
}
