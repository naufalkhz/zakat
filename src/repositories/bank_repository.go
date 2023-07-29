package repositories

import (
	"context"

	"github.com/naufalkhz/zakat/src/models"
	"gorm.io/gorm"
)

type BankRepository interface {
	Create(ctx context.Context, bank *models.Bank) (*models.Bank, error)
	Get(ctx context.Context) ([]*models.Bank, error)
	GetById(ctx context.Context, idBank uint) (*models.Bank, error)
}

type bankRepository struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) BankRepository {
	return &bankRepository{db: db}
}

func (r *bankRepository) Create(ctx context.Context, bank *models.Bank) (*models.Bank, error) {

	tx := r.db.WithContext(ctx).Create(bank)
	return bank, tx.Error
}

func (r *bankRepository) Get(ctx context.Context) ([]*models.Bank, error) {
	var bank []*models.Bank
	tx := r.db.WithContext(ctx).Find(&bank)
	return bank, tx.Error
}

func (r *bankRepository) GetById(ctx context.Context, idBank uint) (*models.Bank, error) {
	var bank *models.Bank
	tx := r.db.WithContext(ctx).Find(&bank, idBank)
	return bank, tx.Error
}
