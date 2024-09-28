package server

import (
	"fmt"
	"log"
	"net/http"
	"github.com/agoudjiliss/auth-system/internal/config"
)



func Run() { 
 config,err := config.NewConfig()
 if err != nil{
   log.Fatal("error to lead configuration ",err)
 }
 fmt.Printf("server work in Port: %v",config.Serverconfig.Port)

 r :=Routing()
 err = http.ListenAndServe(":"+config.Serverconfig.Port, r)
 
 if  err != nil {
   log.Fatal("error to ranning server",err)
 }
}

