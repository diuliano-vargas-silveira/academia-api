package user

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/diuliano-vargas-silveira/academia-api/server/internal/auth"
	"github.com/diuliano-vargas-silveira/academia-api/server/internal/dto/request"
	"github.com/diuliano-vargas-silveira/academia-api/server/internal/dto/response"
)

type Service struct {
	Repository Repository
}

const NO_ROWS = 0

func (s Service) Login(ctx context.Context, request request.LoginRequest) response.Response {
	user, err := s.Repository.FindUserByLogin(ctx, request.Login)
	if err != nil {
		response := response.Response{
			Code:        http.StatusInternalServerError,
			Status:      "Internal Server Error",
			Data:        nil,
			Error:       err.Error(),
			RequestedAt: time.Now(),
		}

		return response
	}

	if user.Id == NO_ROWS {
		response := response.Response{
			Code:        http.StatusNotFound,
			Status:      "Not Found",
			Data:        nil,
			Error:       "user not found",
			RequestedAt: time.Now(),
		}

		return response
	}

	isPassCorrect := auth.ComparePasswords(user.Password, []byte(request.Password))
	if !isPassCorrect {
		response := response.Response{
			Code:        http.StatusUnauthorized,
			Status:      "Unauthorized",
			Data:        nil,
			Error:       "incorrect password",
			RequestedAt: time.Now(),
		}

		return response
	}

	secret := []byte(os.Getenv("SECRET"))
	token, err := auth.CreateJWT(secret, user.Id)
	if err != nil {
		response := response.Response{
			Code:        http.StatusInternalServerError,
			Status:      "Internal Server Error",
			Data:        nil,
			Error:       "error when generating token",
			RequestedAt: time.Now(),
		}

		return response
	}

	response := response.Response{
		Code:        http.StatusOK,
		Status:      "OK",
		Data:        token,
		Error:       nil,
		RequestedAt: time.Now(),
	}

	return response
}

func (s Service) Register(ctx context.Context, request request.CreateUserRequest) response.Response {
	hashedPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		response := response.Response{
			Code:        http.StatusInternalServerError,
			Status:      "Internal Server Error",
			Data:        nil,
			Error:       err.Error(),
			RequestedAt: time.Now(),
		}

		return response
	}

	var repositoryErr = s.Repository.Create(ctx, request.Login, hashedPassword)
	if repositoryErr != nil {
		response := response.Response{
			Code:        http.StatusInternalServerError,
			Status:      "Internal Server Error",
			Data:        nil,
			Error:       repositoryErr.Error(),
			RequestedAt: time.Now(),
		}

		return response
	}

	response := response.Response{
		Code:        http.StatusNoContent,
		Status:      "No Content",
		Data:        nil,
		Error:       nil,
		RequestedAt: time.Now(),
	}

	return response
}
