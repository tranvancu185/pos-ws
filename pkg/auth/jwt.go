package auth

import (
	"time"
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/pkg/utils/utime"

	"github.com/golang-jwt/jwt/v4"
)

type JWTDataset struct {
	UserID     int64 `json:"user_id"`
	UserStatus int64 `json:"user_status"`
	RoleID     int64 `json:"role_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId int64, roleId int64, userStatus int64) (string, error) {
	// Set expiration time
	expirationTime := utime.GetCurrentTime().Add(24 * time.Hour) // ===> 24h

	// Create Claims JWT
	claims := JWTDataset{
		UserID:     userId,
		UserStatus: userStatus,
		RoleID:     roleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	tokenString, err := token.SignedString([]byte(global.Config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
