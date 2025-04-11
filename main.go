package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jplesperance/passwordless-auth-system/db"
	"github.com/jplesperance/passwordless-auth-system/module/auth"
	"github.com/jplesperance/passwordless-auth-system/rdb"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := gin.Default()
	api := app.Group("/api")
	DB := db.Init()
	RDB := rdb.Init()

	auth.RegisterAuthRouter(api, DB, RDB)

	app.Run(":8080")
}
