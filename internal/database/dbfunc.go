package database
import (
  "github.com/agoudjiliss/auth-system/data"
)


func InsertUser(user datatype.User) error{
  query :="INSERT INTO users (username,password,) VALUES ($1,$2)"
  _,err := db.Exec(query,user.UserName,user.Password)
  if err != nil {
    return err
  }
  return nil
}


func InsertToken(user datatype.User)error{
  query :="INSERT INTO refresh_tokens (user_id, token, expires_at)VALUES ($1,$2,$3);"

  _,err := db.Exec(query,user.Id,user.Token.Token,user.Token.Expires_at)
  if err != nil{
    return err
  }
  return nil
}

func SelectID(user datatype.User) (datatype.User,error){
  query :="SELECT id FROM users WHERE username = $1"
  err := db.QueryRow(query,user.UserName).Scan(&user.Id)
  if err != nil{
    return user ,err

  }
  return user,nil
}

