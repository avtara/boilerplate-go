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
)
