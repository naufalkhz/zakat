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
