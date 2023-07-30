package services

import (
	"context"

	"github.com/naufalkhz/zakat/src/gateway"
	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/repositories"
	"go.uber.org/zap"
)

type EmasService interface {
	InquryHargaEmas()
	Get(ctx context.Context) (*models.Emas, error)
}

type emasService struct {
	gateway    gateway.EmasGateway
	repository repositories.EmasRepository
}

func NewEmasService(gateway gateway.EmasGateway, repository repositories.EmasRepository) EmasService {
	return &emasService{
		gateway:    gateway,
		repository: repository,
	}
}

func (e *emasService) Get(ctx context.Context) (*models.Emas, error) {
	res, err := e.repository.Get(ctx)
	if err != nil {
		zap.L().Error("error get emas", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (e *emasService) InquryHargaEmas() {
	emas, err := e.gateway.GetHargaEmas(context.Background())
	if err != nil {
		zap.L().Error("error hit endpoint harga emas", zap.Error(err))
		return
	}

	if _, err = e.repository.Create(context.Background(), emas); err != nil {
		zap.L().Error("error save harga emas", zap.Error(err))
		return
	}
}
