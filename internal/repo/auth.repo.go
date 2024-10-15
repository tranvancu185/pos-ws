package repo

import (
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
)

type IAuthRepo interface {
	GetUserInfo(params database.GetUserProfileParams) (*database.GetUserProfileRow, error)
	CreateUser(user_form database.CreateUserParams) (int64, error)
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

func (ar *authRepo) GetUserInfo(params database.GetUserProfileParams) (*database.GetUserProfileRow, error) {
	user, err := ar.sqlc.GetUserProfile(ctx, params)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ar *authRepo) CreateUser(user_form database.CreateUserParams) (int64, error) {
	userId, err := ar.sqlc.CreateUser(ctx, user_form)
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
