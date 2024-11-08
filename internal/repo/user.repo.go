package repo

import (
	"strconv"
	"strings"
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/model/rq"
)

type IUserRepo interface {
	GetListUsers(params *rq.GetListUsersRequest) ([]database.GetListUserByFilterRow, error)
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
func (ur *userRepo) GetListUsers(params *rq.GetListUsersRequest) ([]database.GetListUserByFilterRow, error) {
	var input database.GetListUserByFilterParams

	if params.PageSize != 0 {
		input.Limit = params.PageSize
	} else {
		input.Limit = 10
	}

	if params.Page != 0 {
		input.Offset = (params.Page - 1) * params.PageSize
	} else {
		input.Offset = 0
	}

	if params.UserName != "" {
		input.UserDisplayName = params.UserName
	}

	if params.UserPhone != "" {
		input.UserPhone = params.UserPhone
	}

	if params.StatusIds != "" {
		status := strings.Split(params.StatusIds, ",")
		for _, s := range status {
			id, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return nil, err
			}
			input.UserStatusIds = append(input.UserStatusIds, id)
		}
	}

	if params.RoleIds != "" {
		status := strings.Split(params.RoleIds, ",")
		for _, s := range status {
			id, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return nil, err
			}
			input.UserRoleIds = append(input.UserRoleIds, id)
		}
	}

	if params.FromDate > 0 {
		input.CreatedAt = params.FromDate
	}

	if params.ToDate > 0 {
		input.CreatedAt_2 = params.ToDate
	}

	users, err := ur.sqlc.GetListUserByFilter(ctx, input)
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
