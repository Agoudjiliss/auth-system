package server

import (
	"fmt"
	"net/http"

	"github.com/agoudjiliss/auth-system/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func Routing() *chi.Mux{
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Get("/ping",ping)
  r.Put("/createuser",handler.CreateUser)
  r.Put("/connect",handler.Login)
   r.With(JWTMiddleware).Get("/protected-endpoint",hi)
  return r
}

func hi(w http.ResponseWriter,r *http.Request){
  fmt.Fprintln(w,"hellooooo")
}
