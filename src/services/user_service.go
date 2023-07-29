package services

import (
	"context"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/repositories"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type UserService interface {
	Get(ctx context.Context) (*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Edit(ctx *gin.Context, user *models.User) (*models.User, error)
	GetUserSession(ctx *gin.Context) (*models.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (e *userService) Get(ctx context.Context) (*models.User, error) {
	user, err := e.repository.Get(ctx)
	if err != nil {
		zap.L().Error("error user", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (e *userService) Create(ctx context.Context, user *models.User) (*models.User, error) {

	if err := user.HashPassword(user.Password); err != nil {
		zap.L().Error("hasing password error", zap.Error(err))
		return nil, err
	}

	res, err := e.repository.Create(ctx, user)
	if err != nil {
		zap.L().Error("error create user", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (e *userService) Edit(ctx *gin.Context, userReq *models.User) (*models.User, error) {
	userTarget, err := e.GetUserSession(ctx)
	if err != nil {
		zap.L().Error("failed toget user", zap.Error(err))
		return nil, err
	}

	res, err := e.repository.Edit(ctx, userTarget, userReq)
	if err != nil {
		zap.L().Error("error create user", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (e *userService) GetUserSession(ctx *gin.Context) (*models.User, error) {
	userId, err := getUserId(ctx)
	if err != nil {
		zap.L().Error("failed parse token", zap.Error(err))
		return nil, err
	}

	user, err := e.repository.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func getUserId(ctx *gin.Context) (uint, error) {
	tokenString := ctx.GetHeader("Authorization")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		jwtSecret := os.Getenv("JWT_SECRET")
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed parse token")
	}

	var user = make(map[string]interface{})
	for key, val := range claims {
		user[key] = val
	}

	return cast.ToUint(user["id"]), nil
}
