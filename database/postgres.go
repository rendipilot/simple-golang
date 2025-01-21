package database

import (
    "context"
    "fmt"
    "os"

    "github.com/jackc/pgx/v5"
)


func ConnectDatabase() (*pgx.Conn, error) {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
    conn, err := pgx.Connect(context.Background(), connStr)

    if err != nil {
        return nil, err
    }

    err = conn.Ping(context.Background())
    if err != nil {
        return nil, err
    }

    return conn, nil
}
