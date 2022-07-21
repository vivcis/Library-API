package db

import (
	"fmt"
	"github.com/vivcis/library-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

//type Database struct {
//	sql *sql.DB
//	DB  *gorm.DB
//}

func SetUpDB() error {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbName=%s port=%ss sslmode=disable TimeZone=Africa/Lags", host, user, password, dbName, port)
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if db == nil {
		return fmt.Errorf("database was not initialized")
	}

	err = db.AutoMigrate(&models.Book{}, &models.User{}, &models.Register{})
	if err != nil {
		return fmt.Errorf("migration error: %v", err)
	}

	DB = db

	return nil
}
