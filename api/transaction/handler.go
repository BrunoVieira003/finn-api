package transaction

import (
	"finnapi/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionHandler struct{
	queries *models.Queries
}

func NewTransactionHandler(queries *models.Queries) *TransactionHandler {
	return &TransactionHandler{queries: queries}
}

func (h *TransactionHandler) CreateTransaction(ctx *gin.Context){
	var newTransaction models.CreateTransactionParams

	if err := ctx.BindJSON(&newTransaction); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
			"message": "Body was not passed correctly",
		})
		return
	}

	transaction, err := h.queries.CreateTransaction(ctx, newTransaction)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   err.Error(),
			"message": "Something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":   transaction.ID,
		"amount": transaction.Amount,
		"type": transaction.Type,
		"date": transaction.Date,
		"account": transaction.AccountID,
		"description": transaction.Description,
		"link": "/transactions/" + transaction.ID.String(),
	})
}

func (h *TransactionHandler) GetTransactions(ctx *gin.Context){
	transactions, err := h.queries.ListTransactions(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}

func (h *TransactionHandler) GetTransactionById(ctx *gin.Context){
	id := ctx.Param("id")
	transactionId, err := uuid.Parse((id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid id passed. the id should be a valid uuid",
		})
		return
	}

	transaction, err := h.queries.FindTransactionById(ctx, transactionId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Transaction not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}