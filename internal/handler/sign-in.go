package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/agoudjiliss/auth-system/data"
	"github.com/agoudjiliss/auth-system/internal/config"
	"github.com/agoudjiliss/auth-system/internal/database"
	"github.com/agoudjiliss/auth-system/tools"
	"github.com/dgrijalva/jwt-go"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userLogin datatype.User
	var err error
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&userLogin)

	// Récupérer l'utilisateur de la base de données
	user, err := database.GetUserByUsername(userLogin.UserName)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Vérifier le mot de passe
	if tools.Verifierpassword(userLogin.Password, user.Password){
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Générez le JWT comme dans CreateUser
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &datatype.Claims{
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.Jwt.Jwtkey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Créez un refresh token si nécessaire
	refreshToken := "some_generated_refresh_token"
	expiresAt := time.Now().Add(24 * time.Hour)
	err = database.InsertToken(int(user.Id), refreshToken, expiresAt)
	if err != nil {
		log.Fatalln("error in Insert Token:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Retournez le JWT et le refresh token
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
  fmt.Fprint(w,tokenString)
}
