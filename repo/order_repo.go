package repo

import (
	"github.com/go_apps_101/models"
)

// DB is the interface wrapping our database
type DB interface {
	DeleteOrder(id string) error
	FindAllOrders() []models.Order
	FindOrder(id string) (models.Order, error)
	UpsertOrder(o models.Order) error
}

// OrderRepo is the repo for all things Order
type OrderRepo struct {
	DB DB
}

// EXERCISE 7: Implement the OrderRepo
// ---------------------------------

// FindAll returns all orders existing from the database
func (or *OrderRepo) FindAll() []models.Order {
	return or.DB.FindAllOrders()
}

// Find returns an order with a specific id from the database
func (or *OrderRepo) Find(id string) (*models.Order, error) {
	order, err := or.DB.FindOrder(id)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// Upsert inserts or updates a model into the database
func (or *OrderRepo) Upsert(o *models.Order) error {
	return or.DB.UpsertOrder(*o)
}

// Delete deletes a model from the database
func (or *OrderRepo) Delete(id string) error {
	return or.DB.DeleteOrder(id)

}
