package domain

// RepositorySale is the interface that wraps the basic Sale methods.
type RepositorySale interface {
	FindAll() (s []Sale, err error)
	Save(s *Sale) (err error)
}