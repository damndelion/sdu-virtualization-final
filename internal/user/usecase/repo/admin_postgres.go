package repo

import (
	"context"
	"github.com/damndelion/sdu-virtualization-final/internal/user/controller/http/v1/dto"
	userEntity "github.com/damndelion/sdu-virtualization-final/internal/user/entity"

	"golang.org/x/crypto/bcrypt"
)

func (ur *UserRepo) GetUserByEmail(ctx context.Context, email string) (user *userEntity.User, err error) {
	res := ur.DB.Where("email = ?", email).WithContext(ctx).Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (ur *UserRepo) GetUserByID(ctx context.Context, id int) (user *userEntity.User, err error) {
	res := ur.DB.Where("id = ?", id).WithContext(ctx).Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (ur *UserRepo) DeleteUser(ctx context.Context, id int) error {
	err := ur.DB.Where("id = ?", id).Delete(&userEntity.User{}).WithContext(ctx).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepo) UpdateUser(_ context.Context, userData dto.UserUpdateRequest, email string) error {
	generatedHash, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := userEntity.User{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: string(generatedHash),
		Wallet:   userData.Wallet,
		Valid:    userData.Valid,
		Role:     userData.Role,
	}

	err = ur.DB.Model(&user).Where("email = ?", email).Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}
