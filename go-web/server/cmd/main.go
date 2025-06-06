package main

import (
	"fmt"
	"server/repository"
)

func main() {
	repository := repository.NewRepository()
	fmt.Println(repository)
}
