package repositories

import (
	"context"
	"fmt"

	"github.com/naufalkhz/zakat/src/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	// Create(ctx context.Context, user *models.User) (*models.User, error)
	Get(ctx context.Context, auth *models.AuthRequest) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

// func (r *authRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
// 	tx := r.db.WithContext(ctx).Create(&user)
// 	return user, tx.Error
// }

func (r *authRepository) Get(ctx context.Context, auth *models.AuthRequest) (*models.User, error) {
	var res *models.User
	tx := r.db.WithContext(ctx).Where("email = ?", auth.Email).First(&res)

	fmt.Println(tx.Error)
	return res, tx.Error
}
