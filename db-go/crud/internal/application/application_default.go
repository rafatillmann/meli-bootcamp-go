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

	rp := repository.NewRepositoryDb(db)

	hd := handler.NewHandlerProduct(rp)

	a.rt.Use(middleware.Logger)
	a.rt.Use(middleware.Recoverer)
	a.rt.Route("/products", func(r chi.Router) {
		r.Get("/{id}", hd.GetById())
		r.Post("/", hd.Create())
		r.Put("/{id}", hd.UpdateOrCreate())
		r.Patch("/{id}", hd.Update())
		r.Delete("/{id}", hd.Delete())
	})

	return
}

func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.addr, a.rt)
	return
}
