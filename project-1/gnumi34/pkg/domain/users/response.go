package users

type Response struct {
	ID       uint
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}
