package auth

import (
	"errors"

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
}

func NewAuthRepository(db *gorm.DB, rdb *redis.Client) IAuthRepository {
	return &AuthRepository{
		db:  db,
		rdb: rdb,
	}
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
