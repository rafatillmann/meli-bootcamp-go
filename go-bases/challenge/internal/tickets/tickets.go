package tickets

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type TicketError struct {
	Message string
}

func (e *TicketError) Error() string {
	return e.Message
}

type Ticket struct {
	ID      int
	Name    string
	Email   string
	Country string
	Time    time.Time
	Price   float64
}

const (
	EarlyMorning = "early-morning"
	Morning      = "morning"
	Afternoon    = "afternoon"
	Night        = "night"
)

type TicketRepository interface {
	GetTotalTickets(destination string) (int, error)
	GetCountByPeriod(period string) (int, error)
	AverageDestination(destination string) (float64, error)
}

type repository struct {
	path    string
	tickets map[int]Ticket
}

func NewTicketRepository() *repository {
	repository := repository{
		path: "tickets.csv",
	}

	repository.tickets = repository.getTickets()

	return &repository
}

func (r *repository) getTickets() map[int]Ticket {
	file, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("the file was not found or is damaged: %v", err))
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var tickets = make(map[int]Ticket)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(fmt.Sprintf("error reading CSV record: %v", err))
		}

		ID, err := strconv.Atoi(record[0])
		if err != nil {
			panic(fmt.Sprintf("error converting ID to integer: %v", err))
		}

		layout := "15:04"
		time, err := time.Parse(layout, record[4])
		if err != nil {
			panic(fmt.Sprintf("error converting time to time.Time: %v", err))
		}

		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			panic(fmt.Sprintf("error converting price to float: %v", err))
		}

		ticket := Ticket{
			ID:      ID,
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Time:    time,
			Price:   price,
		}
		tickets[ID] = ticket
	}

	return tickets
}

func (r *repository) GetTotalTickets(destination string) (int, error) {
	if destination == "" {
		return 0, &TicketError{"unspecified destination"}
	}

	total := 0
	for _, ticket := range r.tickets {
		if ticket.Country == destination {
			total++
		}
	}
	return total, nil
}

func (r *repository) GetCountByPeriod(period string) (int, error) {
	var startHour, endHour string
	switch period {
	case EarlyMorning:
		startHour = "00:00"
		endHour = "06:59"
	case Morning:
		startHour = "07:00"
		endHour = "12:59"
	case Afternoon:
		startHour = "13:00"
		endHour = "19:59"
	case Night:
		startHour = "20:00"
		endHour = "23:59"
	default:
		return 0, &TicketError{"unspecified period"}
	}

	layout := "15:04"

	startTime, err := time.Parse(layout, startHour)
	if err != nil {
		return 0, fmt.Errorf("error while parsing time: %v", err)
	}

	endTime, err := time.Parse(layout, endHour)
	if err != nil {
		return 0, fmt.Errorf("error while parsing tim: %v", err)
	}

	total := 0
	for _, ticket := range r.tickets {
		if (ticket.Time.Compare(startTime) >= 0) && (ticket.Time.Compare(endTime) <= 0) {
			total++
		}
	}
	return total, nil
}

func (r *repository) AverageDestination(destination string) (float64, error) {
	if destination == "" {
		return 0.00, &TicketError{"unspecified destination"}
	}

	count := 0
	for _, ticket := range r.tickets {
		if ticket.Country == destination {
			count++
		}
	}

	return (float64(count) / float64(len(r.tickets))) * 100, nil

}
