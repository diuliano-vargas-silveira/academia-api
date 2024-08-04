package user

import (
	"net/http"

	"github.com/diuliano-vargas-silveira/academia-api/server/internal/database"
	"github.com/diuliano-vargas-silveira/academia-api/server/internal/dto/request"
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid payload",
		})
		return
	}

	if requestBody.Login == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "login field must not be empty",
		})
		return
	}

	if requestBody.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "password field must not be empty",
		})
		return
	}

	token, err := service.Login(ctx, requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{"token": token})
}

func createUser(ctx *gin.Context) {
	var requestBody request.CreateUserRequest

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid payload",
		})
		return
	}

	if requestBody.Login == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "login field must not be empty",
		})
		return
	}

	if requestBody.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "password field must not be empty",
		})
		return
	}

	err := service.Register(ctx, requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}
