package main

import (
	"fmt"
)

func verifySalary(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
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
