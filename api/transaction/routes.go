package transaction

import "github.com/gin-gonic/gin"

func (h *TransactionHandler) RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/transactions")

	routes.POST("", h.CreateTransaction)
	routes.GET("", h.GetTransactions)
}