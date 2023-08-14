package usecase

import (
	"context"
	"github.com/avtara/boilerplate-go/internal/models"
	"github.com/avtara/boilerplate-go/internal/service"
	"github.com/avtara/boilerplate-go/utils"
	"golang.org/x/crypto/bcrypt"
	"strings"

	log "github.com/sirupsen/logrus"
)

type userUsecase struct {
	userRepository  service.UserRepository
	asyncRepository service.AsynqRepository
}

func NewUserUseCase(userRepository service.UserRepository, asyncRepository service.AsynqRepository) service.UserUsecase {
	return &userUsecase{
		userRepository:  userRepository,
		asyncRepository: asyncRepository,
	}
}

func (u *userUsecase) GetLastLogin(ctx context.Context, args models.GetLastLoginRequest) (result models.GetLastLoginResponse, err error) {
	result, err = u.userRepository.GetLastLoginByUsernameOrEmail(ctx, args)

	return
}

func (u *userUsecase) Register(ctx context.Context, args models.RegisterUserRequest) (result models.RegisterUserResponse, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(args.Password), bcrypt.DefaultCost)
	if err != nil {
		return result, models.ErrorInternalServer
	}

	args.Password = string(hashedPassword)
	args.Username = strings.ToLower(strings.Replace(args.Username, " ", "", -1))

	id, err := u.userRepository.Save(ctx, args)
	if err != nil {
		return result, err
	}

	token, err := utils.GenerateToken(id)
	if err != nil {
		return result, models.ErrorInternalServer
	}

	result = models.RegisterUserResponse{
		Username: args.Username,
		Email:    args.Email,
		Token:    token,
	}

	_, err = u.asyncRepository.Publish(ctx, models.TypeNameTaskSendEmailWelcome, map[string]string{
		"email": args.Email,
	})
	if err != nil {
		log.Error("[UseCase][Register][Publish] %s", err.Error())
	}

	return
}

func (u *userUsecase) Auth(ctx context.Context, args models.LoginUserRequest) (result models.LoginUserResponse, err error) {
	user, err := u.userRepository.GetUserByUsernameOrEmail(ctx, args)
	if err != nil {
		return result, err
	}

	err = utils.VerifyPassword(args.Password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return result, models.ErrorUserWrongPassword
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return result, models.ErrorInternalServer
	}

	result = models.LoginUserResponse{
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return
}
