package main

import (
	"fmt"
	"os"
)

func read(file string) {
	defer func() {
		fmt.Println("Execution completed")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recovered from panic: %v\n", err)
		}
	}()

	data, err := os.ReadFile(file) // Automatically closes the file after reading

	if err != nil {
		panic(fmt.Sprintf("The indicated file was not found or is damaged: %v", err))
	}

	fmt.Println(string(data))
}

func main() {
	read("customers.txt")
}
