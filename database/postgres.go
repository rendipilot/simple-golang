package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDatabase() (*sqlx.DB, error){
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT") 
	dbUser := os.Getenv("DB_USER") 
	dbPassword := os.Getenv("DB_PASSWORD") 
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser,dbPassword,dbName)
	db, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		return nil, err
	}
	
	err = db.Ping() 
	if err != nil { 
		return nil, err 
	} 
	
	return db, nil
}