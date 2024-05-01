package usecase

import (
	"context"
	"github.com/damndelion/sdu-virtualization-final/internal/auth/controller/http/v1/dto"
	userEntity "github.com/damndelion/sdu-virtualization-final/internal/user/entity"
)

type (

	// AuthUseCase -.
	AuthUseCase interface {
		Register(ctx context.Context, name, email, password string) error
		Login(ctx context.Context, email, password string) (*dto.LoginResponse, error)
	}

	// AuthRepo -.
	AuthRepo interface {
		CreateUser(ctx context.Context, user *userEntity.User) (int, error)
		GetUserByEmail(ctx context.Context, email string) (*userEntity.User, error)
		CheckForEmail(ctx context.Context, email string) error
	}
)
