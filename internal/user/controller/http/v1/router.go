package v1

import (
	"github.com/damndelion/sdu-virtualization-final/config/user"
	"github.com/damndelion/sdu-virtualization-final/internal/user/usecase"
	"github.com/damndelion/sdu-virtualization-final/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewUserRouter(handler *gin.Engine, l logger.Interface, u usecase.UserUseCase, cfg *user.Config) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Routers
	h := handler.Group("/v1")
	{
		newUserRoutes(h, u, l, cfg)
	}
}
