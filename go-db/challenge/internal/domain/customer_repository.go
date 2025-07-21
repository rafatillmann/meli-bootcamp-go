package domain

// RepositoryCustomer is the interface that wraps the basic methods that a customer repository should implement.
type RepositoryCustomer interface {
	FindAll() (c []Customer, err error)
	Save(c *Customer) (err error)
	TotalByCondition() ([]CustomerCondition, error)
	Amount() ([]CustomerAmount, error)
}
