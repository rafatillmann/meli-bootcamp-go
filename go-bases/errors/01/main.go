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

func verifySalary(salary int) error {
	if salary < 150000 {
		return &salaryError{"Error: the salary entered does not reach the taxable minimum"}
	}

	return nil
}

func main() {
	salary := 10000

	err := verifySalary(salary)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Must pay tax")
}
