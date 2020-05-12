package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rikkinovar/movie-catalog-api/models"
)

var instance *gorm.DB

func initDB() (*gorm.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"))

	if os.Getenv("ENV") == "DEV" {
		connString = fmt.Sprintf("%s sslmode=disable", connString)
	}

	return gorm.Open("mysql", connString)
}

func GetConn() (*gorm.DB, error) {
	if instance == nil {
		conn, err := initDB()
		if err != nil {
			return nil, err
		}
		instance = conn
	}

	return instance, nil
}

// Migrate creates tables for available models.
func Migrate(dbConn *gorm.DB) {
	dbConn.AutoMigrate(
		&models.Movie{},
	)
}
