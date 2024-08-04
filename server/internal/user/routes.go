package user

import (
	"net/http"
	"time"

	"github.com/diuliano-vargas-silveira/academia-api/server/internal/database"
	"github.com/diuliano-vargas-silveira/academia-api/server/internal/dto/request"
	"github.com/diuliano-vargas-silveira/academia-api/server/internal/dto/response"
	"github.com/gin-gonic/gin"
)

var service Service

func Configure() {
	service = Service{
		Repository: &RepositoryPostgres{
			Conn: database.Conn,
		},
	}
}

func SetPrivateRoutes(router *gin.RouterGroup) {
	router.POST("/user", createUser)
}

func SetPublicRoutes(router *gin.Engine) {
	router.POST("/login", handleLogin)
}

func handleLogin(ctx *gin.Context) {
	var requestBody request.LoginRequest

	if err := ctx.BindJSON(&requestBody); err != nil {
		res := response.Response{
			Code:        http.StatusInternalServerError,
			Status:      "Internal server error",
			Data:        nil,
			Error:       "invalid payload",
			RequestedAt: time.Now(),
		}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	if requestBody.Login == "" {
		res := response.Response{
			Code:        http.StatusBadRequest,
			Status:      "Bad Request",
			Data:        nil,
			Error:       "login field must not be empty",
			RequestedAt: time.Now(),
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if requestBody.Password == "" {
		res := response.Response{
			Code:        http.StatusBadRequest,
			Status:      "Bad Request",
			Data:        nil,
			Error:       "password field must not be empty",
			RequestedAt: time.Now(),
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	response := service.Login(ctx, requestBody)

	ctx.JSON(response.Code, response)
}

func createUser(ctx *gin.Context) {
	var requestBody request.CreateUserRequest

	if err := ctx.BindJSON(&requestBody); err != nil {
		res := response.Response{
			Code:        http.StatusInternalServerError,
			Status:      "Internal server error",
			Data:        nil,
			Error:       "invalid payload",
			RequestedAt: time.Now(),
		}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	if requestBody.Login == "" {
		res := response.Response{
			Code:        http.StatusBadRequest,
			Status:      "Bad Request",
			Data:        nil,
			Error:       "login field must not be empty",
			RequestedAt: time.Now(),
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if requestBody.Password == "" {
		res := response.Response{
			Code:        http.StatusBadRequest,
			Status:      "Bad Request",
			Data:        nil,
			Error:       "password field must not be empty",
			RequestedAt: time.Now(),
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	response := service.Register(ctx, requestBody)

	ctx.JSON(response.Code, response)
}
