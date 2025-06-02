package main

import "fmt"

func main() {
	var age int
	var employee bool
	var years float64
	var salary float64

	fmt.Print("Qual sua idade? ")
	fmt.Scan(&age)

	fmt.Print("Você está empregado? ")
	fmt.Scan(&employee)

	if employee {
		fmt.Print("Quanto tempo está empregado? ")
		fmt.Scan(&years)

		fmt.Print("Qual seu salário? ")
		fmt.Scan(&salary)
	}

	if age > 22 && years > 1 {
		if salary > 100.000 {
			fmt.Println("Você é elegível ao empréstimo sem juros")
		} else {
			fmt.Println("Você é elegível ao empréstimo com juros")
		}
	} else {
		fmt.Println("Não é possível lhe conceder um empréstimo")
	}
}
