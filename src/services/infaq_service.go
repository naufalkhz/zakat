package services

import (
	"context"

	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/repositories"
	"go.uber.org/zap"
)

type InfaqService interface {
	CreateInfaq(ctx context.Context, zakatPenghasilan *models.Infaq) (*models.Infaq, error)
	GetList(ctx context.Context) ([]*models.Infaq, error)
}

type infaqService struct {
	repository  repositories.InfaqRepository
	userService UserService
	bankService BankService
}

func NewInfaqService(repository repositories.InfaqRepository, userService UserService, bankService BankService) InfaqService {
	return &infaqService{
		repository:  repository,
		userService: userService,
		bankService: bankService,
	}
}

func (e *infaqService) CreateInfaq(ctx context.Context, infaq *models.Infaq) (*models.Infaq, error) {
	res, err := e.repository.CreateInfaq(ctx, infaq)
	if err != nil {
		zap.L().Error("error create infaq", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (e *infaqService) GetList(ctx context.Context) ([]*models.Infaq, error) {
	res, err := e.repository.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
