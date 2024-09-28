package database
import (
  "github.com/agoudjiliss/auth-system/data"
  "time"
  "database/sql"
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

func GetUserByUsername(username string) (datatype.User, error) {
    var user datatype.User
    query := "SELECT id, username, password FROM users WHERE username = $1"
    
    // Exécutez la requête avec le nom d'utilisateur fourni.
    err := db.QueryRow(query, username).Scan(&user.Id, &user.UserName, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return user, nil // Aucun utilisateur trouvé.
        }
        return user, err // Retournez l'erreur en cas de problème.
    }
    
    return user, nil // Retournez l'utilisateur trouvé.
}
