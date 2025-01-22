package data

import (
	"context"
	"log"
	"rendipilot/simple-golang/database"
	"rendipilot/simple-golang/models"
)

func CreateUserDatabase(user *models.User) error {
	db := database.GetDB()
	if db == nil {
		log.Println("database connection is not initialized")
		return nil
	} else {
		query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`
		_, err := db.Exec(context.Background(), query, user.Name, user.Email, user.Password)
		return err
	}

}
