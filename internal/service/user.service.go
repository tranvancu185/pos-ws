package service

import (
	"errors"
	"strconv"
	"strings"
	"tranvancu185/vey-pos-ws/internal/constants"
	"tranvancu185/vey-pos-ws/internal/constants/messagecode"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/pkg/request"
	"tranvancu185/vey-pos-ws/pkg/utils"
)

type IUserService interface {
	GetListUsers(params *request.GetListUsersRequest) ([]database.GetListUserByFilterRow, error)
	GetUserByID(id int64) (*database.GetUserByIDRow, error)
	UpdateUserByID(id int64, params *request.UpdateUserRequest) error
	UpdateUserPasswordByID(id int64, params *request.UpdateUserPasswordRequest) error
	UpdateProfileAvatar(id int64, avatar string) error
}

type userService struct {
	userRepo repo.IUserRepo
	authRepo repo.IAuthRepo
}

func NewUserService(
	userRepo repo.IUserRepo,
	authRepo repo.IAuthRepo,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// GetListUsers implements IUserService.
func (us *userService) GetListUsers(params *request.GetListUsersRequest) ([]database.GetListUserByFilterRow, error) {
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

	return (us.userRepo).GetListUsers(input)
}

func (us *userService) GetUserByID(id int64) (*database.GetUserByIDRow, error) {
	return (us.userRepo).GetUserByID(id)
}

func (us *userService) UpdateUserByID(id int64, params *request.UpdateUserRequest) error {
	return (us.userRepo).UpdateUser(id, params)
}

func (us *userService) UpdateUserPasswordByID(id int64, params *request.UpdateUserPasswordRequest) error {

	if params.IsReset {

		return (us.userRepo).UpdateUserPassword(id, constants.USER_DEFAULT_PASSWORD)

	} else {
		// Check old password
		user, err := (us.authRepo).GetUserInfo(database.GetUserProfileParams{
			UserID: id,
		})
		if err != nil {
			return err
		}

		// Check password
		if !utils.CompareHash(user.UserPassword, params.OldPassword) {
			return errors.New(messagecode.CODE_UPDATE_PASSWORD_INVALID)
		}

		return (us.userRepo).UpdateUserPassword(id, params.NewPassword)
	}
}

func (us *userService) UpdateProfileAvatar(id int64, avatar string) error {
	return (us.userRepo).UpdateUserAvatarByID(id, avatar)
}
