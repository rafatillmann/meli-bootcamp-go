package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Customer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func verifyIfCustomerExists(ID int) {
	data, err := os.ReadFile("customers.json")

	if err != nil {
		panic(fmt.Sprintf("The indicated file was not found or is damaged: %v", err))
	}

	var customers []Customer
	json.Unmarshal(data, &customers)

	for _, customer := range customers {
		if customer.ID == ID {
			panic(fmt.Sprintf("Customer with ID %d already exists", ID))
		}
	}
}

func validateAndInsertCustomer(ID int, name, phone, address string) (bool, error) {
	if ID == 0 {
		return false, errors.New("ID must be greater than 0")
	}
	if name == "" {
		return false, errors.New("name cannot be empty")
	}
	if phone == "" {
		return false, errors.New("phone cannot be empty")
	}
	if address == "" {
		return false, errors.New("address cannot be empty")
	}

	data, err := os.ReadFile("customers.json")

	if err != nil {
		panic(fmt.Sprintf("The indicated file was not found or is damaged: %v", err))
	}

	var customers []Customer
	json.Unmarshal(data, &customers)

	customers = append(customers, Customer{
		ID:      ID,
		Name:    name,
		Phone:   phone,
		Address: address,
	})

	jsonData, err := json.Marshal(customers)

	if err != nil {
		return false, fmt.Errorf("unable to marshal data: %w", err)
	}

	err = os.WriteFile("customers.json", jsonData, 0644)

	if err != nil {
		panic(fmt.Sprintf("Unable to write to file: %v", err))
	}

	return true, nil
}

func main() {
	defer func() {
		fmt.Println("End of execution")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Several errors were detected at runtime: %v\n", err)
		}
	}()

	var ID int
	var name string
	var phone string
	var address string

	fmt.Print("ID: ")
	fmt.Scan(&ID)
	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Phone number: ")
	fmt.Scan(&phone)
	fmt.Print("Address: ")
	fmt.Scan(&address)

	verifyIfCustomerExists(ID)

	_, err := validateAndInsertCustomer(ID, name, phone, address)

	if err != nil {
		fmt.Printf("Unable to insert customer: %v\n", err)
	}

	fmt.Println("Customer inserted successfully")
}
