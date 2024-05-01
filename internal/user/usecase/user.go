package usecase

import (
	"context"
	"github.com/damndelion/sdu-virtualization-final/config/user"
	"github.com/damndelion/sdu-virtualization-final/internal/user/controller/http/v1/dto"
	userEntity "github.com/damndelion/sdu-virtualization-final/internal/user/entity"
)

type User struct {
	repo UserRepo
	cfg  *user.Config
}

func NewUser(repo UserRepo, cfg *user.Config) *User {
	return &User{repo, cfg}
}

func (u *User) Users(ctx context.Context) ([]*userEntity.User, error) {
	return u.repo.GetUsers(ctx)
}

func (u *User) CreateUser(ctx context.Context, user dto.UserCreateRequest) (int, error) {
	return u.repo.CreateUser(ctx, user)
}

func (u *User) UpdateUser(ctx context.Context, userData dto.UserUpdateRequest, id string) error {
	err := u.repo.UpdateUser(ctx, userData, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) DeleteUser(ctx context.Context, id int) error {
	err := u.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetUserByEmail(ctx context.Context, email string) (*userEntity.User, error) {
	return u.repo.GetUserByEmail(ctx, email)
}

func (u *User) GetUserByID(ctx context.Context, id int) (*userEntity.User, error) {
	return u.repo.GetUserByID(ctx, id)
}
