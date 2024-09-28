package database

import (
  "database/sql"
  _"github.com/lib/pq"
  "github.com/agoudjiliss/auth-system/internal/config"
)
var db *sql.DB

func connectiontodb()(*sql.DB ,error){
  cnxstr := "user = "+config.Config.Dbconfig.User+" password ="+config.Config.Dbconfig.Password+" dbname ="+config.Config.Dbconfig.Dbname+" sslmode ="+config.Config.Dbconfig.Sslmode
  db,err := sql.Open("postgres",string(cnxstr))
  if err != nil{
    return nil,err
  }
  return db,nil 
}

func CreateUserTable(db *sql.DB) error{
  query := `CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );`
  _,err := db.Exec(query)
  if err != nil{
    return err
  }


  return nil
}

func CreateTokentable(db *sql.DB) error{
  query := `CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    token TEXT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id) 
        REFERENCES users(id)
        ON DELETE CASCADE
);`
  _,err := db.Exec(query)
  if err != nil {
    return err
  }
  return nil
}
