package user

import (
	"context"

	"github.com/diuliano-vargas-silveira/academia-api/server/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	FindUserByLogin(ctx context.Context, login string) (models.User, error)
	Create(ctx context.Context, login string, password string) error
}

type RepositoryPostgres struct {
	Conn *pgxpool.Pool
}

func (r *RepositoryPostgres) FindUserByLogin(ctx context.Context,
	login string) (models.User, error) {
	var user models.User
	err := r.Conn.QueryRow(
		ctx,
		"SELECT id, login, senha FROM usuario WHERE login = $1;",
		login,
	).Scan(&user.Id, &user.Login, &user.Password)

	if err == pgx.ErrNoRows {
		return models.User{}, nil
	}

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *RepositoryPostgres) Create(ctx context.Context, login string, password string) error {
	/* NOTE: might add a transaction here */
	_, err := r.Conn.Exec(
		ctx,
		`
			INSERT INTO usuario(login, senha)
			VALUES($1, $2);
		`, login, password)

	if err != nil {
		return err
	}

	return nil
}
