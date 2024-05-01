package applicator

import (
	"fmt"
	"github.com/damndelion/sdu-virtualization-final/config/auth"
	v1 "github.com/damndelion/sdu-virtualization-final/internal/auth/controller/http/v1"
	"github.com/damndelion/sdu-virtualization-final/internal/auth/transport"
	"github.com/damndelion/sdu-virtualization-final/internal/auth/usecase"
	"github.com/damndelion/sdu-virtualization-final/internal/auth/usecase/repo"
	"github.com/damndelion/sdu-virtualization-final/pkg/httpserver"
	"github.com/damndelion/sdu-virtualization-final/pkg/logger"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

// Run creates objects via constructors.
func Run(cfg *auth.Config) {
	l := logger.New(cfg.Log.Level)

	userGrpcTransport := transport.NewUserGrpcTransport(cfg.Transport.UserGrpc)

	authUseCase := usecase.NewAuth(repo.NewAuthRepo(userGrpcTransport), cfg)

	handler := gin.New()
	v1.NewAuthRouter(handler, l, authUseCase)

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("auth - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("auth - Run - httpServer.Notify: %w", err))

		err = httpServer.Shutdown()
		if err != nil {
			l.Error(fmt.Errorf("auth - Run - httpServer.Shutdown: %w", err))
		}
	}
}
