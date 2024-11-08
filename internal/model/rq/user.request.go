package rq

type GetListUsersRequest struct {
	PageSize  int64  `form:"page_size"`
	Page      int64  `form:"page"`
	Total     int64  `form:"total"`
	RoleIds   string `form:"role_id"`
	UserPhone string `form:"phone"`
	UserName  string `form:"name"`
	StatusIds string `form:"status"`
	FromDate  int64  `form:"from_date"`
	ToDate    int64  `form:"to_date"`
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
