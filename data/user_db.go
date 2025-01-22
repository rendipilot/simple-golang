package data

import (
	"context"
	"fmt"
	"log"
	"rendipilot/simple-golang/database"
	"rendipilot/simple-golang/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUserDatabase(user *models.User) error {
	db := database.GetDB()
	if db == nil {
		log.Println("database connection is not initialized")
		return nil
	} else {
		hashedPassword, err :=  bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		if err != nil {
            return err
        }

		query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`
		_, err = db.Exec(context.Background(), query, user.Name, user.Email, hashedPassword)
		return err
	}

}

func GetUsersData() ([]*models.User, error) {
	// Get the database connection
	db := database.GetDB()
	if db == nil {
		log.Println("Database connection is not initialized")
		return nil, fmt.Errorf("database connection is not initialized")
	}

	// Query to fetch all users
	rows, err := db.Query(context.Background(), "SELECT name, email, password FROM users")
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	// Slice to store user data
	var users []*models.User

	// Iterate through the rows and scan them into User structs
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.Name, &user.Email, &user.Password); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		users = append(users, user)
	}

	// Check for any error encountered while iterating
	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, err
	}

	return users, nil
}