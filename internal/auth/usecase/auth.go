package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/damndelion/sdu-virtualization-final/config/auth"
	"github.com/damndelion/sdu-virtualization-final/internal/auth/controller/http/v1/dto"
	userEntity "github.com/damndelion/sdu-virtualization-final/internal/user/entity"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	repo AuthRepo
	cfg  *auth.Config
}

func NewAuth(repo AuthRepo, cfg *auth.Config) *Auth {
	return &Auth{repo, cfg}
}

func (u *Auth) Register(ctx context.Context, name, email, password string) error {
	err := u.repo.CheckForEmail(ctx, email)
	if err != nil {
		return err
	}

	_, err = u.repo.CreateUser(ctx, &userEntity.User{
		Name:     name,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *Auth) Login(ctx context.Context, email, password string) (*dto.LoginResponse, error) {

	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("passwords do not match %v", err))
	}
	accessToken, refreshToken, err := u.generateTokens(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Name:         user.Name,
		Email:        user.Email,
		Role:         user.Role,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *Auth) generateTokens(ctx context.Context, user *userEntity.User) (string, string, error) {
	accessTokenClaims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"name":    user.Name,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Duration(u.cfg.AccessTokenTTL) * time.Second).Unix(),
	}

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := access.SignedString([]byte(u.cfg.SecretKey))
	if err != nil {
		return "", "", err
	}

	refreshTokenClaims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Duration(u.cfg.RefreshTokenTTL) * time.Second).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), refreshTokenClaims)

	refreshTokenString, err := refreshToken.SignedString([]byte(u.cfg.SecretKey))
	if err != nil {
		return "", "", err
	}

	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
