package datatype
import(
  "github.com/dgrijalva/jwt-go"
)

type User struct{
   Id int64 
   UserName string `json:"username"`
   Password string `json:"password"`
   role string 
   Token struct{
     Token string
     Expires_at string
   }
}

type Configuration struct{
  Server struct{
    Host string
    Port string
  }
  Db struct{
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

