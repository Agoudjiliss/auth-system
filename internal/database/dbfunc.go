package database
import (
  "github.com/agoudjiliss/auth-system/data"
  "time"
)

func InsertUser(user datatype.User) (int64, error) {
	// Correct SQL query
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"

	// Variable to hold the new user's ID
	var userID int64

	// Execute the query and return the new user's ID
	err := db.QueryRow(query, user.UserName, user.Password).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}


func InsertToken(userID int, token string, expiresAt time.Time) error {
	query := "INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3);"

	_, err := db.Exec(query, userID, token, expiresAt)
	if err != nil {
		return err
	}
	return nil
}

