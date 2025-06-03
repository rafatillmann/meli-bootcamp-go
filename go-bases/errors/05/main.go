package main

import (
	"fmt"
)

type salaryError struct {
	message string
}

func (e *salaryError) Error() string {
	return e.message
}

func salary(hours float64, value float64) (float64, error) {
	if hours < 80 {
		return 0, &salaryError{"Error: the worker cannot have worked less than 80 hours per month"}
	}

	salary := hours * value
	if salary >= 150000.00 {
		salary *= 0.9
	}
	return salary, nil
}

func main() {
	salary, err := salary(100, 5000)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Salary: %.2f\n", salary)
}
