package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n",
			product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}

func getById(id int) Product {
	for _, product := range Products {
		if product.ID == id {
			return product
		}
	}

	return Product{}
}

var Products = []Product{
	{ID: 1, Name: "Pen", Price: 2.5, Description: "Blue pen", Category: "Stationery"},
	{ID: 2, Name: "Notebook", Price: 1500, Description: "MacBook Pro", Category: "Electronics"},
}

func main() {
	product := Product{ID: 3, Name: "Pooh", Price: 50, Description: "Teddy Bear", Category: "Toys"}
	product.Save()
	product.GetAll()

	fmt.Println(getById(3))
}
