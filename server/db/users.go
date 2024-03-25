package db

import (
	"blogs/server/response"
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type Users []User

func UserList(c context.Context, db *sqlx.DB) (*Users, error) {
	var us Users
	err := selectContext(c, db, &us, `
	SELECT
	   id,
	   username,
	   email,
	   password,
	   created_at
	FROM users`)

	if err != nil {
		return nil, err
	}
	return &us, nil
}

func UserByEmail(c context.Context, db *sqlx.DB, email string) (*User, error) {
	var u User
	err := getContext(c, db, &u, `
	SELECT
	   id,
	   username,
	   email, 
	   password,
	   created_at
	FROM users
	WHERE email=$1
	`, email)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &u, nil
}

func UserCreate(c context.Context, db *sqlx.DB, req User) error {
	_, err := db.NamedExecContext(c, `
        INSERT INTO users
            (
			 username,
			 email,
			 password,
             created_at)
        VALUES
            (:username, :email, :password, :created_at)
    `, req)
	if err != nil {
		return err
	}
	return nil
}

func (u User) Response() response.User {
	return response.User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
	}
}
