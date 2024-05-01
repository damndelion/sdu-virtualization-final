package repo

import (
	"context"
	"github.com/damndelion/sdu-virtualization-final/internal/user/controller/http/v1/dto"
	userEntity "github.com/damndelion/sdu-virtualization-final/internal/user/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (ur *UserRepo) GetUsers(_ context.Context) (users []*userEntity.User, err error) {
	res := ur.DB.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (ur *UserRepo) CreateUser(ctx context.Context, userRequest dto.UserCreateRequest) (int, error) {
	generatedHash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user := userEntity.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: string(generatedHash),
		Wallet:   userRequest.Wallet,
		Role:     "user",
	}
	res := ur.DB.WithContext(ctx).Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}

	return user.ID, nil
}
