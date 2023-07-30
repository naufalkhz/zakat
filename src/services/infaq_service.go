package services

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/repositories"
	"github.com/naufalkhz/zakat/utils"
	"go.uber.org/zap"
)

type InfaqService interface {
	CreateInfaq(ctx context.Context, infaq *models.Infaq) (*models.Infaq, error)
	GetList(ctx context.Context) ([]*models.Infaq, error)
	CreateInfaqRiwayat(ctx *gin.Context, infaqRiwayatRequest *models.InfaqRiwayatRequest) (*models.TransaksiResponse, error)
	GetRiwayatInfaqByUserId(ctx *gin.Context, idUser uint) ([]*models.InfaqRiwayat, error)
}

type infaqService struct {
	repository  repositories.InfaqRepository
	authService AuthService
	bankService BankService
}

func NewInfaqService(repository repositories.InfaqRepository, authService AuthService, bankService BankService) InfaqService {
	return &infaqService{
		repository:  repository,
		authService: authService,
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

func (e *infaqService) CreateInfaqRiwayat(ctx *gin.Context, infaqRiwayatRequest *models.InfaqRiwayatRequest) (*models.TransaksiResponse, error) {

	///////////////// Bikin concurrency dan di pisah /////////////////
	// Get User
	user, err := e.authService.GetUserSession(ctx)
	if err != nil {
		zap.L().Error("error get user session", zap.Error(err))
		return nil, err
	}
	// Get Bank
	bank, err := e.bankService.GetBankById(ctx, infaqRiwayatRequest.IdBank)
	if err != nil {
		zap.L().Error("error get bank", zap.Error(err))
		return nil, err
	}

	// Get Infaq
	infaq, err := e.repository.GetById(ctx, infaqRiwayatRequest.IdInfaq)
	if err != nil {
		zap.L().Error("error get infaq", zap.Error(err))
		return nil, err
	}
	///////////////// Bikin concurrency dan di pisah /////////////////

	infaqRiwayat := &models.InfaqRiwayat{
		KodeRiwayat: utils.GenerateCode("IFQ"),
		Nominal:     infaqRiwayatRequest.Nominal,
		Catatan:     infaqRiwayatRequest.Catatan,
		HambaAllah:  infaqRiwayatRequest.HambaAllah,

		IdUser:   user.ID,
		NamaUser: user.Nama,
		Email:    user.Email,

		IdInfaq: infaq.ID,
		Judul:   infaq.Judul,

		TransaksiBank: models.TransaksiBank{
			IdBank:     bank.ID,
			Nama:       bank.NamaBank,
			NoRekening: bank.NoRekening,
			AtasNama:   bank.AtasNama,
		},
	}

	res, err := e.repository.CreateInfaqRiwayat(ctx, infaqRiwayat)
	if err != nil {
		zap.L().Error("error create infaq riwayat", zap.Error(err))
		return nil, err
	}

	// Should handle this when update is failed
	if err := e.repository.UpdateDanaInfaq(ctx, infaqRiwayatRequest.Nominal, infaq); err != nil {
		zap.L().Error("error update dana infaq", zap.Error(err))
		return nil, err
	}

	return &models.TransaksiResponse{KodeRiwayat: res.KodeRiwayat, Bayar: float64(res.Nominal), Bank: *bank}, nil
}

func (e *infaqService) GetRiwayatInfaqByUserId(ctx *gin.Context, idUser uint) ([]*models.InfaqRiwayat, error) {
	infaqRiwayat, err := e.repository.GetInfaqRiwayat(ctx, idUser)
	if err != nil {
		zap.L().Error("error get infaq riwayat", zap.Error(err))
		return nil, err
	}

	return infaqRiwayat, nil
}
