package account

import "github.com/gin-gonic/gin"

func (h *AccountHandler) RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/accounts")

	routes.POST("", h.CreateAccount)
	routes.GET("", h.GetAccounts)
}