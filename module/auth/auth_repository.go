package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jplesperance/passwordless-auth-system/db"
	"github.com/jplesperance/passwordless-auth-system/utils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

type IAuthRepository interface {
	FindFirstUserByEmail(email string) (*db.User, error)
	StoreEmailLoginVerificationSessionToRedis(userId uint, token string) error
}

func NewAuthRepository(db *gorm.DB, rdb *redis.Client) *AuthRepository {
	return &AuthRepository{db, rdb}
}

func (AuthRepository *AuthRepository) FindFirstUserByEmail(email string) (db.User, error) {
	var user db.User
	err := AuthRepository.db.Model(&user).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return db.User{}, utils.NewHttpResponse(404, "User doesn't exist", nil)
		} else {
			return db.User{}, utils.NewHttpResponse(500, "Other db errors", err)
		}
	}

	return user, nil
}

func (authRepository *AuthRepository) StoreEmailLoginVerificationSessionToRedis(userId uint, token string) error {
	ctx := context.Background()

	key := fmt.Sprintf("login_attempt_%v:%v", userId, token)

	err := authRepository.rdb.Set(ctx, key, "unverified", time.Minute*15).Err()
	if err != nil {
		return utils.NewHttpResponse(500, "Unknown error", err)
	}
	return nil
}
