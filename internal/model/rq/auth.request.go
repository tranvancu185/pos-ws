package rq

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=30"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	UserName        string `json:"username" validate:"required,min=3,max=30"`
	UserPassword    string `json:"password" validate:"required"`
	UserDisplayName string `json:"user_display_name" validate:"required,min=3"`
	UserPhone       string `json:"user_phone" validate:"required,min=9,max=15,numeric"`
	UserAvatar      string `json:"user_avatar"`
	UserRoleID      int64  `json:"user_role_id"`
	UserStatus      int64  `json:"user_status"`
}

type GetProfileRequest struct {
	UserID    int64  `json:"user_id"`
	UserName  string `json:"user_name"`
	UserPhone string `json:"user_phone"`
}
