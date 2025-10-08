package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"finnapi/api/account"
	"finnapi/api/user"
	"finnapi/db/models"
)

type CreateAccount struct{
	Name string `json:"name" binding:"required"`
}

func main() {
	router := gin.Default()

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgresql://postgres:postgres@localhost:5432/finn")
	if err != nil {
		panic("Cannot connect to database")
	}
	defer conn.Close(ctx)

	queries := models.New(conn)

	accountHandler := account.NewAccountHandler(queries)
	accountHandler.RegisterRoutes(router)

	userHandler := user.NewUserHandler(queries)
	userHandler.RegisterRoutes(router)

	router.Run()
}
