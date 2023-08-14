package models

import "errors"

const (
	TypeNameTaskSendEmailWelcome = "task:send_email_welcome"
)

var (
	ErrorInternalServer = errors.New("internal server error")

	ErrorUserNotFound      = errors.New("user not found")
	ErrorUserWrongPassword = errors.New("wrong password")
	ErrorUserDuplicate     = errors.New("user duplicate")
)
