package repositories

import (
	"context"

	"github.com/naufalkhz/zakat/src/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Get(ctx context.Context) (*models.User, error)
	GetUserById(ctx context.Context, userId uint) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	tx := r.db.WithContext(ctx).Create(&user)
	return user, tx.Error
}

func (r *userRepository) Get(ctx context.Context) (*models.User, error) {
	var user *models.User
	tx := r.db.WithContext(ctx).Last(&user)
	return user, tx.Error
}
func (r *userRepository) GetUserById(ctx context.Context, userId uint) (*models.User, error) {
	var user *models.User
	tx := r.db.WithContext(ctx).Find(&user, userId)
	return user, tx.Error
}
