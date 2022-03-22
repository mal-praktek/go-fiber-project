package database

import (
	"fmt"
	"github.com/tegarsubkhan236/go-fiber-project/app/models"
	"github.com/tegarsubkhan236/go-fiber-project/pkg/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		env.GetEnv("DB_HOST", "127.0.0.1"),
		env.GetEnv("DB_USER", ""),
		env.GetEnv("DB_PASSWORD", ""),
		env.GetEnv("DB_NAME", ""),
		env.GetEnv("DB_PORT", "5432"),
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&models.User{})
}
