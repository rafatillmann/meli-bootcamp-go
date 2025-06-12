package api

import (
	"regexp"
	"time"
)

type ProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	CodeValue   string  `json:"code_value" validate:"required"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
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
