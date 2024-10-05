package repositories

import (
	"golang_server.dankbueno.com/internal/models"
)

// Insert user into database
func CreateUser(user models.User) error {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	_, err := db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(username string) (models.User, error) {
	query := "SELECT id, password FROM users WHERE username = ?"
	var user models.User
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
