package postgres

import (
	"context"
	"github.com/avtara/boilerplate-go/internal/models"
	"github.com/avtara/boilerplate-go/internal/service"
	"github.com/avtara/boilerplate-go/internal/service/repository/postgres/queries"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(
	conn *sqlx.DB,
) service.UserRepository {
	return &userRepository{
		conn: conn,
	}
}

func (a *userRepository) GetLastLoginByUsernameOrEmail(ctx context.Context, args models.GetLastLoginRequest) (result models.GetLastLoginResponse, err error) {
	err = a.conn.QueryRowContext(
		ctx,
		queries.GetLastLoginByUsernameOrEmail,
		args.Username,
		args.Email,
	).Scan(&result.LastLogin)
	if err != nil {
		return
	}

	return
}

func (a *userRepository) Save(ctx context.Context, args models.RegisterUserRequest) (id int64, err error) {
	err = a.conn.QueryRowContext(ctx, queries.CreateAccount, args.Username, args.Password, args.Email).Scan(&id)
	if err != nil {
		return
	}

	return
}

func (a *userRepository) GetUserByUsernameOrEmail(ctx context.Context, args models.LoginUserRequest) (result models.User, err error) {
	err = a.conn.QueryRowContext(
		ctx,
		queries.GetUserByUsernameOrEmail,
		args.Username,
		args.Email,
	).Scan(&result.ID, &result.Username, &result.Email, &result.Password)
	if err != nil {
		return
	}

	return
}
