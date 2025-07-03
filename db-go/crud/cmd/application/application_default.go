package application

import (
	"app/cmd/router"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
)

type ApplicationDefault struct {
	rt   *chi.Mux
	addr string
}

func NewApplicationDefault(addr string) (a *ApplicationDefault) {
	// default config
	defaultRouter := chi.NewRouter()
	defaultAddr := ":8080"
	if addr != "" {
		defaultAddr = addr
	}

	a = &ApplicationDefault{
		rt:   defaultRouter,
		addr: defaultAddr,
	}
	return
}

func (a *ApplicationDefault) TearDown() (err error) {
	return
}

func (a *ApplicationDefault) SetUp() (err error) {
	config := &mysql.Config{
		User:   "root",
		Passwd: "1234",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "my_db",
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	a.rt.Use(middleware.Logger)
	a.rt.Use(middleware.Recoverer)

	a.rt.Route("/products", func(r chi.Router) {
		r.Mount("/", router.ProductRouter(db))
	})

	a.rt.Route("/warehouses", func(r chi.Router) {
		r.Mount("/", router.WarehouseRouter(db))
	})

	return
}

func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.addr, a.rt)
	return
}
