package domain

import (
	"regexp"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func ValidateExpiration(date string) bool {
	dateRegex := regexp.MustCompile(`^(\d{2})/(\d{2})/(\d{4})$`)

	matches := dateRegex.FindStringSubmatch(date)
	if matches == nil {
		return false
	}

	dateString := matches[3] + "-" + matches[2] + "-" + matches[1]
	_, err := time.Parse("2006-01-02", dateString)
	return err == nil
}
