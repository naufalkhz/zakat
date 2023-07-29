package repositories

import (
	"context"

	"github.com/naufalkhz/zakat/src/models"
	"gorm.io/gorm"
)

type ZakatRepository interface {
	CreateZakatPenghasilan(ctx context.Context, zakatPenghasilan *models.ZakatPenghasilan) (*models.ZakatPenghasilan, error)
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
