package repositories

import (
	"context"
	"fmt"

	"github.com/naufalkhz/zakat/src/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Get(ctx context.Context, auth *models.AuthRequest) (*models.User, error)
	GetUserById(ctx context.Context, userId uint) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Get(ctx context.Context, auth *models.AuthRequest) (*models.User, error) {
	var res *models.User
	tx := r.db.WithContext(ctx).Where("email = ?", auth.Email).First(&res)

	fmt.Println(tx.Error)
	return res, tx.Error
}

func (r *authRepository) GetUserById(ctx context.Context, userId uint) (*models.User, error) {
	var res *models.User
	tx := r.db.WithContext(ctx).Find(&res, userId)
	return res, tx.Error
}
