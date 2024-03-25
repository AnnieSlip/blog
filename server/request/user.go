package request

import "blogs/server/db"

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (uc User) DB() db.User {
	return db.User{
		Username: uc.Username,
		Email:    uc.Email,
		Password: uc.Password,
	}
}
