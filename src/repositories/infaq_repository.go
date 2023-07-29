package repositories

import (
	"context"

	"github.com/naufalkhz/zakat/src/models"
	"gorm.io/gorm"
)

type InfaqRepository interface {
	CreateInfaq(ctx context.Context, infaq *models.Infaq) (*models.Infaq, error)
	UpdateDanaInfaq(ctx context.Context, nominal int64, infaq *models.Infaq) error
	GetList(ctx context.Context) ([]*models.Infaq, error)
	CreateInfaqRiwayat(ctx context.Context, infaqRiwayat *models.InfaqRiwayat) (*models.InfaqRiwayat, error)
	GetById(ctx context.Context, idInfaq uint) (*models.Infaq, error)
}

type infaqRepository struct {
	db *gorm.DB
}

func NewInfaqRepository(db *gorm.DB) InfaqRepository {
	return &infaqRepository{db: db}
}

func (r *infaqRepository) CreateInfaq(ctx context.Context, infaq *models.Infaq) (*models.Infaq, error) {
	tx := r.db.WithContext(ctx).Create(&infaq)
	return infaq, tx.Error
}

func (r *infaqRepository) GetList(ctx context.Context) ([]*models.Infaq, error) {
	var infaq []*models.Infaq
	tx := r.db.WithContext(ctx).Find(&infaq)
	return infaq, tx.Error
}

func (r *infaqRepository) CreateInfaqRiwayat(ctx context.Context, infaqRiwayat *models.InfaqRiwayat) (*models.InfaqRiwayat, error) {
	tx := r.db.WithContext(ctx).Create(&infaqRiwayat)
	return infaqRiwayat, tx.Error
}

func (r *infaqRepository) GetById(ctx context.Context, idInfaq uint) (*models.Infaq, error) {
	var infaq *models.Infaq
	tx := r.db.WithContext(ctx).Find(&infaq, idInfaq)
	return infaq, tx.Error
}

func (r *infaqRepository) UpdateDanaInfaq(ctx context.Context, nominal int64, infaq *models.Infaq) error {
	tx := r.db.Debug().WithContext(ctx).Model(&models.Infaq{}).Where("id = ?", infaq.ID).Update("dana_terkumpul", gorm.Expr("dana_terkumpul + ?", nominal))
	return tx.Error
}
