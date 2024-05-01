package grpc

import (
	"context"
	"fmt"
	"github.com/damndelion/sdu-virtualization-final/internal/user/controller/http/v1/dto"
	"github.com/damndelion/sdu-virtualization-final/internal/user/usecase/repo"
	"github.com/damndelion/sdu-virtualization-final/pkg/logger"
	pb "github.com/damndelion/sdu-virtualization-final/pkg/protobuf/userService/gw"
	"strconv"
)

type Service struct {
	pb.UnimplementedUserServiceServer
	logger *logger.Logger
	repo   *repo.UserRepo
}

func NewService(logger *logger.Logger, repo *repo.UserRepo) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
	}
}

func (s *Service) GetUserByID(ctx context.Context, request *pb.GetUserByIDRequest) (*pb.User, error) {
	id, err := strconv.Atoi(request.Id)
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to GetUserByID err: %v", err)

		return nil, fmt.Errorf("GetUserById err: %w", err)
	}

	return &pb.User{
		Id:       int32(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Wallet:   user.Wallet,
		Valid:    user.Valid,
		Role:     user.Role,
	}, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, request *pb.GetUserByEmailRequest) (*pb.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, request.Email)
	if err != nil {
		s.logger.Error("failed to GetUserByEmail err: %v", err)

		return nil, fmt.Errorf("GetUserByEmail err: %w", err)
	}

	return &pb.User{
		Id:       int32(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Wallet:   user.Wallet,
		Valid:    user.Valid,
		Role:     user.Role,
	}, nil
}

func (s *Service) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := request.GetUser()

	newUser := dto.UserCreateRequest{
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Wallet:   user.GetWallet(),
		Valid:    user.GetValid(),
	}

	id, err := s.repo.CreateUser(ctx, newUser)
	if err != nil {
		s.logger.Error("failed to CreateUser err: %v", err)

		return nil, fmt.Errorf("CreateUser err: %w", err)
	}

	return &pb.CreateUserResponse{Id: int32(id)}, nil
}
