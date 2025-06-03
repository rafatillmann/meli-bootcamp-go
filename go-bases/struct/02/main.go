package main

import (
	"fmt"
	"time"
)

type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

type Employee struct {
	Person
    Position string
}

func (e Employee) PrintEmployee(){
    fmt.Printf("ID: %d, Name: %s, Date of Birth: %s, Position: %s\n", e.ID, e.Name, e.DateOfBirth.Format("2006-01-02"), e.Position)
}

func main() {
    person := Person{
        ID:          1,
        Name:        "Rafaela Tillmann",
        DateOfBirth: time.Date(2000, 12, 11, 0, 0, 0, 0, time.UTC),
    }

    employee := Employee{
        Person:    person,
        Position: "Software Developer",
    }

    employee.PrintEmployee()
}
