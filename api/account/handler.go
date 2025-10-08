package account

import (
	"finnapi/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AccountHandler struct {
	queries *models.Queries
}

func NewAccountHandler(queries *models.Queries) *AccountHandler{
	return &AccountHandler{queries: queries}
}

func (h *AccountHandler) GetAccounts(ctx *gin.Context){
	accounts, err := h.queries.ListAccounts(ctx)
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
}

type CreateAccount struct{
	Name string `json:"name" binding:"required"`
}

func (h *AccountHandler) CreateAccount(ctx *gin.Context){
	var newAccount CreateAccount

	if err := ctx.BindJSON(&newAccount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"message": "Body was not passed correctly",
		})
		return
	}

	account, err := h.queries.CreateAccount(ctx, newAccount.Name)
	if(err != nil){
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"message": "Something went wrong",
		})
		return
	}
	
	ctx.JSON(http.StatusCreated, gin.H{
		"id":   account.ID,
		"name": account.Name,
		"link": "/accounts/" + account.ID.String(),
	})
}

func (h *AccountHandler) GetAccountById(ctx *gin.Context){
	id := ctx.Param("id")
	accountId, err := uuid.Parse((id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid id passed. the id should be a valid uuid",
		})
		return
	}
	account, err := h.queries.FindAccountById(ctx, accountId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Account not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, account)
}