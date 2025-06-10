package product

import (
	"encoding/json"
	"fmt"
	"os"
	"server/internal/domain"
)

type Repository interface {
	Get() []domain.Product
	GetByID(ID int) domain.Product
	GetFilterByPrice(price float64) []domain.Product
	Create(request domain.ProductRequest) domain.Product
	load(path string) ([]domain.Product, error)
}

type repository struct {
	data []domain.Product
}

func NewRepository() *repository {
	repository := repository{
		data: make([]domain.Product, 0),
	}

	data, err := repository.load("products.json")
	if err != nil {
		panic(err)
	}
	repository.data = data
	return &repository
}

func (r *repository) Get() []domain.Product {
	return r.data
}

func (r *repository) GetByID(ID int) domain.Product {
	var result domain.Product
	for _, product := range r.data {
		if product.ID == ID {
			result = product
		}
	}
	return result
}

func (r *repository) GetFilterByPrice(price float64) []domain.Product {
	var result []domain.Product
	for _, product := range r.data {
		if product.Price > price {
			result = append(result, product)
		}
	}
	return result
}

func (r *repository) Create(request domain.ProductRequest) domain.Product {
	product := domain.Product{
		ID:          len(r.data) + 1,
		Name:        request.Name,
		Quantity:    request.Quantity,
		CodeValue:   "8392",
		IsPublished: false,
		Expiration:  "10/12/2025",
		Price:       90.00,
	}
	r.data = append(r.data, product)
	return product
}

func (r *repository) load(path string) ([]domain.Product, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("the file was not found: %w", err)
	}
	defer file.Close()

	var data []domain.Product
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}
	return data, nil
}
