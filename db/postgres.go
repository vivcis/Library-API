package db

import (
	"fmt"
	"github.com/vivcis/library-app/helpers"
	"github.com/vivcis/library-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

//PostgresDb implements the DB interface
type PostgresDb struct {
	DB *gorm.DB
}

// SetUpDB sets up the postgres instance
func (pdb *PostgresDb) SetUpDB(config *helpers.Config) (*gorm.DB, error) {
	//fmt.Println("connecting to database...")
	//
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbName=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, user, password, dbName, port)
	//var err error
	//if os.Getenv("DATABASE_URL") != "" {
	//	dsn = os.Getenv("DATABASE_URL")
	//}
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//if db == nil {
	//	return fmt.Errorf("database was not initialized")
	//} else {
	//	fmt.Println("connected to database")
	//}

	var dsn string
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort, config.DBMode, config.DBTimeZone)
	} else {
		dsn = databaseURL
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Established database connection")

	pdb.DB = db

	err = pdb.DB.AutoMigrate(&models.Register{}, &models.Book{}, &models.User{})
	if err != nil {
		log.Fatalf("migration error: %v", err)
	}
	//err = pdb.PrePopulateTables()
	//if err != nil {
	//	log.Println("error repopulating database", err)
	//	return err
	//}

	return db, nil
}

func (pdb *PostgresDb) PrePopulateTables() error {
	err := pdb.DB.AutoMigrate(&models.User{}, &models.Register{}, &models.Book{})
	if err != nil {
		return fmt.Errorf("error migrating data: %v", err)
	}
	register := models.Register{
		Model:        models.Model{},
		UserID:       "7777777777",
		User:         models.User{},
		BookID:       "33333333",
		Book:         models.Book{},
		BorrowedDate: 55555,
		ReturnDate:   777777,
	}
	result := pdb.DB.Where("register = ?", "7777777777").Find(&register)
	if result.RowsAffected < 1 {
		pdb.DB.Create(&register)
	}

	user := models.User{
		Model:        models.Model{},
		UserName:     "cece",
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "cecilia.orji@decagon.dev",
		Password:     "1234567890",
		PasswordHash: "$2a$12$T2wSf1qgpTyhLOons3u4JOCqCwKDDL4J3UhGdOTEBL/CmAS/RNCPm",
	}
	result = pdb.DB.Where("user = ?", "John").Find(&user)
	if result.RowsAffected < 1 {
		pdb.DB.Create(&user)
	}

	book := models.Book{
		Model:    models.Model{},
		Title:    "Are you afraid of the dark?",
		Author:   "Sidney Sheldon",
		Copies:   5,
		ISBN:     "55555566666777778888899999",
		Returned: false,
	}
	result = pdb.DB.Where("book = ?", "Sidney Sheldon").Find(&book)
	if result.RowsAffected < 1 {
		pdb.DB.Create(&book)
	}
	return nil
}
