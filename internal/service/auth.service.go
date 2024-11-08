package service

import (
	"errors"
	"tranvancu185/vey-pos-ws/internal/constants"
	"tranvancu185/vey-pos-ws/internal/constants/messagecode"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/internal/model/rs"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/pkg/auth"
	"tranvancu185/vey-pos-ws/pkg/utils/ucrypto"
)

type IAuthService interface {
	Login(userName, password string) (*rs.LoginResponse, error)
	Register(user_data *rq.RegisterRequest) (int64, error)
}

type authService struct {
	authRepo repo.IAuthRepo
	userRepo repo.IUserRepo
}

func NeuAuthService(
	authRepo repo.IAuthRepo,
	userRepo repo.IUserRepo,
) IAuthService {
	return &authService{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (as *authService) Login(userName, password string) (*rs.LoginResponse, error) {
	// Check user exist
	userInfo, err := as.authRepo.GetUserInfo(rq.GetProfileRequest{
		UserName: userName,
	})
	if err != nil {
		return nil, err
	}
	// Check password
	if !ucrypto.CompareHash(userInfo.UserPassword, password) {
		return nil, errors.New(messagecode.CODE_INVALID_LOGIN)
	}
	// Create token
	token, err := auth.GenerateJWT(userInfo.UserID, userInfo.UserRoleID, userInfo.UserStatus)
	if err != nil {
		return nil, err
	}
	result := &rs.LoginResponse{
		UserID: userInfo.UserID,
		Token:  token,
	}
	// Return token
	return result, nil
}

func (as *authService) Register(userData *rq.RegisterRequest) (int64, error) {
	// Check user exist
	isExist := as.authRepo.CheckUserExist(userData.UserName, userData.UserPhone)
	if isExist == constants.BOOL_TRUE {
		return 0, errors.New(messagecode.CODE_USER_EXIST)
	}

	// Create user
	userId, err := as.authRepo.CreateUser(userData)
	if err != nil {
		return 0, err
	}

	// Return token
	return userId, nil
}
