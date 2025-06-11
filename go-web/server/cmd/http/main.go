package main

import (
	"log"
	"net/http"
	"server/cmd/http/router"

	"github.com/go-chi/chi/v5"
	"github.com/lpernett/godotenv"
)

func main() {
	// Adicionar um config file
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	r := chi.NewRouter()

	r.Use(router.Logger, router.Auth)

	r.Route("/products", func(r chi.Router) {
		r.Mount("/", router.Routers())
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
