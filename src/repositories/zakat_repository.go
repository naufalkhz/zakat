package repositories

import (
	"context"

	"github.com/naufalkhz/zakat/src/models"
	"gorm.io/gorm"
)

type ZakatRepository interface {
	CreateZakatPenghasilan(ctx context.Context, zakatPenghasilan *models.ZakatPenghasilan) (*models.ZakatPenghasilan, error)
	CreateZakatTabungan(ctx context.Context, zakatTabungan *models.ZakatTabungan) (*models.ZakatTabungan, error)
	CreateZakatPerdagangan(ctx context.Context, zakatPerdagangan *models.ZakatPerdagangan) (*models.ZakatPerdagangan, error)
	CreateZakatEmas(ctx context.Context, zakatEmas *models.ZakatEmas) (*models.ZakatEmas, error)

	GetRiwayatZakatPenghasilan(ctx context.Context, idUser uint) ([]*models.ZakatPenghasilan, error)
	GetRiwayatZakatTabungan(ctx context.Context, idUser uint) ([]*models.ZakatTabungan, error)
	GetRiwayatZakatPerdagangan(ctx context.Context, idUser uint) ([]*models.ZakatPerdagangan, error)
	GetRiwayatZakatEmas(ctx context.Context, idUser uint) ([]*models.ZakatEmas, error)
}

type zakatRepository struct {
	db *gorm.DB
}

func NewZakatRepository(db *gorm.DB) ZakatRepository {
	return &zakatRepository{db: db}
}

func (r *zakatRepository) CreateZakatPenghasilan(ctx context.Context, zakatPenghasilan *models.ZakatPenghasilan) (*models.ZakatPenghasilan, error) {
	tx := r.db.WithContext(ctx).Create(&zakatPenghasilan)
	return zakatPenghasilan, tx.Error
}

func (r *zakatRepository) CreateZakatTabungan(ctx context.Context, zakatTabungan *models.ZakatTabungan) (*models.ZakatTabungan, error) {
	tx := r.db.WithContext(ctx).Create(&zakatTabungan)
	return zakatTabungan, tx.Error
}

func (r *zakatRepository) CreateZakatPerdagangan(ctx context.Context, zakatPerdagangan *models.ZakatPerdagangan) (*models.ZakatPerdagangan, error) {
	tx := r.db.WithContext(ctx).Create(&zakatPerdagangan)
	return zakatPerdagangan, tx.Error
}

func (r *zakatRepository) CreateZakatEmas(ctx context.Context, zakatEmas *models.ZakatEmas) (*models.ZakatEmas, error) {
	tx := r.db.WithContext(ctx).Create(&zakatEmas)
	return zakatEmas, tx.Error
}

func (r *zakatRepository) GetRiwayatZakatPenghasilan(ctx context.Context, idUser uint) ([]*models.ZakatPenghasilan, error) {
	var zakatPenghasilan []*models.ZakatPenghasilan
	tx := r.db.Debug().WithContext(ctx).Where("id_user = ?", idUser).Find(&zakatPenghasilan)
	return zakatPenghasilan, tx.Error
}

func (r *zakatRepository) GetRiwayatZakatTabungan(ctx context.Context, idUser uint) ([]*models.ZakatTabungan, error) {
	var zakatTabungan []*models.ZakatTabungan
	tx := r.db.Debug().WithContext(ctx).Where("id_user = ?", idUser).Find(&zakatTabungan)
	return zakatTabungan, tx.Error
}

func (r *zakatRepository) GetRiwayatZakatPerdagangan(ctx context.Context, idUser uint) ([]*models.ZakatPerdagangan, error) {
	var zakatPerdagangan []*models.ZakatPerdagangan
	tx := r.db.Debug().WithContext(ctx).Where("id_user = ?", idUser).Find(&zakatPerdagangan)
	return zakatPerdagangan, tx.Error
}

func (r *zakatRepository) GetRiwayatZakatEmas(ctx context.Context, idUser uint) ([]*models.ZakatEmas, error) {
	var zakatEmas []*models.ZakatEmas
	tx := r.db.Debug().WithContext(ctx).Where("id_user = ?", idUser).Find(&zakatEmas)
	return zakatEmas, tx.Error
}
