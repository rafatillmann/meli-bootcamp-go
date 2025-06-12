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
	data   map[int]domain.Product
	lastID int
}

func NewRepository() *repository {
	repository := repository{
		path: "products.json",
	}

	err := repository.load()
	if err != nil {
		log.Fatal(err)
	}
	return &repository
}

func (r *repository) GetAll() []domain.Product {
	var values []domain.Product
	for _, product := range r.data {
		values = append(values, product)
	}
	return values
}

func (r *repository) GetByID(ID int) (*domain.Product, error) {
	product, found := r.data[ID]

	if !found {
		return nil, fmt.Errorf("resource not found")
	}

	return &product, nil
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

	r.data[product.ID] = *product

	if err := r.write(); err != nil {
		return fmt.Errorf("error while writing data: %w", err)
	}

	return nil
}

func (r *repository) Update(update *domain.Product) error {
	if _, found := r.data[update.ID]; !found {
		return fmt.Errorf("resource not found")
	}

	r.data[update.ID] = *update

	if err := r.write(); err != nil {
		return fmt.Errorf("error while writing data: %w", err)
	}
	return nil
}

func (r *repository) Delete(ID int) error {
	if _, found := r.data[ID]; !found {
		return fmt.Errorf("resource not found")
	}

	delete(r.data, ID)

	if err := r.write(); err != nil {
		return fmt.Errorf("error while writing data: %w", err)
	}
	return nil
}

func (r *repository) load() error {
	file, err := os.Open(r.path)
	if err != nil {
		return fmt.Errorf("the file was not found: %w", err)
	}
	defer file.Close()

	var values []domain.Product
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&values); err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	data := make(map[int]domain.Product)
	for _, product := range values {
		data[product.ID] = product
	}

	r.data = data
	r.lastID = values[len(values)-1].ID

	return nil
}

func (r *repository) write() error {
	file, err := os.OpenFile(r.path, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("the file was not found: %w", err)
	}
	defer file.Close()

	var values []domain.Product
	for _, product := range r.data {
		values = append(values, product)
	}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(values); err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}

	return nil
}
