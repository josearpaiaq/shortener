package db

import (
	"fmt"
	"log"

	"github.com/josearpaiaq/shortener/models"
	"github.com/josearpaiaq/shortener/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func ConnectToDB() (*gorm.DB, error) {
	dbHost := utils.GET_ENV("DB_HOST", "localhost")
	dbPort := utils.GET_ENV("DB_PORT", "5432")
	dbUser := utils.GET_ENV("DB_USER", "postgres")
	dbPassword := utils.GET_ENV("DB_PASSWORD", "password")
	dbName := utils.GET_ENV("DB_NAME", "db")
	dbSslmode := utils.GET_ENV("DB_SSLMODE", "disable")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
		dbSslmode,
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	
	log.Println("DB Connection started successfully")

	// Assign the connection to the global variable
	Connection = db

	// Migrate the schema
	migrateTables()

	return db, nil
}

func migrateTables() error {
	err := Connection.AutoMigrate(&models.URL{})
	if err != nil {
		return err
	}
	return nil
}