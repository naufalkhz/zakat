package repositories

import (
	"context"

	"github.com/naufalkhz/zakat/src/models"
	"gorm.io/gorm"
)

type InfaqRepository interface {
	CreateInfaq(ctx context.Context, infaq *models.Infaq) (*models.Infaq, error)
	GetList(ctx context.Context) ([]*models.Infaq, error)
}

type infaqRepository struct {
	db *gorm.DB
}

func NewInfaqRepository(db *gorm.DB) InfaqRepository {
	return &infaqRepository{db: db}
}

func (r *infaqRepository) CreateInfaq(ctx context.Context, infaq *models.Infaq) (*models.Infaq, error) {
	tx := r.db.WithContext(ctx).Omit("dana_terkumpul").Create(&infaq)
	return infaq, tx.Error
}

func (r *infaqRepository) GetList(ctx context.Context) ([]*models.Infaq, error) {
	var infaq []*models.Infaq
	tx := r.db.WithContext(ctx).Find(&infaq)
	return infaq, tx.Error
}
