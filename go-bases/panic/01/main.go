package main

import (
	"fmt"
	"os"
)

func open(file string) {
	defer func() {
		fmt.Println("Execution completed")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recovered from panic: %v\n", err)
		}
	}()

	_, err := os.Open(file)

	if err != nil {
		panic(fmt.Sprintf("The indicated file was not found or is damaged: %v", err))
	}
}

func main() {
	open("customers.txt")
}
