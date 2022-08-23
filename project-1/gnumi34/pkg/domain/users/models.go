package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
}

func (u *User) ToResponse() *Response {
	return &Response{
		ID:       u.ID,
		UserName: u.UserName,
		Email:    u.Email,
	}
}

func ToMultipleResponse(req []User) (out []Response) {
	for idx := range req {
		out = append(out, Response{
			ID:       req[idx].ID,
			UserName: req[idx].UserName,
			Email:    req[idx].Email,
		})
	}

	return out
}
