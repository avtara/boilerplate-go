package usecase

import (
	"context"
	"github.com/avtara/boilerplate-go/internal/models"
	"github.com/avtara/boilerplate-go/internal/service"
)

type userUsecase struct {
	userRepository service.UserRepository
}

func NewUserUseCase(userRepository service.UserRepository) service.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) GetLastLogin(ctx context.Context, args models.GetLastLoginRequest) (result models.GetLastLoginResponse, err error) {
	result, err = u.userRepository.GetLastLoginByUsernameOrEmail(ctx, args)

	return
}
