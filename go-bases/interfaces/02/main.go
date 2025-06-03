package main

import "fmt"

const (
	small  = "small"
	medium = "medium"
	large  = "large"
)

type Product interface {
	Price() float64
}

type ProductSmall struct {
	PricePerUnit float64
}

func (p ProductSmall) Price() float64 {
	return p.PricePerUnit
}

type ProductMedium struct {
	PricePerUnit float64
}

func (p ProductMedium) Price() float64 {
	return p.PricePerUnit + (p.PricePerUnit * 0.03) + (p.PricePerUnit * 0.03)
}

type ProductLarge struct {
	PricePerUnit float64
}

func (p ProductLarge) Price() float64 {
	return p.PricePerUnit + (p.PricePerUnit * 0.06) + 2500
}

func factory(product string, pricePerUnit float64) Product {
	switch product {
	case small:
		return ProductSmall{pricePerUnit}
	case medium:
		return ProductMedium{pricePerUnit}
	case large:
		return ProductLarge{pricePerUnit}
	default:
		return nil
	}
}

func main() {
	productSmall := factory(small, 100)
	productMedium := factory(medium, 2000)
	productLarge := factory(large, 3000)

	fmt.Printf("Small Price: %.2f\n", productSmall.Price())
	fmt.Printf("Medium Price: %.2f\n", productMedium.Price())
	fmt.Printf("Large Price: %.2f\n", productLarge.Price())
}
