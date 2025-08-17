package repository

import "app/internal"

func NewProductsMap(db map[int]internal.Product) *ProductsMap {
	defaultDb := make(map[int]internal.Product)
	if db != nil {
		defaultDb = db
	}

	return &ProductsMap{
		db: defaultDb,
	}
}

type ProductsMap struct {
	db map[int]internal.Product
}

func (r *ProductsMap) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	p = make(map[int]internal.Product)

	for k, v := range r.db {
		if query.ID > 0 && query.ID != v.ID {
			continue
		}

		p[k] = v
	}

	return
}
