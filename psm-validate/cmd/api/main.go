package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"psm-validate/internal/infrastructure/router"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/account", router.CreateAccount)
	r.Get("/account/{id}", router.FindByID)

	r.Post("/transactions", router.SaveTransaction)

	http.ListenAndServe(":8080", r)
}
