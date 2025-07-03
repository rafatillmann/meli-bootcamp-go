package router

import (
	"app/internal/handler"
	"app/internal/repository"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ProductRouter(db *sql.DB) http.Handler {
	rt := chi.NewRouter()

	rp := repository.NewRepositoryProductSql(db)
	hd := handler.NewHandlerProduct(rp)

	rt.Get("/{id}", hd.GetById())
	rt.Post("/", hd.Create())
	rt.Patch("/{id}", hd.Update())
	rt.Delete("/{id}", hd.Delete())

	return rt
}
