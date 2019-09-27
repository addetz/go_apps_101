package repo

import (
	"github.com/go_apps_101/db"
	"github.com/go_apps_101/models"
)

// OrderRepo is the repo for all things Order
type OrderRepo struct {
	DB *db.DB
}

// EXERCISE 7: Implement the OrderRepo
// ---------------------------------

// FindAll returns all orders existing from the database
func (or *OrderRepo) FindAll() []models.Order {
	// implementation here
	return []models.Order{}
}

// Find returns an order with a specific id from the database
func (or *OrderRepo) Find(id string) (*models.Order, error) {
	// implementation here
	return nil, nil
}

// Upsert inserts or updates a model into the database
func (or *OrderRepo) Upsert(o *models.Order) error {
	// implementation here
	return nil
}

// Delete deletes a model from the database
func (or *OrderRepo) Delete(id string) error {
	// implementation here
	return nil
}
