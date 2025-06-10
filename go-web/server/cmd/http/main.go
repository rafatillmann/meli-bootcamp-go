package main

import (
	"log"
	"net/http"
	"server/cmd/http/router"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/products", func(r chi.Router) {
		r.Mount("/", router.ProductRouters())
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
