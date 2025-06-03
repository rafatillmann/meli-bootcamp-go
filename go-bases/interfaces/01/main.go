package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Student struct {
	ID      int
	Name    string
	Surname string
	Date    time.Time
}

func (s Student) PrintStudent() {
	fmt.Printf("Name: %s\nSurname: %s\nID: %d\nDate of admission: %s\n", s.Name, s.Surname, s.ID, s.Date.Format("2006-01-02"))
}

func main() {
	var name string
	var surname string
	var date string

	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Surname: ")
	fmt.Scan(&surname)
	fmt.Print("Date (YYYY-MM-DD): ")
	fmt.Scan(&date)

	layout := "2006-01-02"
	time, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing:", err)
	}

	fmt.Println("-----")

	student := Student{
		ID:      rand.Intn(100),
		Name:    name,
		Surname: surname,
		Date:    time,
	}

	student.PrintStudent()
}
