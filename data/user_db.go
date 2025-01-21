package data

import (
	"rendipilot/simple-golang/models"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func CreateUserDatabase(user *models.User) error {
	query := `INSER INTO users (name, email, password) VALUES (:name, :email, :password)`
	_, err := db.NamedExec(query, user)
	return err
}