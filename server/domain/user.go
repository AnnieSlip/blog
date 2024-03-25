package domain

import (
	"blogs/server/common"
	"blogs/server/db"
	"blogs/server/request"
	"blogs/server/response"
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

func UserCreate(c context.Context, d *sqlx.DB, ucr request.User) (*int, error) {
	pass, err := HashPassword(ucr.Password)
	if err != nil {
		return nil, err
	}

	ucr.Password = pass
	req := ucr.DB()
	req.CreatedAt = time.Now()

	err = db.UserCreate(context.Background(), d, req)
	if err != nil {
		return nil, err
	}

	return nil, nil

}

func Login(c context.Context, d *sqlx.DB, req request.User) (*response.User, *string, error) {
	user, err := db.UserByEmail(c, d, req.Email)
	if err != nil {
		return nil, nil, err
	}
	if user == nil {
		return nil, nil, fmt.Errorf("no such user")
	}

	isPasswordValid := CheckPasswordHash(req.Password, user.Password)

	if !isPasswordValid {
		return nil, nil, fmt.Errorf("password is not correct")
	}

	token, err := common.GenerateToken(user.ID)
	if err != nil {
		return nil, nil, err
	}

	res := user.Response()

	return &res, &token, nil

}
