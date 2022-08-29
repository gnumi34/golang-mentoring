package auth

type ValidateUserRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
