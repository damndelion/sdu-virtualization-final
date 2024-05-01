package v1

import (
	"github.com/damndelion/sdu-virtualization-final/internal/auth/usecase"
	"github.com/damndelion/sdu-virtualization-final/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAuthRouter(handler *gin.Engine, l logger.Interface, u usecase.AuthUseCase) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	h := handler.Group("/v1")
	{
		newAuthRoutes(h, u, l)
	}
}
