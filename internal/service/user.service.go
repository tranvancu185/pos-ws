package service

import (
	"errors"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/internal/uconst"
	"tranvancu185/vey-pos-ws/internal/uconst/messagecode"
	"tranvancu185/vey-pos-ws/pkg/utils/ucrypto"
)

type IUserService interface {
	GetListUsers(params *rq.GetListUsersRequest) ([]database.GetListUserByFilterRow, error)
	GetUserByID(id int64) (*database.GetUserByIDRow, error)
	UpdateUserByID(id int64, params *rq.UpdateUserRequest) error
	UpdateUserPasswordByID(id int64, params *rq.UpdateUserPasswordRequest) error
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
func (us *userService) GetListUsers(params *rq.GetListUsersRequest) ([]database.GetListUserByFilterRow, error) {
	return (us.userRepo).GetListUsers(params)
}

func (us *userService) GetUserByID(id int64) (*database.GetUserByIDRow, error) {
	return (us.userRepo).GetUserByID(id)
}

func (us *userService) UpdateUserByID(id int64, params *rq.UpdateUserRequest) error {
	return (us.userRepo).UpdateUser(id, params)
}

func (us *userService) UpdateUserPasswordByID(id int64, params *rq.UpdateUserPasswordRequest) error {
	if params.IsReset {
		return (us.userRepo).UpdateUserPassword(id, uconst.USER_DEFAULT_PASSWORD)
	} else {
		// Check old password
		user, err := (us.authRepo).GetUserInfo(rq.GetProfileRequest{
			UserID: id,
		})
		if err != nil {
			return err
		}
		// Check password
		if !ucrypto.CompareHash(user.UserPassword, params.OldPassword) {
			return errors.New(messagecode.CODE_UPDATE_PASSWORD_INVALID)
		}

		return (us.userRepo).UpdateUserPassword(id, params.NewPassword)
	}
}

func (us *userService) UpdateProfileAvatar(id int64, avatar string) error {
	return (us.userRepo).UpdateUserAvatarByID(id, avatar)
}
