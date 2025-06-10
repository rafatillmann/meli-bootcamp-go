package product

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"server/internal/domain"
)

type Repository interface {
	GetAll() []domain.Product
	GetByID(ID int) (*domain.Product, error)
	GetFilterByPrice(price float64) []domain.Product
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(ID int) error
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

	data, err := repository.read()
	if err != nil {
		log.Fatal(err)
	}
	repository.data = data
	repository.lastID = data[len(data)-1].ID
	return &repository
}

func (r *repository) GetAll() []domain.Product {
	return r.data
}

func (r *repository) GetByID(ID int) (*domain.Product, error) {
	var result *domain.Product
	for _, product := range r.data {
		if product.ID == ID {
			result = &product
		}
	}

	if result == nil {
		return nil, fmt.Errorf("resource not found")
	}

	return result, nil
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

func (r *repository) Create(product *domain.Product) error {
	r.lastID++
	product.ID = r.lastID
	r.data = append(r.data, *product)

	if err := r.write(); err != nil {
		return fmt.Errorf("error while writing data: %w", err)
	}

	return nil
}

func (r *repository) Update(product *domain.Product) error {
	for i, p := range r.data {
		if p.ID == product.ID {
			r.data[i] = *product

			if err := r.write(); err != nil {
				return fmt.Errorf("error while writing data: %w", err)
			}
			return nil
		}
	}

	return fmt.Errorf("resource not found")
}

func (r *repository) Delete(ID int) error {
	for i, p := range r.data {
		if p.ID == ID {
			r.data = append(r.data[:i], r.data[i+1:]...)

			if err := r.write(); err != nil {
				return fmt.Errorf("error while writing data: %w", err)
			}
			return nil
		}
	}

	return fmt.Errorf("resource not found")
}

func (r *repository) read() ([]domain.Product, error) {
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

func (r *repository) write() error {
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
