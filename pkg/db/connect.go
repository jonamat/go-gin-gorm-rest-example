package db

import (
	"fmt"
	"os"

	"github.com/jonamat/go-gin-gorm-rest-example/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Gorm *gorm.DB

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TZ"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.Service{},
		&models.Booking{},
	)

	Gorm = db
	return db, nil
}
