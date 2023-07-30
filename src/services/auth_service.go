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

type AuthService interface {
	Login(ctx context.Context, auth *models.AuthRequest) (*models.User, error)
	GetUserSession(ctx *gin.Context) (*models.User, error)
}

type authService struct {
	repository repositories.AuthRepository
}

func NewAuthService(repository repositories.AuthRepository) AuthService {
	return &authService{
		repository: repository,
	}
}

func (e *authService) Login(ctx context.Context, auth *models.AuthRequest) (*models.User, error) {

	res, err := e.repository.Get(ctx, auth)
	if err != nil {
		zap.L().Error("error get user", zap.Error(err))
		return nil, err
	}

	if err := res.CheckPassword(auth.Password); err != nil {
		zap.L().Error("invalid credential", zap.Error(err))
		return nil, fmt.Errorf("invalid credential")
	}

	return res, nil
}

func (e *authService) GetUserSession(ctx *gin.Context) (*models.User, error) {
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
