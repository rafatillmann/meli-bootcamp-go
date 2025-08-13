package internal

type RepositoryProducts interface {
	SearchProducts(query ProductQuery) (p map[int]Product, err error)
}