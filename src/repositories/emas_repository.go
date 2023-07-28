package repositories

import (
	"context"

	"github.com/naufalkhz/zakat/src/models"
	"gorm.io/gorm"
)

type EmasRepository interface {
	Create(ctx context.Context, emas *models.Emas) (*models.Emas, error)
	Get(ctx context.Context) (*models.Emas, error)
}

type emasRepository struct {
	db *gorm.DB
}

func NewEmasRepository(db *gorm.DB) EmasRepository {
	return &emasRepository{db: db}
}

func (r *emasRepository) Create(ctx context.Context, emas *models.Emas) (*models.Emas, error) {
	var result *models.Emas
	tx := r.db.WithContext(ctx).Create(emas)
	return result, tx.Error
}

func (r *emasRepository) Get(ctx context.Context) (*models.Emas, error) {
	var emas *models.Emas
	tx := r.db.WithContext(ctx).Last(&emas)
	return emas, tx.Error
}
