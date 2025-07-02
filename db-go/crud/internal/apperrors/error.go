package apperrors

import "errors"

var (
	ErrRepositoryProductNotFound = errors.New("product not found")
	ErrRepositoryFetchData       = errors.New("error fetching data from database")
	ErrRepositorySaveData        = errors.New("error saving data to database")
	ErrRepositoryUpdateData      = errors.New("error updating data to database")
	ErrRepositoryDeleteData      = errors.New("error deleting data to database")
	ErrRepositoryGetLastId       = errors.New("error getting last id from database")
)
