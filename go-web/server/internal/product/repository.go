package product

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"server/internal/domain"
)

type Repository interface {
	Get() []domain.Product
	GetByID(ID int) domain.Product
	GetFilterByPrice(price float64) []domain.Product
	Create(request domain.ProductRequest) domain.Product
}
type repository struct {
	path   string
	data   []domain.Product
	lastID int
}

func NewRepository() *repository {
	repository := repository{
		path: "products.json",
		data: make([]domain.Product, 0),
	}

	data, err := repository.load()
	if err != nil {
		log.Fatal(err)
	}
	repository.data = data
	repository.lastID = data[len(data)-1].ID
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
	r.lastID++

	product := domain.Product{
		ID:          r.lastID,
		Name:        request.Name,
		Quantity:    request.Quantity,
		CodeValue:   request.CodeValue,
		IsPublished: request.IsPublished,
		Expiration:  request.Expiration,
		Price:       request.Price,
	}
	r.data = append(r.data, product)

	if err := r.save(); err != nil {
		log.Fatal(err)
	}

	return product
}

func (r *repository) load() ([]domain.Product, error) {
	file, err := os.Open(r.path)
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

func (r *repository) save() error {
	file, err := os.OpenFile(r.path, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("the file was not found: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(r.data); err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}

	return nil
}
