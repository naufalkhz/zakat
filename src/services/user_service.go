package services

import (
	"context"

	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/repositories"
	"go.uber.org/zap"
)

type UserService interface {
	Get(ctx context.Context) (*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (e *userService) Get(ctx context.Context) (*models.User, error) {
	user, err := e.repository.Get(ctx)
	if err != nil {
		zap.L().Error("error get harga emas", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (e *userService) Create(ctx context.Context, user *models.User) (*models.User, error) {

	if err := user.HashPassword(user.Password); err != nil {
		zap.L().Error("hasing password error", zap.Error(err))
		return nil, err
	}

	res, err := e.repository.Create(ctx, user)
	if err != nil {
		zap.L().Error("error create user", zap.Error(err))
		return nil, err
	}
	return res, nil
}
