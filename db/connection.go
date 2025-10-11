package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/josearpaiaq/shortener/utils"
	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	dbHost := utils.GET_ENV("DB_HOST")
	dbPort := utils.GET_ENV("DB_PORT")
	dbUser := utils.GET_ENV("DB_USER")
	dbPassword := utils.GET_ENV("DB_PASSWORD")
	dbName := utils.GET_ENV("DB_NAME")
	dbSslmode := utils.GET_ENV("DB_SSLMODE")

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
		dbSslmode,
	)

	db, err := sql.Open("postgres", connStr)
	
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	
	if err = db.Ping(); err != nil {
		log.Println("DB Ping Failed")
		log.Fatal(err)
	}
	log.Println("DB Connection started successfully")

	return db, nil
}