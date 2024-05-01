package v1

import (
	"fmt"
	"github.com/damndelion/sdu-virtualization-final/config/user"
	"github.com/damndelion/sdu-virtualization-final/internal/user/controller/http/middleware"
	"github.com/damndelion/sdu-virtualization-final/internal/user/controller/http/v1/dto"
	"github.com/damndelion/sdu-virtualization-final/internal/user/usecase"
	"github.com/damndelion/sdu-virtualization-final/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userRoutes struct {
	u usecase.UserUseCase
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, u usecase.UserUseCase, l logger.Interface, cfg *user.Config) {
	r := &userRoutes{u, l}

	userHandler := handler.Group("user")
	{
		userHandler.Use(middleware.JwtVerify(cfg.SecretKey))

		userHandler.GET("/", r.GetUser)
		userHandler.GET("/all", r.GetAllUser)
		userHandler.PUT("/:id", r.Update)
		userHandler.DELETE("/:id", r.Delete)

	}
}

func (ur *userRoutes) GetUser(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")

	resUser, err := ur.u.GetUserByID(ctx, userID.(int))
	if err != nil {
		ur.l.Error(fmt.Errorf("http - v1 - user - getUsersById: %w", err))
		errorResponse(ctx, http.StatusInternalServerError, "getUsersById error")

		return
	}

	ctx.JSON(http.StatusOK, resUser)
}

func (ur *userRoutes) GetAllUser(ctx *gin.Context) {

	resUser, err := ur.u.Users(ctx)
	if err != nil {
		ur.l.Error(fmt.Errorf("http - v1 - user - get all user: %w", err))
		errorResponse(ctx, http.StatusInternalServerError, "get all user error")

		return
	}

	ctx.JSON(http.StatusOK, resUser)
}

func (ur *userRoutes) Update(ctx *gin.Context) {
	var client dto.UserUpdateRequest
	if err := ctx.BindJSON(&client); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	idStr := ctx.Param("id")

	err := ur.u.UpdateUser(ctx, client, idStr)
	if err != nil {
		ur.l.Error(fmt.Errorf("http - v1 - user - update user: %w", err))
		errorResponse(ctx, http.StatusInternalServerError, "update  user error")

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (ur *userRoutes) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	err = ur.u.DeleteUser(ctx, id)
	if err != nil {
		ur.l.Error(fmt.Errorf("http - v1 - user - delete user: %w", err))
		errorResponse(ctx, http.StatusInternalServerError, "update  user error")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
