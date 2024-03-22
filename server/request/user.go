package request

import "blogs/server/db"

type UserCreate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (uc UserCreate) DB() db.User {
	return db.User{
		Username: uc.Username,
		Email:    uc.Email,
		Password: uc.Password,
	}
}
