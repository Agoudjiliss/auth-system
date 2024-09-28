package server

import (
  "github.com/go-chi/chi/v5"
  "github.com/go-chi/chi/v5/middleware"
)


func Routing() *chi.Mux{
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Get("/ping",ping)
  return r
}
