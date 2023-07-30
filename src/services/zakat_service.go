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
	CreatePenghasilan(ctx *gin.Context, zakatPenghasilan *models.ZakatPenghasilanRequest) (*models.PembayaranResponse, error)
	CreateTabungan(ctx *gin.Context, zakatTabungan *models.ZakatTabunganRequest) (*models.PembayaranResponse, error)
	CreatePerdagangan(ctx *gin.Context, zakatPerdagangan *models.ZakatPerdaganganRequest) (*models.PembayaranResponse, error)
	CreateEmas(ctx *gin.Context, zakatEmas *models.ZakatEmasRequest) (*models.PembayaranResponse, error)

	GetRiwayatZakatPenghasilanByUserId(ctx *gin.Context, idUser uint) ([]*models.ZakatPenghasilan, error)
	GetRiwayatZakatTabunganByUserId(ctx *gin.Context, idUser uint) ([]*models.ZakatTabungan, error)
	GetRiwayatZakatPerdaganganByUserId(ctx *gin.Context, idUser uint) ([]*models.ZakatPerdagangan, error)
	GetRiwayatZakatEmasByUserId(ctx *gin.Context, idUser uint) ([]*models.ZakatEmas, error)
}

type zakatService struct {
	repository  repositories.ZakatRepository
	authService AuthService
	bankService BankService
	emasService EmasService
}

func NewZakatService(repository repositories.ZakatRepository, authService AuthService, bankService BankService, emasService EmasService) ZakatService {
	return &zakatService{
		repository:  repository,
		authService: authService,
		bankService: bankService,
		emasService: emasService,
	}
}

