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
)
