package response

type LoginResponse struct {
	Token  string `json:"token"`
	UserID int64  `json:"user_id"`
}

type RegisterResponse struct {
	UserID int64 `json:"user_id"`
}
