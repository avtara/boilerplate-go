package usecase

import (
	"context"
	"fmt"
	"github.com/avtara/boilerplate-go/internal/models"
	"github.com/avtara/boilerplate-go/internal/service"
	"github.com/avtara/boilerplate-go/utils"
	"golang.org/x/crypto/bcrypt"
	"strings"
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

func (u *userUsecase) Register(ctx context.Context, args models.RegisterUserRequest) (result models.RegisterUserResponse, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(args.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	args.Password = string(hashedPassword)
	args.Username = strings.ToLower(strings.Replace(args.Username, " ", "", -1))

	id, err := u.userRepository.Save(ctx, args)
	if err != nil {
		return
	}

	token, err := utils.GenerateToken(id)
	if err != nil {
		return
	}

	result = models.RegisterUserResponse{
		Username: args.Username,
		Email:    args.Email,
		Token:    token,
	}

	return
}

func (u *userUsecase) Auth(ctx context.Context, args models.LoginUserRequest) (result models.LoginUserResponse, err error) {
	user, err := u.userRepository.GetUserByUsernameOrEmail(ctx, args)
	fmt.Println(err, user)

	if err != nil {
		return
	}

	err = utils.VerifyPassword(args.Password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return
	}

	result = models.LoginUserResponse{
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return
}
