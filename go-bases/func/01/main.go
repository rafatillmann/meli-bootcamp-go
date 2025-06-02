package main

import (
	"fmt"
	"math"
)

func tax(salary float64) float64 {
	switch {
	case salary > 50000:
		return math.Trunc(salary * 0.17)
	case salary > 150000:
		return math.Trunc(salary * 0.27)
	default:
		return 0
	}
}

func main() {
	var salary float64
	fmt.Print("Informe seu salário: ")
	fmt.Scan(&salary)

	fmt.Printf("Seu salário é %.2f e o total de imposto deduzido é %.2f\n", salary, tax(salary))
}
