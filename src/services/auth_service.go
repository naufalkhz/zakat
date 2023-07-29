package services

import (
	"context"
	"fmt"

	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/repositories"
	"go.uber.org/zap"
)

type AuthService interface {
	Login(ctx context.Context, auth *models.AuthRequest) (*models.User, error)
}

type authService struct {
	repository repositories.AuthRepository
}

func NewAuthService(repository repositories.AuthRepository) AuthService {
	return &authService{
		repository: repository,
	}
}

func (e *authService) Login(ctx context.Context, auth *models.AuthRequest) (*models.User, error) {

	res, err := e.repository.Get(ctx, auth)
	if err != nil {
		zap.L().Error("error get user", zap.Error(err))
		return nil, err
	}

	if err := res.CheckPassword(auth.Password); err != nil {
		zap.L().Error("invalid credential", zap.Error(err))
		return nil, fmt.Errorf("invalid credential")
	}

	return res, nil
}
