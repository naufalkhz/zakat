package services

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/repositories"
	"go.uber.org/zap"
)

type UserService interface {
	Get(ctx context.Context) (*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Edit(ctx *gin.Context, user *models.User) (*models.User, error)
	GetRiwayatUser(ctx *gin.Context) (*models.RiwayatTransaksiResponse, error)
}

type userService struct {
	repository   repositories.UserRepository
	authService  AuthService
	zakatService ZakatService
	infaqService InfaqService
}

func NewUserService(repository repositories.UserRepository, authService AuthService, zakatService ZakatService, infaqService InfaqService) UserService {
	return &userService{
		repository:   repository,
		authService:  authService,
		zakatService: zakatService,
		infaqService: infaqService,
	}
}

func (e *userService) Get(ctx context.Context) (*models.User, error) {
	user, err := e.repository.Get(ctx)
	if err != nil {
		zap.L().Error("error user", zap.Error(err))
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

func (e *userService) Edit(ctx *gin.Context, userReq *models.User) (*models.User, error) {
	userTarget, err := e.authService.GetUserSession(ctx)
	if err != nil {
		zap.L().Error("failed toget user", zap.Error(err))
		return nil, err
	}

	res, err := e.repository.Edit(ctx, userTarget, userReq)
	if err != nil {
		zap.L().Error("error create user", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (e *userService) GetRiwayatUser(ctx *gin.Context) (*models.RiwayatTransaksiResponse, error) {
	// Get User
	user, err := e.authService.GetUserSession(ctx)
	if err != nil {
		zap.L().Error("error get user session", zap.Error(err))
		return nil, err
	}

	///////////// TODO: Buat ini parallel //////////////////
	zakatPenghasilan, err := e.zakatService.GetRiwayatZakatPenghasilanByUserId(ctx, user.ID)
	if err != nil {
		zap.L().Error("error get riwayat zakat penghasilan", zap.Error(err))
		return nil, err
	}

	zakatTabungan, err := e.zakatService.GetRiwayatZakatTabunganByUserId(ctx, user.ID)
	if err != nil {
		zap.L().Error("error get riwayat zakat tabungan", zap.Error(err))
		return nil, err
	}

	zakatPerdagangan, err := e.zakatService.GetRiwayatZakatPerdaganganByUserId(ctx, user.ID)
	if err != nil {
		zap.L().Error("error get riwayat zakat perdagangan", zap.Error(err))
		return nil, err
	}

	zakatEmas, err := e.zakatService.GetRiwayatZakatEmasByUserId(ctx, user.ID)
	if err != nil {
		zap.L().Error("error get riwayat zakat perdagangan", zap.Error(err))
		return nil, err
	}

	infaqRiwayat, err := e.infaqService.GetRiwayatInfaqByUserId(ctx, user.ID)
	if err != nil {
		zap.L().Error("error get riwayat zakat perdagangan", zap.Error(err))
		return nil, err
	}

	///////////// TODO: Buat ini parallel //////////////////

	return &models.RiwayatTransaksiResponse{ZakatPenghasilan: zakatPenghasilan, ZakatTabungan: zakatTabungan, ZakatPerdagangan: zakatPerdagangan, ZakatEmas: zakatEmas, InfaqRiwayat: infaqRiwayat}, nil
}
