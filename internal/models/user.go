package models

import "time"

type (
	GetLastLoginRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	GetLastLoginResponse struct {
		LastLogin time.Time `db:"last_login"`
	}

	RegisterUserRequest struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	RegisterUserResponse struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	}

	LoginUserRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password" binding:"required"`
	}

	LoginUserResponse struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	}

	User struct {
		ID       int64  `db:"user_id"`
		Username string `db:"username"`
		Email    string `db:"email"`
		Password string `db:"password"`
	}
)
