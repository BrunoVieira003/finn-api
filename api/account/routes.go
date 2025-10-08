package account

import "github.com/gin-gonic/gin"

func (h *AccountHandler) RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/:userId/accounts")

	routes.POST("", h.CreateAccount)
	routes.GET("", h.GetAccounts)
	routes.GET("/:id", h.GetAccountById)
	routes.DELETE("/:id", h.DeleteAccount)
}