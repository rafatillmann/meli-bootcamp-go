package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println(employees["Benjamin"])

	for key, value := range employees {
		if value > 20 {
			fmt.Println(key)
		}
	}

	employees["Frederico"] = 25
	delete(employees, "Pedro")
}
