package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"

	"finnapi/api/account"
	"finnapi/api/transaction"
	"finnapi/db/models"
)

type CreateAccount struct{
	Name string `json:"name" binding:"required"`
}

func main() {
	router := gin.Default()

	ctx := context.Background()
	godotenv.Load()
	
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Cannot connect to database")
	}
	defer conn.Close(ctx)

	queries := models.New(conn)

	accountHandler := account.NewAccountHandler(queries)
	accountHandler.RegisterRoutes(router)

	transactionHandler := transaction.NewTransactionHandler(queries)
	transactionHandler.RegisterRoutes(router)

	router.Run()
}
