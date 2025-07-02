package ticket

import (
	"chanllenge/internal/domain"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// NewLoaderCsv creates a new ticket loader from a CSV file
func NewLoaderCsv(filePath string) *LoaderCsv {
	return &LoaderCsv{
		filePath: filePath,
	}
}

// LoaderCsv represents a ticket loader from a CSV file
type LoaderCsv struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (t *LoaderCsv) Load() (db map[int]domain.TicketAttributes, err error) {
	// open the file
	f, err := os.Open(t.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records
	db = make(map[int]domain.TicketAttributes)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			err = fmt.Errorf("error reading record: %v", err)
			return nil, err
		}

		// serialize the record
		id, _ := strconv.Atoi(record[0])
		price, _ := strconv.ParseFloat(record[5], 64)
		ticket := domain.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   price,
		}

		// add the ticket to the map
		db[id] = ticket
	}

	return
}
