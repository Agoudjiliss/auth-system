package server

import (
  "github.com/go-chi/chi/v5"
  "github.com/go-chi/chi/v5/middleware"
  "github.com/agoudjiliss/auth-system/internal/handler"
)


func Routing() *chi.Mux{
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Get("/ping",ping)
  r.Put("/createuser",handler.CreateUser)
  return r
}
