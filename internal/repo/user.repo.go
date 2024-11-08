package repo

import (
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/model/rq"
)

type IUserRepo interface {
	GetListUsers(params database.GetListUserByFilterParams) ([]database.GetListUserByFilterRow, error)
	GetUserByID(id int64) (*database.GetUserByIDRow, error)
	UpdateUser(id int64, params *rq.UpdateUserRequest) error
	UpdateUserPassword(id int64, password string) error
	UpdateUserAvatarByID(id int64, avatar string) error
}

type userRepo struct {
	sqlc *database.Queries
}

func NewUserRepo() IUserRepo {
	return &userRepo{
		sqlc: database.New(global.Mdbc),
	}
}

// GetListUsers implements IUserRepo.
func (ur *userRepo) GetListUsers(params database.GetListUserByFilterParams) ([]database.GetListUserByFilterRow, error) {
	users, err := ur.sqlc.GetListUserByFilter(ctx, params)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepo) GetTotalUsers(params database.GetTotalUserByFilterParams) (int64, error) {
	total, err := ur.sqlc.GetTotalUserByFilter(ctx, params)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (ur *userRepo) GetUserByID(id int64) (*database.GetUserByIDRow, error) {
	user, err := ur.sqlc.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepo) UpdateUser(id int64, params *rq.UpdateUserRequest) error {
	input := database.UpdateUserByIDParams{
		UserID: id,
	}

	if params.RoleId > 0 {
		input.UserRoleID = params.RoleId
	}

	if params.UserPhone != "" {
		input.UserPhone = params.UserPhone
	}

	if params.StatusId > 0 {
		input.UserStatus = params.StatusId
	}

	if params.UserDisplayName != "" {
		input.UserDisplayName = params.UserDisplayName
	}

	err := ur.sqlc.UpdateUserByID(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) UpdateUserPassword(id int64, password string) error {
	input := database.UpdateUserPasswordByIDParams{
		UserID:       id,
		UserPassword: password,
	}

	err := ur.sqlc.UpdateUserPasswordByID(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) UpdateUserAvatarByID(id int64, avatar string) error {
	input := database.UpdateUserAvatarByIDParams{
		UserID:     id,
		UserAvatar: avatar,
	}

	err := ur.sqlc.UpdateUserAvatarByID(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
