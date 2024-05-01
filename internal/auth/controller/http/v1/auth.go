package v1

import (
	"fmt"
	"github.com/damndelion/sdu-virtualization-final/internal/auth/controller/http/v1/dto"
	"github.com/damndelion/sdu-virtualization-final/internal/auth/usecase"
	"github.com/damndelion/sdu-virtualization-final/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authRoutes struct {
	u usecase.AuthUseCase
	l logger.Interface
}

func newAuthRoutes(handler *gin.RouterGroup, u usecase.AuthUseCase, l logger.Interface) {
	r := &authRoutes{u, l}

	userHandler := handler.Group("/auth")
	{
		userHandler.POST("/register", r.Register)
		userHandler.POST("/login", r.Login)
	}
}

func (ar *authRoutes) Register(ctx *gin.Context) {
	var registerRequest dto.RegisterRequest
	err := ctx.ShouldBindJSON(&registerRequest)

	if err != nil {
		ar.l.Error(fmt.Errorf("http - v1 - auth - register: %w", err))
		errorResponse(ctx, http.StatusBadRequest, "Registration form is not correct")

		return
	}

	err = ar.u.Register(ctx, registerRequest.Name, registerRequest.Email, registerRequest.Password)
	if err != nil {
		ar.l.Error(fmt.Errorf("http - v1 - auth - register: %w", err))
		errorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Error: %v", err))

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user successfully registered"})
}

func (ar *authRoutes) Login(ctx *gin.Context) {

	var loginRequest dto.LoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		ar.l.Error(fmt.Errorf("http - v1 - auth - login: %w", err))
		errorResponse(ctx, http.StatusBadRequest, "Login form error")

		return
	}

	token, err := ar.u.Login(ctx, loginRequest.Email, loginRequest.Password)
	if err != nil {
		ar.l.Error(fmt.Errorf("http - v1 - auth - login: %w", err))
		errorResponse(ctx, http.StatusInternalServerError, "Login error")

		return
	}

	ctx.JSON(http.StatusOK, token)
}
