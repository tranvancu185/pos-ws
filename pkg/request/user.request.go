package request

type GetListUsersRequest struct {
	GetListRequest
	RoleIds   string `json:"role_id"`
	UserPhone string `json:"phone"`
	UserName  string `json:"name"`
	StatusIds string `json:"status"`
	FromDate  int64  `json:"from_date"`
	ToDate    int64  `json:"to_date"`
}

type UpdateUserRequest struct {
	RoleId          int64  `json:"role_id"`
	UserPhone       string `json:"phone"`
	StatusId        int64  `json:"status"`
	UserPassword    string `json:"password"`
	UserAvatar      string `json:"avatar"`
	UserDisplayName string `json:"display_name"`
}

type UpdateUserPasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password" validate:"required"`
	IsReset     bool   `json:"is_reset"`
}
