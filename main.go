package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

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

	router.GET("/accounts", func(ctx *gin.Context) {
		accounts, err := queries.ListAccounts(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "Internal Server Error",
				"error":   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"accounts": accounts,
		})
	})

	router.POST("/accounts", func(ctx *gin.Context) {
		var newAccount CreateAccount

		if err := ctx.BindJSON(&newAccount); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": err.Error(),
				"message": "Body was not passed correctly",
			})
			return
		}

		if account, err := queries.CreateAccount(ctx, newAccount.Name); err != nil {
			ctx.JSON(http.StatusCreated, gin.H{
				"id":   account.ID,
				"name": account.Name,
				"link": "/accounts/" + account.ID.String(),
			})
			return
		}
	})

	router.GET("/accounts/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		accountId, err := uuid.Parse((id))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "Invalid id passed. the id should be a valid uuid",
			})
			return
		}
		account, err := queries.FindAccountById(ctx, accountId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Account not found",
			})
			return
		}

		ctx.JSON(http.StatusOK, account)
	})

	router.Run()
}
