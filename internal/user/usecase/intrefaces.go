// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"github.com/damndelion/sdu-virtualization-final/internal/user/controller/http/v1/dto"
	userEntity "github.com/damndelion/sdu-virtualization-final/internal/user/entity"
)

type (

	// UserUseCase -.
	UserUseCase interface {
		Users(ctx context.Context) ([]*userEntity.User, error)
		CreateUser(ctx context.Context, user dto.UserCreateRequest) (int, error)
		UpdateUser(ctx context.Context, userData dto.UserUpdateRequest, email string) error
		GetUserByEmail(ctx context.Context, email string) (*userEntity.User, error)
		GetUserByID(ctx context.Context, id int) (*userEntity.User, error)
		DeleteUser(ctx context.Context, id int) error
	}

	// UserRepo -.
	UserRepo interface {
		GetUsers(ctx context.Context) ([]*userEntity.User, error)
		GetUserByEmail(ctx context.Context, email string) (*userEntity.User, error)
		GetUserByID(ctx context.Context, id int) (*userEntity.User, error)
		CreateUser(ctx context.Context, user dto.UserCreateRequest) (int, error)
		UpdateUser(ctx context.Context, userData dto.UserUpdateRequest, email string) error
		DeleteUser(ctx context.Context, id int) error
	}
)
