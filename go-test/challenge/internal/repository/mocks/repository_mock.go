package mocks

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func NewProductsRepositoryMock() *repositoryMock {
	return &repositoryMock{}
}

func (m *repositoryMock) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	args := m.Called(query)
	return args.Get(0).(map[int]internal.Product), args.Error(1)
}
