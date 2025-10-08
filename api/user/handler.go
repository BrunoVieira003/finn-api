package user

import (
	"finnapi/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	queries *models.Queries
}

func NewUserHandler(queries *models.Queries) *UserHandler{
	return &UserHandler{queries: queries}
}

type NewUser struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) CreateUser(ctx *gin.Context){
	var newUser NewUser

	if err := ctx.BindJSON(&newUser); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
			"message": "Body was not passed correctly",
		})
		return
	}

	user, err := h.queries.CreateUser(ctx, models.CreateUserParams{
		Username: newUser.Username,
		Email: newUser.Email,
		Password: newUser.Password,
	})

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   err.Error(),
			"message": "Something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":   user.ID,
		"username": user.Username,
		"email": user.Email,
		"link": "/user/" + user.ID.String(),
	})
}

func (h *UserHandler) GetUsers(ctx *gin.Context){
	users, err := h.queries.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *UserHandler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := uuid.Parse((id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid id passed. The id should be a valid uuid",
		})
		return
	}
	user, err := h.queries.FindUserById(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := uuid.Parse((id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid id passed. the id should be a valid uuid",
		})
		return
	}

	err = h.queries.DeleteUser(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   err.Error(),
			"message": "Something went wrong",
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}