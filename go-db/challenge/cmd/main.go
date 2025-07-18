package main

import (
	"app/cmd/application"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := &application.ConfigApplicationDefault{
		Db: &mysql.Config{
			User:   "root",
			Passwd: "root",
			Net:    "tcp",
			Addr:   "localhost:3306",
			DBName: "products",
		},
		Addr: "127.0.0.1:8080",
	}
	app := application.NewApplicationDefault(cfg)
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
