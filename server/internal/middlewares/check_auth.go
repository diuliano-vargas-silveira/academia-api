package middlewares

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/diuliano-vargas-silveira/academia-api/server/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CheckAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authToken[1]

	token, err := auth.ValidateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	claims := token.Claims.(jwt.MapClaims)

	if float64(time.Now().Unix()) > claims["expiresAt"].(float64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userIdClaim := claims["userID"].(string)

	userId, err := strconv.Atoi(userIdClaim)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to get user id"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("currentUser", userId)

	c.Next()
}
