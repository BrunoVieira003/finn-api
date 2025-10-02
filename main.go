package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"finnapi/db/models"
);

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
		if(err != nil){
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"message": "Internal Server Error",
				"error": err.Error(),
			})

			return
		}
		
		ctx.JSON(http.StatusOK, gin.H{
			"accounts": accounts,
		})
	})


	router.Run()
}