package controller

import (
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/internal/model/rs"
	"tranvancu185/vey-pos-ws/internal/service"
	"tranvancu185/vey-pos-ws/internal/uconst/messagecode"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(
	authService service.IAuthService,
) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	// Get username, password from request
	var req rq.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	// Validate username, password
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.Error(err)
		return
	}
	// Call authService.Login
	result, err := ac.authService.Login(req.Username, req.Password)
	if err != nil {
		c.Error(err)
		return
	}
	// Return token
	data := rs.LoginResponse{
		UserID: result.UserID,
		Token:  result.Token,
	}
	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        data,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (ac *AuthController) Register(c *gin.Context) {
	// Get username, password from request
	var req rq.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	// Validate username, password
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.Error(err)
		return
	}
	// Call authService.Register
	userId, err := ac.authService.Register(&req)
	if err != nil {
		c.Error(err)
		return
	}

	result := rs.RegisterResponse{
		UserID: userId,
	}

	// Return token
	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        result,
		MessageCode: messagecode.CODE_REGISTER_SUCCESS,
	})
}
