package internal

type ProductAttributes struct {
	Description	string
	Price    	float64
	SellerId 	int
}

type Product struct {
	Id       	int
	ProductAttributes
}

type ProductQuery struct {
	Id	   	int
}