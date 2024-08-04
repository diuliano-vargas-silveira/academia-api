package user

import (
	"context"
	"errors"
	"os"

	"github.com/diuliano-vargas-silveira/academia-api/server/internal/auth"
	"github.com/diuliano-vargas-silveira/academia-api/server/internal/dto/request"
)

type Service struct {
	Repository Repository
}

const NO_ROWS = 0

func (s Service) Login(ctx context.Context, request request.LoginRequest) (string, error) {
	user, err := s.Repository.FindUserByLogin(ctx, request.Login)
	if err != nil {
		return "", err
	}

	if user.Id == NO_ROWS {
		return "", errors.New("user not found")
	}

	isPassCorrect := auth.ComparePasswords(user.Password, []byte(request.Password))
	if !isPassCorrect {
		return "", errors.New("incorrect password")
	}

	secret := []byte(os.Getenv("SECRET"))
	token, err := auth.CreateJWT(secret, user.Id)
	if err != nil {
		return "", errors.New("error when generating token")
	}

	return token, nil
}

func (s Service) Register(ctx context.Context, request request.CreateUserRequest) error {
	hashedPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		return err
	}

	var repositoryErr = s.Repository.Create(ctx, request.Login, hashedPassword)
	if repositoryErr != nil {
		return repositoryErr
	}

	return nil
}
