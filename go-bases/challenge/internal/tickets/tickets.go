package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Ticket struct {
	ID      int
	Name    string
	Email   string
	Country string
	Time    time.Time
	Price   float64
}

var tickets map[int]Ticket = make(map[int]Ticket)

const (
	EarlyMorning = "early-morning"
	Morning      = "morning"
	Afternoon    = "afternoon"
	Night        = "night"
)

func PopulateTickets() {
	file, err := os.Open("tickets.csv")
	if err != nil {
		panic(fmt.Sprintf("The file was not found or is damaged: %v", err))
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		panic(fmt.Sprint("Error reading CSV header: %v", err))
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(fmt.Sprintf("Error reading CSV record: %v", err))
		}

		ID, err := strconv.Atoi(record[0])
		if err != nil {
			panic(fmt.Sprintf("Error converting ID to integer: %v", err))
		}

		layout := "15:04"
		time, err := time.Parse(layout, record[4])
		if err != nil {
			panic(fmt.Sprintf("Error converting time to time.Time: %v", err))
		}

		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			panic(fmt.Sprintf("Error converting price to float: %v", err))
		}

		ticket := Ticket{
			ID:      ID,
			Name:    record[2],
			Email:   record[3],
			Country: record[4],
			Time:    time,
			Price:   price,
		}
		tickets[ID] = ticket
	}
}

func GetTotalTickets(destination string) (int, error) {
	if destination == "" {
		return 0, errors.New("Unspecified destination")
	}

	total := 0
	for _, ticket := range tickets {
		if ticket.Country == destination {
			total++
		}
	}
	return total, nil
}

func GetCountByPeriod(period string) (int, error) {
	switch period {
	case EarlyMorning:
		return 0, nil
	case Morning:
		return 0, nil
	case Afternoon:
		return 0, nil
	case Night:
		return 0, nil
	default:
		return 0, errors.New("Unspecified period")
	}
}

func AverageDestination(destination string) (float64, error) {}
