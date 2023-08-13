package service

import (
	"context"
	"github.com/avtara/boilerplate-go/internal/models"
)

type UserUsecase interface {
	GetLastLogin(ctx context.Context, args models.GetLastLoginRequest) (result models.GetLastLoginResponse, err error)
}
