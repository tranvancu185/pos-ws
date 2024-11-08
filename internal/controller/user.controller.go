package controller

import (
	"fmt"
	"path/filepath"
	"strconv"
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/constants/messagecode"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/internal/model/rs"
	"tranvancu185/vey-pos-ws/internal/service"
	"tranvancu185/vey-pos-ws/pkg/response"
	"tranvancu185/vey-pos-ws/pkg/utils/ufile"
	urand "tranvancu185/vey-pos-ws/pkg/utils/urandom"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) GetProfile(c *gin.Context) {
	// Get user_id from context by j	wt
	id := c.GetInt64("user_id")

	result, err := uc.userService.GetUserByID(id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        result,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	// Get user_id from context by jwt
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := uc.userService.GetUserByID(idInt)
	if err != nil {
		c.Error(err)
		return
	}
	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        result,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (uc *UserController) GetListUsers(c *gin.Context) {
	var queryParams rq.GetListUsersRequest

	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.Error(err)
		return
	}

	result, err := uc.userService.GetListUsers(&queryParams)
	if err != nil {
		c.Error(err)
		return
	}

	data := rs.GetListResponse{
		Page:     queryParams.Page,
		PageSize: queryParams.PageSize,
		Total:    int64(len(result)),
		Rows:     result,
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        data,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (uc *UserController) UpdateProfile(c *gin.Context) {
	idParams := c.Param("id")
	id, errParams := strconv.ParseInt(idParams, 10, 64)
	if errParams != nil {
		c.Error(errParams)
		return
	}
	// Get username, password from request
	var req *rq.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	// Call userService.UpdateProfile
	err := uc.userService.UpdateUserByID(id, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (uc *UserController) UpdatePassword(c *gin.Context) {
	idParams := c.Param("id")
	id, errParams := strconv.ParseInt(idParams, 10, 64)
	if errParams != nil {
		c.Error(errParams)
		return
	}

	// Get username, password from request
	var req *rq.UpdateUserPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	// Call userService.UpdatePassword
	err := uc.userService.UpdateUserPasswordByID(id, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (uc *UserController) UpdateAvatar(c *gin.Context) {
	// Get user_id from context by jwt
	idParams := c.Param("id")
	id, errParams := strconv.ParseInt(idParams, 10, 64)
	if errParams != nil {
		c.Error(errParams)
		return
	}

	// Get User by id
	user, err := uc.userService.GetUserByID(id)
	if err != nil {
		c.Error(err)
		return
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		c.Error(err)
		return
	}

	// Tạo đường dẫn lưu trữ file
	filename := fmt.Sprintf("%d_%s_%s", id, urand.RandomTenDigit(), "avatar.png")
	filePath := filepath.Join(global.Config.Path.PathAvatar, filename)

	// Sử dụng một goroutine duy nhất để xử lý tất cả các tác vụ
	go func() {
		// Remove old avatar
		if user.UserAvatar != "" {
			avatarPath := filepath.Join(global.Config.Path.PathAvatar, user.UserAvatar)
			errRemove := ufile.RemoveFile(avatarPath)
			if errRemove != nil {
				global.SendLog("Error remove file", "error", errRemove)
			}
		}
	}()

	// Lưu file vào thư mục
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.Error(err) // Ghi nhận lỗi vào context
		return
	}

	// Call userService.UpdateAvatar
	if err := uc.userService.UpdateProfileAvatar(id, filename); err != nil {
		c.Error(err) // Ghi nhận lỗi vào context
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}
