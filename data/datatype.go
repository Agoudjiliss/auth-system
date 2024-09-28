package datatype
import(
  "github.com/dgrijalva/jwt-go"
)

type User struct{
   Id int64 
   UserName string
   Password string
   role string
   Token struct{
     Token string
     Expires_at string
   }
}

type Configuration struct{
  Serverconfig struct{
    Host string
    Port string
  }
  Dbconfig struct{
    User string
    Password string
    Dbname string
    Sslmode string
  }
  Jwt struct{
   Jwtkey string 
  }

}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

