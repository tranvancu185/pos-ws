package repo

import (
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/pkg/utils/ucrypto"
	"tranvancu185/vey-pos-ws/pkg/utils/utime"
)

type IAuthRepo interface {
	GetUserInfo(params rq.GetProfileRequest) (*database.GetUserProfileRow, error)
	CreateUser(userData *rq.RegisterRequest) (int64, error)
	CheckUserExist(user_name string, user_phone string) bool
}

type authRepo struct {
	sqlc *database.Queries
}

func NewAuthRepo() IAuthRepo {
	return &authRepo{
		sqlc: database.New(global.Mdbc),
	}
}

func (ar *authRepo) GetUserInfo(params rq.GetProfileRequest) (*database.GetUserProfileRow, error) {
	var input database.GetUserProfileParams

	if params.UserID != 0 {
		input.UserID = params.UserID
	}
	if params.UserName != "" {
		input.UserName = params.UserName
	}
	if params.UserPhone != "" {
		input.UserPhone = params.UserPhone
	}

	user, err := ar.sqlc.GetUserProfile(ctx, input)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ar *authRepo) CreateUser(userData *rq.RegisterRequest) (int64, error) {

	userForm := database.CreateUserParams{
		UserName:        userData.UserName,
		UserPassword:    ucrypto.GetHash(userData.UserPassword),
		UserDisplayName: userData.UserDisplayName,
		UserPhone:       userData.UserPhone,
		UserAvatar:      userData.UserAvatar,
		CreatedAt:       utime.GetCurrentTimeUnix(),
		UpdatedAt:       utime.GetCurrentTimeUnix(),
	}

	if userData.UserStatus != 0 {
		userForm.UserStatus = userData.UserStatus
	}

	userId, err := ar.sqlc.CreateUser(ctx, userForm)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (ar *authRepo) CheckUserExist(user_name string, user_phone string) bool {
	total_user, err := ar.sqlc.CheckUserExist(ctx, database.CheckUserExistParams{
		UserName:  user_name,
		UserPhone: user_phone,
	})
	if err != nil || total_user > 0 {
		return true
	}
	return false
}
