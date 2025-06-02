package main

import (
	"errors"
	"fmt"
)

const (
	CategoryA = "A"
	CategoryB = "B"
	CategoryC = "C"
)

func salary(category string, minutes float64) (float64, error) {
	switch category {
	case CategoryA:
		return 1000 * (minutes / 60), nil
	case CategoryB:
		result := 1500 * (minutes / 60)
		return result * 1.20, nil
	case CategoryC:
		result := 3000 * (minutes / 60)
		return result * 1.50, nil
	default:
		return 0, errors.New("Unexpected error")
	}
}

func main() {
	var category string
	var minutes float64
	fmt.Print("Informe a categoria: ")
	fmt.Scan(&category)
	fmt.Print("Informe a quantia de minutos trabalhados: ")
	fmt.Scan(&minutes)

	salary, err := salary(category, minutes)

	if err == nil {
		fmt.Printf("O salário é %.2f\n", salary)
	} else {
		fmt.Println("Não foi possível calcular o salário")
	}
}
