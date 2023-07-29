package services

import (
	"context"
	"fmt"

	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/repositories"
	"go.uber.org/zap"
)

type BankService interface {
	GetListBank(ctx context.Context) ([]*models.Bank, error)
	GetBankById(ctx context.Context, idBank uint) (*models.Bank, error)
	Create(ctx context.Context, bank *models.Bank) (*models.Bank, error)
}

type bankService struct {
	repository repositories.BankRepository
}

func NewBankService(repository repositories.BankRepository) BankService {
	return &bankService{
		repository: repository,
	}
}

func (e *bankService) Create(ctx context.Context, bank *models.Bank) (*models.Bank, error) {
	res, err := e.repository.Create(ctx, bank)
	if err != nil {
		zap.L().Error("error create bank", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (e *bankService) GetListBank(ctx context.Context) ([]*models.Bank, error) {
	res, err := e.repository.Get(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *bankService) GetBankById(ctx context.Context, idBank uint) (*models.Bank, error) {
	res, err := e.repository.GetById(ctx, idBank)
	if err != nil {
		return nil, err
	}
	if res.ID == 0 {
		zap.L().Error("id bank found")
		return nil, fmt.Errorf("id bank found")
	}
	return res, nil
}
