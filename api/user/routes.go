package user

import "github.com/gin-gonic/gin"

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/users")

	routes.POST("", h.CreateUser)
	routes.GET("", h.GetUsers)
	routes.GET("/:id", h.GetUserById)
	routes.DELETE("/:id", h.DeleteUser)
}