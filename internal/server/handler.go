package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/agoudjiliss/auth-system/data"
	"github.com/agoudjiliss/auth-system/internal/config"
	"github.com/agoudjiliss/auth-system/internal/database"
	"github.com/agoudjiliss/auth-system/tools"
	"github.com/dgrijalva/jwt-go"
)



func CreateUser(w http.ResponseWriter,r *http.Request){
  var NewUser datatype.User
  var err error
  w.Header().Set("Content-Type","application/json")
  json.NewDecoder(r.Body).Decode(&NewUser)
  NewUser.Password,err = tools.HachePassword(NewUser.Password)
  if err != nil{
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatalln(w,"error to haching Password: %s",err)
  }
  
  expirationTime := time.Now().Add(10 * time.Minute)
	claims := &datatype.Claims {
		Username: NewUser.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},}
    // Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.Jwt.Jwtkey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
  NewUser.Id,err = database.InsertUser(NewUser)
  if err != nil{
    log.Fatalln("error to Insert User: ",err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  err = database.InsertToken(int(NewUser.Id),tokenString,time.Now())
  if err != nil{
    log.Fatalln("error in Insert Token:",err)
    w.WriteHeader(http.StatusInternalServerError)
  }

	// Return the token
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
  response := map[string]string{"token": tokenString}
	json.NewEncoder(w).Encode(response)
}


