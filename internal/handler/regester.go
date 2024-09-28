package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	datatype "github.com/agoudjiliss/auth-system/data"
	"github.com/agoudjiliss/auth-system/internal/config"
	"github.com/agoudjiliss/auth-system/internal/database"
	"github.com/agoudjiliss/auth-system/tools"
	"github.com/dgrijalva/jwt-go"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var NewUser datatype.User
	var err error
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&NewUser)

	// Hash the password
	NewUser.Password, err = tools.HachePassword(NewUser.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(w, "error to haching Password: %s", err)
		return
	}

	// Insert user into the database
	NewUser.Id, err = database.InsertUser(NewUser)
	if err != nil {
		log.Fatalln("error to Insert User: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generate JWT
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &datatype.Claims{
		Username: NewUser.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.Jwt.Jwtkey)) // Assurez-vous que Jwtkey est de type []byte
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Optional: Create a refresh token and store it
	refreshToken := "some_generated_refresh_token" // Générez ici un refresh token sécurisé
	expiresAt := time.Now().Add(24 * time.Hour)    // Durée de validité du refresh token
	err = database.InsertToken(int(NewUser.Id), refreshToken, expiresAt)
	if err != nil {
		log.Fatalln("error in Insert Token:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Return the JWT and optional refresh token
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
  fmt.Fprint(w,tokenString)
}
