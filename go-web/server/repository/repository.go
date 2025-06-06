package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"server/model"
)

type Repository interface {
	load()
}

type repository struct {
	data []model.Product
}

func NewRepository() *repository {
	data, err := load("products.json")
	if err != nil {
		panic(err)
	}
	repository := repository{
		data: data,
	}
	return &repository
}

func load(path string) ([]model.Product, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("the file was not found: %w", err)
	}
	defer file.Close()

	var data []model.Product
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}
	return data, nil
}
