package main

import (
	"context"
	"log"
	"os"

	"github.com/diuliano-vargas-silveira/academia-api/server/internal/database"
	"github.com/diuliano-vargas-silveira/academia-api/server/internal/middlewares"
	"github.com/diuliano-vargas-silveira/academia-api/server/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	connectionString := os.Getenv("POSTGRES_URI")
	conn, err := database.NewConnection(ctx, connectionString)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	g := gin.Default()

	user.Configure()

	private := g.Group("/api")
	{
		private.Use(middlewares.CheckAuth)
		user.SetPrivateRoutes(private)
	}

	user.SetPublicRoutes(g)

	g.Run(":8080")
}
