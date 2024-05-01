package repo

import (
	"context"
	"fmt"
	"github.com/damndelion/sdu-virtualization-final/internal/auth/transport"
	userEntity "github.com/damndelion/sdu-virtualization-final/internal/user/entity"
)

type AuthRepo struct {
	userGrpcTransport *transport.UserGrpcTransport
}

func NewAuthRepo(userGrpcTransport *transport.UserGrpcTransport) *AuthRepo {
	return &AuthRepo{userGrpcTransport}
}

func (t *AuthRepo) CreateUser(ctx context.Context, user *userEntity.User) (int, error) {
	grpcUser, err := t.userGrpcTransport.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return int(grpcUser.Id), nil
}

func (t *AuthRepo) GetUserByEmail(ctx context.Context, email string) (*userEntity.User, error) {
	grpcUser, err := t.userGrpcTransport.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	user := &userEntity.User{
		ID:       int(grpcUser.Id),
		Name:     grpcUser.Name,
		Email:    grpcUser.Email,
		Password: grpcUser.Password,
		Wallet:   grpcUser.Wallet,
		Valid:    grpcUser.Valid,
		Role:     grpcUser.Role,
	}

	return user, nil
}

func (t *AuthRepo) CheckForEmail(ctx context.Context, email string) error {
	grpcUser, _ := t.userGrpcTransport.GetUserByEmail(ctx, email)
	if grpcUser.Id != 0 {
		return fmt.Errorf("user with this email alraedy exists")
	}

	return nil
}
