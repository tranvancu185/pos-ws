package service

import (
	"errors"
	"tranvancu185/vey-pos-ws/internal/constants"
	"tranvancu185/vey-pos-ws/internal/constants/messagecode"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/pkg/auth"
	"tranvancu185/vey-pos-ws/pkg/request"
	"tranvancu185/vey-pos-ws/pkg/response"
	"tranvancu185/vey-pos-ws/pkg/utils"
)

type IAuthService interface {
	Login(userName, password string) (*response.LoginResponse, error)
	Register(user_data *request.RegisterRequest) (int64, error)
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

func (as *authService) Login(userName, password string) (*response.LoginResponse, error) {
	// Check user exist
	userInfo, err := as.authRepo.GetUserInfo(database.GetUserProfileParams{
		UserName: userName,
	})
	if err != nil {
		return nil, err
	}
	// Check password
	if !utils.CompareHash(userInfo.UserPassword, password) {
		return nil, errors.New(messagecode.CODE_INVALID_LOGIN)
	}
	// Create token
	token, err := auth.GenerateJWT(userInfo.UserID, userInfo.UserRoleID, userInfo.UserStatus)
	if err != nil {
		return nil, err
	}
	result := &response.LoginResponse{
		UserID: userInfo.UserID,
		Token:  token,
	}
	// Return token
	return result, nil
}

func (as *authService) Register(userData *request.RegisterRequest) (int64, error) {
	// Check user exist
	isExist := as.authRepo.CheckUserExist(userData.UserName, userData.UserPhone)
	if isExist == constants.BOOL_TRUE {
		return 0, errors.New(messagecode.CODE_USER_EXIST)
	}
	// Create user
	userForm := database.CreateUserParams{
		UserName:        userData.UserName,
		UserPassword:    utils.GetHash(userData.UserPassword),
		UserDisplayName: userData.UserDisplayName,
		UserPhone:       userData.UserPhone,
		UserAvatar:      userData.UserAvatar,
		UserStatus:      userData.UserStatus,
		CreatedAt:       utils.GetCurrentTimeUnix(),
		UpdatedAt:       utils.GetCurrentTimeUnix(),
	}

	userId, err := as.authRepo.CreateUser(userForm)
	if err != nil {
		return 0, err
	}

	// Return token
	return userId, nil
}
