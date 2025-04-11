package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RegisterAuthRouter(router *gin.RouterGroup, db *gorm.DB, rdb *redis.Client) {

}
