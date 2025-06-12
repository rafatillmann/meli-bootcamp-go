package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rafatillmann/meli-bootcamp-go/go-bases/challenge/internal/tickets"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Application encountered an error: %v\n", err)
		}
	}()

	var repository tickets.TicketRepository = tickets.NewTicketRepository()

	reader := bufio.NewReader(os.Stdin)

	var destination string
	fmt.Println("Enter the destination country: ")
	destination, _ = reader.ReadString('\n')
	destination = strings.TrimSpace(destination)

	total, err := repository.GetTotalTickets(destination)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Total tickets to %s: %d\n", destination, total)

	//--------------------

	var period string
	fmt.Println("Enter the period (early-morning, morning, afternoon, night): ")
	period, _ = reader.ReadString('\n')
	period = strings.TrimSpace(period)

	count, err := repository.GetCountByPeriod(period)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Count of people who bought tickets in %s: %d\n", period, count)

	//--------------------

	fmt.Println("Enter the destination country:")
	destination, _ = reader.ReadString('\n')
	destination = strings.TrimSpace(destination)

	average, err := repository.AverageDestination(destination)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Porcentage of people who bought tickets to %s: %.2f%%\n", destination, average)
}
