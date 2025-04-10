package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("Missing required database environment variables")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database")

	db.AutoMigrate(&User{})

	email := "testuser@gmail.com"
	var user User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		user = User{Email: &email}
		db.Create(&user)
		log.Println("Test user created successfully")
	} else {
		log.Println("Test user already exists")
	}
	return db
}
