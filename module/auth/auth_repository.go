package auth

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

type IAuthRepository interface {
}

func NewAuthRepository(db *gorm.DB, rdb *redis.Client) IAuthRepository {
	return &AuthRepository{
		db:  db,
		rdb: rdb,
	}
}
