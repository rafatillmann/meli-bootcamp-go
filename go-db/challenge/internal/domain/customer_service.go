package domain

// ServiceCustomer is the interface that wraps the basic methods that a customer service should implement.
type ServiceCustomer interface {
	FindAll() (c []Customer, err error)
	Save(c *Customer) (err error)
	TotalByCondition() ([]CustomerCondition, error)
	Amount() ([]CustomerAmount, error)
}
