package service

import (
	"context"
	"github.com/avtara/boilerplate-go/internal/models"
)

type UserRepository interface {
	GetLastLoginByUsernameOrEmail(ctx context.Context, args models.GetLastLoginRequest) (result models.GetLastLoginResponse, err error)
	Save(ctx context.Context, args models.RegisterUserRequest) (id int64, err error)
	GetUserByUsernameOrEmail(ctx context.Context, args models.LoginUserRequest) (user models.User, err error)
}
