package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
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
	// - db: init
	a.db, err = sql.Open("mysql", a.cfgDb.FormatDSN())
	if err != nil {
		return
	}
	// - db: ping
	err = a.db.Ping()
	if err != nil {
		return
	}
	// - repository
	rpCustomer := repository.NewCustomersMySQL(a.db)
	rpProduct := repository.NewProductsMySQL(a.db)
	rpInvoice := repository.NewInvoicesMySQL(a.db)
	rpSale := repository.NewSalesMySQL(a.db)
	// - service
	svCustomer := service.NewCustomersDefault(rpCustomer)
	svProduct := service.NewProductsDefault(rpProduct)
	svInvoice := service.NewInvoicesDefault(rpInvoice)
	svSale := service.NewSalesDefault(rpSale)
	// - handler
	hdCustomer := handler.NewCustomersDefault(svCustomer)
	hdProduct := handler.NewProductsDefault(svProduct)
	hdInvoice := handler.NewInvoicesDefault(svInvoice)
	hdSale := handler.NewSalesDefault(svSale)

	a.router = chi.NewRouter()
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)

	a.router.Route("/customers", func(r chi.Router) {
		r.Get("/", hdCustomer.GetAll())
		r.Post("/", hdCustomer.Create())
	})
	a.router.Route("/products", func(r chi.Router) {
		r.Get("/", hdProduct.GetAll())
		r.Post("/", hdProduct.Create())
	})
	a.router.Route("/invoices", func(r chi.Router) {
		r.Get("/", hdInvoice.GetAll())
		r.Post("/", hdInvoice.Create())
	})
	a.router.Route("/sales", func(r chi.Router) {
		r.Get("/", hdSale.GetAll())
		r.Post("/", hdSale.Create())
	})

	return
}

func (a *ApplicationDefault) Run() (err error) {
	defer a.db.Close()
	err = http.ListenAndServe(a.cfgAddr, a.router)
	return
}
