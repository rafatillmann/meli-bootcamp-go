package main

import (
	"log"
	"net/http"
	"os"
	"server/cmd/http/router"
	"server/pkg/response"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lpernett/godotenv"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != os.Getenv("API_TOKEN") {
			response.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger, AuthMiddleware)

	r.Route("/products", func(r chi.Router) {
		r.Mount("/", router.ProductRouters())
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
