package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ConfigApplicationDefault struct {
	Addr string
}

func NewApplicationDefault(cfg *ConfigApplicationDefault) *ApplicationDefault {
	defaultRt  := chi.NewRouter()
	defaultCfg := &ConfigApplicationDefault{
		Addr: ":8080",
	}
	if cfg != nil {
		if cfg.Addr != "" {
			defaultCfg.Addr = cfg.Addr
		}
	}

	return &ApplicationDefault{
		rt:   defaultRt,
		addr: defaultCfg.Addr,
	}
}

type ApplicationDefault struct {
	rt *chi.Mux
	addr string
}

func (a *ApplicationDefault) TearDown() (err error) {
	return
}

func (a *ApplicationDefault) SetUp() (err error) {
	rpProduct := repository.NewProductsMap(nil)
	hdProduct := handler.NewProductsDefault(rpProduct)

	a.rt.Use(middleware.Logger)
	a.rt.Use(middleware.Recoverer)
	a.rt.Route("/product", func(r chi.Router) {
		r.Get("/", hdProduct.Get())
	})
	return
}

func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.addr, a.rt)
	return
}