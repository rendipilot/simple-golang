package database

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
)

var (
    db   *pgx.Conn
    once sync.Once
)

func ConnectDatabase() (*pgx.Conn, error) {
	var err error

	once.Do(func() {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
		db, err = pgx.Connect(context.Background(), connStr)
		
		if err != nil {
			return
		}

		err = db.Ping(context.Background())
	})

	return db, err
}

func GetDB() *pgx.Conn {
    return db
}
