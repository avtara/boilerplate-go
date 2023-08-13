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
