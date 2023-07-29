package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/repositories"
	"github.com/naufalkhz/zakat/utils"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type ZakatService interface {
	CreatePenghasilan(ctx *gin.Context, zakatPenghasilan *models.ZakatPenghasilanRequest) (*models.ZakatPenghasilanResponse, error)
}

type zakatService struct {
	repository  repositories.ZakatRepository
	userService UserService
	bankService BankService
	emasService EmasService
}

func NewZakatService(repository repositories.ZakatRepository, userService UserService, bankService BankService, emasService EmasService) ZakatService {
	return &zakatService{
		repository:  repository,
		userService: userService,
		bankService: bankService,
		emasService: emasService,
	}
}

func (e *zakatService) CreatePenghasilan(ctx *gin.Context, zakatPenghasilanReq *models.ZakatPenghasilanRequest) (*models.ZakatPenghasilanResponse, error) {

	///////////////// Bikin concurrency dan di pisah /////////////////
	// Get User
	user, err := e.userService.GetUserSession(ctx)
	if err != nil {
		zap.L().Error("error get user session", zap.Error(err))
		return nil, err
	}
	// Get Bank
	bank, err := e.bankService.GetBankById(ctx, zakatPenghasilanReq.IdBank)
	if err != nil {
		zap.L().Error("error get bank", zap.Error(err))
		return nil, err
	}
	// Get Emas
	emas, err := e.emasService.Get(ctx)
	if err != nil {
		zap.L().Error("error get emas", zap.Error(err))
		return nil, err
	}

	fmt.Println(bank, emas, user)
	///////////////// Bikin concurrency dan di pisah /////////////////

	zakatPenghasilan := &models.ZakatPenghasilan{
		KodeRiwayat:          utils.GenerateCode("PHS"),
		Penghasilan:          zakatPenghasilanReq.Penghasilan,
		PendapatanLain:       zakatPenghasilanReq.PendapatanLain,
		PengeluaranKebutuhan: zakatPenghasilanReq.PengeluaranKebutuhan,
		JenisPenghasilan:     zakatPenghasilanReq.JenisPenghasilan,

		IdUser:    user.ID,
		EmailUser: user.Email,

		IdBank:     bank.ID,
		NamaBank:   bank.Nama,
		AtasNama:   bank.AtasNama,
		NoRekening: bank.NoRekening,
		HargaEmas:  emas.Harga,
		Bayar:      cast.ToFloat64(zakatPenghasilanReq.Penghasilan+zakatPenghasilanReq.PendapatanLain-zakatPenghasilanReq.PengeluaranKebutuhan) * 0.025,
	}

	res, err := e.repository.CreateZakatPenghasilan(ctx, zakatPenghasilan)
	if err != nil {
		zap.L().Error("error create zakatPenghasilan", zap.Error(err))
		return nil, err
	}
	return &models.ZakatPenghasilanResponse{KodeRiwayat: res.KodeRiwayat, Bayar: res.Bayar, Bank: *bank}, nil
}
