package db

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Database interface {
}

type DBParams struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

func InitDBParams() DBParams {
	ginMode := os.Getenv("GIN_MODE")
	log.Println(ginMode)
	if ginMode != "release" {
		errEnv := godotenv.Load()
		if errEnv != nil {
			log.Fatal("Error loading .env file")
		}
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	return DBParams{
		Host:     host,
		User:     user,
		Password: password,
		DbName:   dbName,
		Port:     port,
	}
}
