package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
)

// NewApplicationDefault creates a new default application.
func NewApplicationDefault(addr, filePathStore string) (a *ApplicationDefault) {
	// default config
	defaultRouter := chi.NewRouter()
	defaultAddr := ":8080"
	if addr != "" {
		defaultAddr = addr
	}

	a = &ApplicationDefault{
		rt:            defaultRouter,
		addr:          defaultAddr,
		filePathStore: filePathStore,
	}
	return
}

// ApplicationDefault is the default application.
type ApplicationDefault struct {
	// rt is the router.
	rt *chi.Mux
	// addr is the address to listen.
	addr string
	// filePathStore is the file path to store.
	filePathStore string
}

// TearDown tears down the application.
func (a *ApplicationDefault) TearDown() (err error) {
	return
}

// SetUp sets up the application.
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// - store
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

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	// - repository
	rp := repository.NewRepositoryDb(db)
	// - handler
	hd := handler.NewHandlerProduct(rp)

	// router
	// - middlewares
	a.rt.Use(middleware.Logger)
	a.rt.Use(middleware.Recoverer)
	// - endpoints
	a.rt.Route("/products", func(r chi.Router) {
		// GET /products/{id}
		r.Get("/{id}", hd.GetById())
		// POST /products
		r.Post("/", hd.Create())
		// PUT /products/{id}
		r.Put("/{id}", hd.UpdateOrCreate())
		// PATCH /products/{id}
		r.Patch("/{id}", hd.Update())
		// DELETE /products/{id}
		r.Delete("/{id}", hd.Delete())
	})

	return
}

// Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.addr, a.rt)
	return
}
