package application

import (
	"chanllenge/cmd/http/handler"
	"chanllenge/internal/ticket"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ApplicationDefault struct {
	rt         *chi.Mux
	serverAddr string
	dbFile     string
}

type ConfigAppDefault struct {
	ServerAddr string
	DbFile     string
}

func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "tickets.csv",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}

	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}

func (a *ApplicationDefault) SetUp() (err error) {
	db, err := ticket.NewLoaderCsv(a.dbFile).Load()
	if err != nil {
		log.Fatal(err)
	}

	rp := ticket.NewRepository(db)
	sv := ticket.NewService(rp)
	handler := handler.NewTicketHandler(sv)

	a.rt.Route("/tickets", func(r chi.Router) {
		r.Get("/amount", handler.GetAmount())
		r.Get("/amount/{country}", handler.GetAmountByCountry())
		r.Get("/percentage/{country}", handler.GetPercentageByCountry())
	})

	return
}

func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
