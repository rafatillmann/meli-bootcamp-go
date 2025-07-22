package test

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
)

func init() {
	config := &mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "localhost:3307",
		DBName: "products",
	}

	txdb.Register("txdb", "mysql", config.FormatDSN())
}

func GetTxdb() (*sql.DB, error) {
	return sql.Open("txdb", "products")
}
