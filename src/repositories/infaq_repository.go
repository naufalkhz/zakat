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
	GetInfaqRiwayat(ctx context.Context, idUser uint) ([]*models.InfaqRiwayat, error)
	GetInfaqRiwayatLastLimit(ctx context.Context, limit int) ([]*models.InfaqRiwayat, error)
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

func (r *infaqRepository) GetInfaqRiwayat(ctx context.Context, idUser uint) ([]*models.InfaqRiwayat, error) {
	var infaqRiwayat []*models.InfaqRiwayat
	tx := r.db.Debug().WithContext(ctx).Where("id_user = ?", idUser).Find(&infaqRiwayat)
	return infaqRiwayat, tx.Error
}

func (r *infaqRepository) GetInfaqRiwayatLastLimit(ctx context.Context, limit int) ([]*models.InfaqRiwayat, error) {
	var infaqRiwayat []*models.InfaqRiwayat
	tx := r.db.Debug().WithContext(ctx).Limit(limit).Order("id desc").Find(&infaqRiwayat)
	return infaqRiwayat, tx.Error
}
