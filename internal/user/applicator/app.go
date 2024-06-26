package applicator

import (
	"database/sql"
	"fmt"
	"github.com/damndelion/sdu-virtualization-final/config/user"
	"github.com/damndelion/sdu-virtualization-final/internal/user/controller/grpc"
	v1 "github.com/damndelion/sdu-virtualization-final/internal/user/controller/http/v1"
	userEntity "github.com/damndelion/sdu-virtualization-final/internal/user/entity"
	"github.com/damndelion/sdu-virtualization-final/internal/user/usecase"
	"github.com/damndelion/sdu-virtualization-final/internal/user/usecase/repo"
	"github.com/damndelion/sdu-virtualization-final/pkg/httpserver"
	"github.com/damndelion/sdu-virtualization-final/pkg/logger"
	"github.com/damndelion/sdu-virtualization-final/pkg/postgres"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run(cfg *user.Config) {
	l := logger.New(cfg.Log.Level)

	db, _, err := postgres.New(cfg.PG.URL)
	if err != nil {
		l.Fatal(fmt.Errorf("user - Run - postgres.New: %w", err))
	}
	sqlDB, err := db.DB()
	defer func(sqlDB *sql.DB) {
		err = sqlDB.Close()
		if err != nil {
			l.Fatal(err)
		}
	}(sqlDB)

	userRepo := repo.NewUserRepo(db)
	userUseCase := usecase.NewUser(userRepo, cfg)

	err = db.AutoMigrate(&userEntity.User{})
	if err != nil {
		l.Error(err)
	}

	handler := gin.New()
	v1.NewUserRouter(handler, l, userUseCase, cfg)

	grpcService := grpc.NewService(l, userRepo)
	grpcServer := grpc.NewServer(cfg.GrpcServer.Port, grpcService)
	err = grpcServer.Start()
	if err != nil {
		l.Fatal("failed to start grpc-server err: %v", err)
	}

	defer grpcServer.Close()

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("user - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("user - Run - httpServer.Notify: %w", err))

		err = httpServer.Shutdown()
		if err != nil {
			l.Error(fmt.Errorf("user - Run - httpServer.Shutdown: %w", err))
		}
	}
}
