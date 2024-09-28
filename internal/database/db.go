package database

import (
  "database/sql"
  _"github.com/lib/pq"
  "github.com/agoudjiliss/auth-system/internal/config"
  "fmt"
)
var db *sql.DB


func Connectiontodb() (*sql.DB, error) {
    config,err := config.NewConfig()
    cnxstr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
        config.Db.User,
        config.Db.Password,
        config.Db.Dbname,
        config.Db.Sslmode,
    )
    
    db, err = sql.Open("postgres", cnxstr)
    if err != nil {
        return nil, err
    }

    // Ping the database to ensure the connection is established
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}

func CreateUserTable(db *sql.DB) error{
  query := `CREATE TABLE IF NOT EXISTS users (
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
  query := `CREATE TABLE IF NOT EXISTS refresh_tokens (
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
