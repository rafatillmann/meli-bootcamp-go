package main

import (
	"fmt"

	"github.com/rafatillmann/meli-bootcamp-go/go-bases/challenge/internal/tickets"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Application encountered an error: %v\n", err)
		}
	}()

	tickets.PopulateTickets()

	var destination string
	fmt.Println("Enter the destination country: ")
	fmt.Scanln(&destination)

	total, err := tickets.GetTotalTickets(destination)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Total tickets to %s: %d\n", destination, total)

	//--------------------

	var period string
	fmt.Println("Enter the time (early-morning, morning, afternoon, night): ")
	fmt.Scanln(&period)

	count, err := tickets.GetCountByPeriod(period)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Count of people who bought tickets in %s: %d\n", period, count)

	//--------------------

	fmt.Println("Enter the destination country:")
	fmt.Scanln(&destination)

	average, err := tickets.AverageDestination(destination)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Porcentage of people who bought tickets to %s: %.2f%%\n", destination, average)
}
