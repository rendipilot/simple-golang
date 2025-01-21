package data

import (
    "context"
    "rendipilot/simple-golang/models"

    "github.com/jackc/pgx/v5"
)

var db *pgx.Conn

func CreateUserDatabase(user *models.User) error {
    query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`
    _, err := db.Exec(context.Background(), query, user.Name, user.Email, user.Password)
    return err
}