func (e *zakatService) CreatePenghasilan(ctx *gin.Context, zakatPenghasilanReq *models.ZakatPenghasilanRequest) (*models.PembayaranResponse, error) {

	///////////////// Bikin concurrency dan di pisah /////////////////
	// Get User
	user, err := e.authService.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}
	// Get Bank
	bank, err := e.bankService.GetBankById(ctx, zakatPenghasilanReq.IdBank)
	if err != nil {
		return nil, err
	}
	// Get Emas
	emas, err := e.emasService.Get(ctx)
	if err != nil {
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

		PembayaranInfo: models.PembayaranInfo{
			IdUser:    user.ID,
			NamaUser:  user.Nama,
			EmailUser: user.Email,
			HargaEmas: emas.Harga,
			PembayaranBank: models.PembayaranBank{
				IdBank:     bank.ID,
				Nama:       bank.NamaBank,
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
	return &models.PembayaranResponse{KodeRiwayat: res.KodeRiwayat, Bayar: res.Bayar, Bank: *bank}, nil
}

func (e *zakatService) CreateTabungan(ctx *gin.Context, zakatTabunganReq *models.ZakatTabunganRequest) (*models.PembayaranResponse, error) {

	///////////////// Bikin concurrency dan di pisah /////////////////
	// Get User
	user, err := e.authService.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}
	// Get Bank
	bank, err := e.bankService.GetBankById(ctx, zakatTabunganReq.IdBank)
	if err != nil {
		return nil, err
	}
	// Get Emas
	emas, err := e.emasService.Get(ctx)
	if err != nil {
		return nil, err
	}
	///////////////// Bikin concurrency dan di pisah /////////////////

	zakatTabungan := &models.ZakatTabungan{
		KodeRiwayat:   utils.GenerateCode("TBG"),
		SaldoTabungan: zakatTabunganReq.SaldoTabungan,
		Bunga:         zakatTabunganReq.Bunga,
		Bayar:         cast.ToFloat64(zakatTabunganReq.SaldoTabungan-zakatTabunganReq.Bunga) * 0.025,

		PembayaranInfo: models.PembayaranInfo{
			IdUser:    user.ID,
			NamaUser:  user.Nama,
			EmailUser: user.Email,
			HargaEmas: emas.Harga,
			PembayaranBank: models.PembayaranBank{
				IdBank:     bank.ID,
				Nama:       bank.NamaBank,
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
	return &models.PembayaranResponse{KodeRiwayat: res.KodeRiwayat, Bayar: res.Bayar, Bank: *bank}, nil
}

func (e *zakatService) CreatePerdagangan(ctx *gin.Context, zakatPerdaganganReq *models.ZakatPerdaganganRequest) (*models.PembayaranResponse, error) {

	///////////////// Bikin concurrency dan di pisah /////////////////
	// Get User
	user, err := e.authService.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}
	// Get Bank
	bank, err := e.bankService.GetBankById(ctx, zakatPerdaganganReq.IdBank)
	if err != nil {
		return nil, err
	}
	// Get Emas
	emas, err := e.emasService.Get(ctx)
	if err != nil {
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

		PembayaranInfo: models.PembayaranInfo{
			IdUser:    user.ID,
			NamaUser:  user.Nama,
			EmailUser: user.Email,
			HargaEmas: emas.Harga,
			PembayaranBank: models.PembayaranBank{
				IdBank:     bank.ID,
				Nama:       bank.NamaBank,
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
	return &models.PembayaranResponse{KodeRiwayat: res.KodeRiwayat, Bayar: res.Bayar, Bank: *bank}, nil
}

func (e *zakatService) CreateEmas(ctx *gin.Context, ZakatEmasReq *models.ZakatEmasRequest) (*models.PembayaranResponse, error) {

	///////////////// Bikin concurrency dan di pisah /////////////////
	// Get User
	user, err := e.authService.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}
	// Get Bank
	bank, err := e.bankService.GetBankById(ctx, ZakatEmasReq.IdBank)
	if err != nil {
		return nil, err
	}
	// Get Emas
	emas, err := e.emasService.Get(ctx)
	if err != nil {
		return nil, err
	}
	///////////////// Bikin concurrency dan di pisah /////////////////

	zakatEmas := &models.ZakatEmas{
		KodeRiwayat: utils.GenerateCode("EMS"),
		Emas:        ZakatEmasReq.Emas,
		Bayar:       cast.ToFloat64(ZakatEmasReq.Emas*emas.Harga) * 0.025,

		PembayaranInfo: models.PembayaranInfo{
			IdUser:    user.ID,
			NamaUser:  user.Nama,
			EmailUser: user.Email,
			HargaEmas: emas.Harga,
			PembayaranBank: models.PembayaranBank{
				IdBank:     bank.ID,
				Nama:       bank.NamaBank,
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
	return &models.PembayaranResponse{KodeRiwayat: res.KodeRiwayat, Bayar: res.Bayar, Bank: *bank}, nil
}

func (e *zakatService) GetRiwayatZakatPenghasilanByUserId(ctx *gin.Context, idUser uint) ([]*models.ZakatPenghasilan, error) {
	zakatPenghasilan, err := e.repository.GetRiwayatZakatPenghasilan(ctx, idUser)
	if err != nil {
		zap.L().Error("error get riwayat zakat penghasilan", zap.Error(err))
		return nil, err
	}

	return zakatPenghasilan, nil
}

func (e *zakatService) GetRiwayatZakatTabunganByUserId(ctx *gin.Context, idUser uint) ([]*models.ZakatTabungan, error) {
	zakatTabungan, err := e.repository.GetRiwayatZakatTabungan(ctx, idUser)
	if err != nil {
		zap.L().Error("error get riwayat zakat tabungan", zap.Error(err))
		return nil, err
	}

	return zakatTabungan, nil
}

func (e *zakatService) GetRiwayatZakatPerdaganganByUserId(ctx *gin.Context, idUser uint) ([]*models.ZakatPerdagangan, error) {
	zakatPerdagangan, err := e.repository.GetRiwayatZakatPerdagangan(ctx, idUser)
	if err != nil {
		zap.L().Error("error get riwayat zakat perdagangan", zap.Error(err))
		return nil, err
	}

	return zakatPerdagangan, nil
}

func (e *zakatService) GetRiwayatZakatEmasByUserId(ctx *gin.Context, idUser uint) ([]*models.ZakatEmas, error) {
	zakatEmas, err := e.repository.GetRiwayatZakatEmas(ctx, idUser)
	if err != nil {
		zap.L().Error("error get riwayat zakat emas", zap.Error(err))
		return nil, err
	}

	return zakatEmas, nil
}
