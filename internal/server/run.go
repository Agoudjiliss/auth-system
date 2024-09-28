package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/agoudjiliss/auth-system/internal/config"
	"github.com/agoudjiliss/auth-system/internal/database"
)

var db *sql.DB

func Run() { 
 config,err := config.NewConfig()
 fmt.Print(config)
 if err != nil{
   log.Fatal("error to lead configuration ",err)
 }
 log.Println("lead configuration succesfuly !!!")
 db,err = database.Connectiontodb()
 if err != nil{
   log.Fatal("error Connection to db :",err)
 }
 log.Println("Connection to db succesfuly !!!!")
 err = database.CreateUserTable(db)
 if err != nil{
   log.Fatal("error to Create User Table :",err)
 }
 log.Println("Create User table succesfuly !!!!")
 err = database.CreateTokentable(db)
 if err != nil {
   log.Fatal("error to Create Token table",err)
 }
 log.Println("Create Token table succesfuly !!!")

 fmt.Printf("server work in Port: %v",config.Server.Port)
 r :=Routing()
 err = http.ListenAndServe(":"+config.Server.Port, r)
 
 if  err != nil {
   log.Fatal("error to ranning server",err)
 }
}

