package main

import (
	"app/cmd/application"
	"fmt"
)

func main() {
	app := application.NewApplicationDefault("")
	// - tear down
	defer app.TearDown()
	// - set up
	if err := app.SetUp(); err != nil {
		fmt.Println(err)
		return
	}
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
