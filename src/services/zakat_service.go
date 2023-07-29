package services

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/repositories"
	"github.com/naufalkhz/zakat/utils"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type ZakatService interface {
	CreatePenghasilan(ctx *gin.Context, zakatPenghasilan *models.ZakatPenghasilanRequest) (*models.TransaksiResponse, error)
	CreateTabungan(ctx *gin.Context, zakatTabungan *models.ZakatTabunganRequest) (*models.TransaksiResponse, error)
	CreatePerdagangan(ctx *gin.Context, zakatPerdagangan *models.ZakatPerdaganganRequest) (*models.TransaksiResponse, error)
	CreateEmas(ctx *gin.Context, zakatEmas *models.ZakatEmasRequest) (*models.TransaksiResponse, error)
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

func (e *zakatService) CreatePenghasilan(ctx *gin.Context, zakatPenghasilanReq *models.ZakatPenghasilanRequest) (*models.TransaksiResponse, error) {

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
	///////////////// Bikin concurrency dan di pisah /////////////////

	zakatPenghasilan := &models.ZakatPenghasilan{
		KodeRiwayat:          utils.GenerateCode("PHS"),
		Penghasilan:          zakatPenghasilanReq.Penghasilan,
		PendapatanLain:       zakatPenghasilanReq.PendapatanLain,
		PengeluaranKebutuhan: zakatPenghasilanReq.PengeluaranKebutuhan,
		JenisPenghasilan:     zakatPenghasilanReq.JenisPenghasilan,
		Bayar:                cast.ToFloat64(zakatPenghasilanReq.Penghasilan+zakatPenghasilanReq.PendapatanLain-zakatPenghasilanReq.PengeluaranKebutuhan) * 0.025,

		TransaksiInfo: models.TransaksiInfo{
			IdUser:    user.ID,
			EmailUser: user.Email,
			HargaEmas: emas.Harga,
			TransaksiBank: models.TransaksiBank{
				IdBank:     bank.ID,
				Nama:       bank.Nama,
				AtasNama:   bank.AtasNama,
				NoRekening: bank.NoRekening,
			},
		},
	}

	res, err := e.repository.CreateZakatPenghasilan(ctx, zakatPenghasilan)
	if err != nil {
		zap.L().Error("error create zakat penghasilan", zap.Error(err))
		return nil, err
	}
	return &models.TransaksiResponse{KodeRiwayat: res.KodeRiwayat, Bayar: res.Bayar, Bank: *bank}, nil
}

func (e *zakatService) CreateTabungan(ctx *gin.Context, zakatTabunganReq *models.ZakatTabunganRequest) (*models.TransaksiResponse, error) {

	///////////////// Bikin concurrency dan di pisah /////////////////
	// Get User
	user, err := e.userService.GetUserSession(ctx)
	if err != nil {
		zap.L().Error("error get user session", zap.Error(err))
		return nil, err
	}
	// Get Bank
	bank, err := e.bankService.GetBankById(ctx, zakatTabunganReq.IdBank)
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
	///////////////// Bikin concurrency dan di pisah /////////////////

	zakatTabungan := &models.ZakatTabungan{
		KodeRiwayat:   utils.GenerateCode("TBG"),
		SaldoTabungan: zakatTabunganReq.SaldoTabungan,
		Bunga:         zakatTabunganReq.Bunga,
		Bayar:         cast.ToFloat64(zakatTabunganReq.SaldoTabungan-zakatTabunganReq.Bunga) * 0.025,

		TransaksiInfo: models.TransaksiInfo{
			IdUser:    user.ID,
			EmailUser: user.Email,
			HargaEmas: emas.Harga,
			TransaksiBank: models.TransaksiBank{
				IdBank:     bank.ID,
				Nama:       bank.Nama,
				AtasNama:   bank.AtasNama,
				NoRekening: bank.NoRekening,
			},
		},
	}

	res, err := e.repository.CreateZakatTabungan(ctx, zakatTabungan)
	if err != nil {
		zap.L().Error("error create zakat tabungan", zap.Error(err))
		return nil, err
	}
	return &models.TransaksiResponse{KodeRiwayat: res.KodeRiwayat, Bayar: res.Bayar, Bank: *bank}, nil
}

func (e *zakatService) CreatePerdagangan(ctx *gin.Context, zakatPerdaganganReq *models.ZakatPerdaganganRequest) (*models.TransaksiResponse, error) {

	///////////////// Bikin concurrency dan di pisah /////////////////
	// Get User
	user, err := e.userService.GetUserSession(ctx)
	if err != nil {
		zap.L().Error("error get user session", zap.Error(err))
		return nil, err
	}
	// Get Bank
	bank, err := e.bankService.GetBankById(ctx, zakatPerdaganganReq.IdBank)
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
	///////////////// Bikin concurrency dan di pisah /////////////////

	zakatPerdagangan := &models.ZakatPerdagangan{
		KodeRiwayat: utils.GenerateCode("PDG"),
		Modal:       zakatPerdaganganReq.Modal,
		Keuntungan:  zakatPerdaganganReq.Keuntungan,
		Piutang:     zakatPerdaganganReq.Piutang,
		Utang:       zakatPerdaganganReq.Utang,
		Kerugian:    zakatPerdaganganReq.Kerugian,

		Bayar: cast.ToFloat64(zakatPerdaganganReq.Modal+zakatPerdaganganReq.Keuntungan+zakatPerdaganganReq.Piutang-zakatPerdaganganReq.Utang-zakatPerdaganganReq.Kerugian) * 0.025,

		TransaksiInfo: models.TransaksiInfo{
			IdUser:    user.ID,
			EmailUser: user.Email,
			HargaEmas: emas.Harga,
			TransaksiBank: models.TransaksiBank{
				IdBank:     bank.ID,
				Nama:       bank.Nama,
				AtasNama:   bank.AtasNama,
				NoRekening: bank.NoRekening,
			},
		},
	}

	res, err := e.repository.CreateZakatPerdagangan(ctx, zakatPerdagangan)
	if err != nil {
		zap.L().Error("error create zakat perdagangan", zap.Error(err))
		return nil, err
	}
	return &models.TransaksiResponse{KodeRiwayat: res.KodeRiwayat, Bayar: res.Bayar, Bank: *bank}, nil
}

func (e *zakatService) CreateEmas(ctx *gin.Context, ZakatEmasReq *models.ZakatEmasRequest) (*models.TransaksiResponse, error) {

	///////////////// Bikin concurrency dan di pisah /////////////////
	// Get User
	user, err := e.userService.GetUserSession(ctx)
	if err != nil {
		zap.L().Error("error get user session", zap.Error(err))
		return nil, err
	}
	// Get Bank
	bank, err := e.bankService.GetBankById(ctx, ZakatEmasReq.IdBank)
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
	///////////////// Bikin concurrency dan di pisah /////////////////

	zakatEmas := &models.ZakatEmas{
		KodeRiwayat: utils.GenerateCode("EMS"),
		Emas:        ZakatEmasReq.Emas,
		Bayar:       cast.ToFloat64(ZakatEmasReq.Emas*emas.Harga) * 0.025,

		TransaksiInfo: models.TransaksiInfo{
			IdUser:    user.ID,
			EmailUser: user.Email,
			HargaEmas: emas.Harga,
			TransaksiBank: models.TransaksiBank{
				IdBank:     bank.ID,
				Nama:       bank.Nama,
				AtasNama:   bank.AtasNama,
				NoRekening: bank.NoRekening,
			},
		},
	}

	res, err := e.repository.CreateZakatEmas(ctx, zakatEmas)
	if err != nil {
		zap.L().Error("error create zakat emas", zap.Error(err))
		return nil, err
	}
	return &models.TransaksiResponse{KodeRiwayat: res.KodeRiwayat, Bayar: res.Bayar, Bank: *bank}, nil
}
