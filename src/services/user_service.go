package services

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/leekchan/accounting"
	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/repositories"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type UserService interface {
	Get(ctx context.Context) (*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Edit(ctx *gin.Context, user *models.User) (*models.User, error)
	GetRiwayatPembayaranUser(ctx *gin.Context) (*models.RiwayatPembayaranResponse, error)
	ExportRiwayatPembayaranUser(ctx *gin.Context) ([]*models.PDF, error)
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
		return nil, err
	}

	res, err := e.repository.Edit(ctx, userTarget, userReq)
	if err != nil {
		zap.L().Error("error create user", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (e *userService) GetRiwayatPembayaranUser(ctx *gin.Context) (*models.RiwayatPembayaranResponse, error) {
	// Get User
	user, err := e.authService.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}

	zakatPenghasilan, zakatTabungan, zakatPerdagangan, zakatEmas, infaqRiwayat, err := e.RiwayatPembayaranUser(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &models.RiwayatPembayaranResponse{ZakatPenghasilan: zakatPenghasilan, ZakatTabungan: zakatTabungan, ZakatPerdagangan: zakatPerdagangan, ZakatEmas: zakatEmas, InfaqRiwayat: infaqRiwayat}, nil
}

func (e *userService) ExportRiwayatPembayaranUser(ctx *gin.Context) ([]*models.PDF, error) {
	// Get User
	user, err := e.authService.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}

	zakatPenghasilan, zakatTabungan, zakatPerdagangan, zakatEmas, infaqRiwayat, err := e.RiwayatPembayaranUser(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	var data []*models.PDF
	ac := accounting.Accounting{Symbol: "Rp. ", Thousand: "."}
	var wg sync.WaitGroup

	wg.Add(5)
	go func() {
		defer wg.Done()
		for _, v := range zakatPenghasilan {
			data = append(data, &models.PDF{
				KodeRiwayat: v.KodeRiwayat,
				Tipe:        "Zakat Penghasilan",
				Bayar:       ac.FormatMoney(v.Bayar),
				Tanggal:     v.Model.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}
	}()
	go func() {
		defer wg.Done()
		for _, v := range zakatEmas {
			data = append(data, &models.PDF{
				KodeRiwayat: v.KodeRiwayat,
				Tipe:        "Zakat Emas",
				Bayar:       ac.FormatMoney(v.Bayar),
				Tanggal:     v.Model.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}
	}()
	go func() {
		defer wg.Done()
		for _, v := range zakatPerdagangan {
			data = append(data, &models.PDF{
				KodeRiwayat: v.KodeRiwayat,
				Tipe:        "Zakat Perdagangan",
				Bayar:       ac.FormatMoney(v.Bayar),
				Tanggal:     v.Model.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}
	}()
	go func() {
		defer wg.Done()
		for _, v := range zakatTabungan {
			data = append(data, &models.PDF{
				KodeRiwayat: v.KodeRiwayat,
				Tipe:        "Zakat Tabungan",
				Bayar:       ac.FormatMoney(v.Bayar),
				Tanggal:     v.Model.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}
	}()
	go func() {
		defer wg.Done()
		for _, v := range infaqRiwayat {
			data = append(data, &models.PDF{
				KodeRiwayat: v.KodeRiwayat,
				Tipe:        v.Judul,
				Bayar:       ac.FormatMoney(v.Nominal),
				Tanggal:     v.Model.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}
	}()
	wg.Wait()

	sort.Slice(data, func(i, j int) bool {
		return data[i].Tanggal > data[j].Tanggal
	})

	return data, nil
}

func (e *userService) RiwayatPembayaranUser(ctx *gin.Context, userId uint) ([]*models.ZakatPenghasilan, []*models.ZakatTabungan, []*models.ZakatPerdagangan, []*models.ZakatEmas, []*models.InfaqRiwayat, error) {
	var zakatPenghasilan []*models.ZakatPenghasilan
	var zakatTabungan []*models.ZakatTabungan
	var zakatPerdagangan []*models.ZakatPerdagangan
	var zakatEmas []*models.ZakatEmas
	var infaqRiwayat []*models.InfaqRiwayat

	var g errgroup.Group

	g.Go(func() error {
		res, err := e.zakatService.GetRiwayatZakatPenghasilanByUserId(ctx, userId)
		if err != nil {
			return fmt.Errorf("failed to get riwayat penghasilan by user id: %w", err)
		}
		zakatPenghasilan = res
		return nil
	})
	g.Go(func() error {
		res, err := e.zakatService.GetRiwayatZakatTabunganByUserId(ctx, userId)
		if err != nil {
			return fmt.Errorf("failed to get riwayat tabungan by user id: %w", err)
		}
		zakatTabungan = res
		return nil
	})
	g.Go(func() error {
		res, err := e.zakatService.GetRiwayatZakatPerdaganganByUserId(ctx, userId)
		if err != nil {
			return fmt.Errorf("failed to get riwayat perdagangan by user id: %w", err)
		}
		zakatPerdagangan = res
		return nil
	})
	g.Go(func() error {
		res, err := e.zakatService.GetRiwayatZakatEmasByUserId(ctx, userId)
		if err != nil {
			return fmt.Errorf("failed to get riwayat emas by user id: %w", err)
		}
		zakatEmas = res
		return nil
	})
	g.Go(func() error {
		res, err := e.infaqService.GetRiwayatInfaqByUserId(ctx, userId)
		if err != nil {
			return fmt.Errorf("failed to get riwayat infaq by user id: %w", err)
		}
		infaqRiwayat = res
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, nil, nil, nil, nil, err
	}

	return zakatPenghasilan, zakatTabungan, zakatPerdagangan, zakatEmas, infaqRiwayat, nil
}
