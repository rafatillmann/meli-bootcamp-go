package main

import (
	"errors"
	"fmt"
)

var salaryErr error = errors.New("Error: salary is less than 10000")

func verifySalary(salary int) error {
	if salary < 10000 {
		return salaryErr
	}
	return nil
}

func main() {
	salary := 1000
    verifyErr := salaryErr

	err := verifySalary(salary)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(errors.Is(err, verifyErr))
		return
	}
}
