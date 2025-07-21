package application

import (
	"app/cmd/router"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migrate_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type ConfigApplicationDefault struct {
	Db   *mysql.Config
	Addr string
}

func NewApplicationDefault(config *ConfigApplicationDefault) *ApplicationDefault {
	defaultCfg := &ConfigApplicationDefault{
		Db:   nil,
		Addr: ":8080",
	}
	if config != nil {
		if config.Db != nil {
			defaultCfg.Db = config.Db
		}
		if config.Addr != "" {
			defaultCfg.Addr = config.Addr
		}
	}

	return &ApplicationDefault{
		cfgDb:   defaultCfg.Db,
		cfgAddr: defaultCfg.Addr,
	}
}

type ApplicationDefault struct {
	cfgDb   *mysql.Config
	cfgAddr string
	db      *sql.DB
	router  *chi.Mux
}

func (a *ApplicationDefault) SetUp() (err error) {

	a.db, err = sql.Open("mysql", a.cfgDb.FormatDSN())
	if err != nil {
		return
	}

	err = a.db.Ping()
	if err != nil {
		return
	}

	driver, err := migrate_mysql.WithInstance(a.db, &migrate_mysql.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./docs/db/mysql/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	a.router = chi.NewRouter()
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)

	a.router.Route("/customers", func(r chi.Router) {
		r.Mount("/", router.CustomerRouter(a.db))
	})
	a.router.Route("/products", func(r chi.Router) {
		r.Mount("/", router.ProductRouter(a.db))
	})
	a.router.Route("/invoices", func(r chi.Router) {
		r.Mount("/", router.InvoiceRouter(a.db))
	})
	a.router.Route("/sales", func(r chi.Router) {
		r.Mount("/", router.SaleRouter(a.db))
	})

	return
}

func (a *ApplicationDefault) Run() (err error) {
	defer a.db.Close()
	err = http.ListenAndServe(a.cfgAddr, a.router)
	return
}
