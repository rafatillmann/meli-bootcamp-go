package service

import "app/internal/domain"

// NewCustomersDefault creates new default service for customer entity.
func NewCustomersDefault(rp domain.RepositoryCustomer) *CustomersDefault {
	return &CustomersDefault{rp}
}

// CustomersDefault is the default service implementation for customer entity.
type CustomersDefault struct {
	// rp is the repository for customer entity.
	rp domain.RepositoryCustomer
}

// Amount implements domain.ServiceCustomer.
func (s *CustomersDefault) Amount() ([]domain.CustomerAmount, error) {
	return s.rp.Amount()
}

// TotalByCondition implements domain.ServiceCustomer.
func (s *CustomersDefault) TotalByCondition() ([]domain.CustomerCondition, error) {
	return s.rp.TotalByCondition()
}

// FindAll returns all customers.
func (s *CustomersDefault) FindAll() (c []domain.Customer, err error) {
	c, err = s.rp.FindAll()
	return
}

// Save saves the customer.
func (s *CustomersDefault) Save(c *domain.Customer) (err error) {
	err = s.rp.Save(c)
	return
}
