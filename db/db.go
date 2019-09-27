package db

import (
	"errors"
	"fmt"
	"sync"

	"github.com/go_apps_101/models"
)

type DB struct {
	m      sync.Mutex
	orders map[string]models.Order
}

func InitDB() *DB {
	return &DB{
		orders: make(map[string]models.Order),
	}
}

func (db *DB) UpsertOrder(o models.Order) error {
	db.m.Lock()
	defer db.m.Unlock()
	if _, ok := db.orders[o.ID]; !ok {
		o.ID = fmt.Sprintf("order-%d", len(db.orders))
	}
	db.orders[o.ID] = o
	return nil
}

func (db *DB) FindOrder(id string) (models.Order, error) {
	o, ok := db.orders[id]
	if !ok {
		return models.Order{}, errors.New(fmt.Sprintf("No model for id %s found", id))
	}
	return o, nil
}

// Maybe this can work without ranging
func (db *DB) FindAllOrders() []models.Order {
	var o []models.Order
	for _, order := range db.orders {
		o = append(o, order)
	}
	return o
}

func (db *DB) DeleteOrder(id string) error {
	db.m.Lock()
	defer db.m.Unlock()
	if _, ok := db.orders[id]; !ok {
		return errors.New(fmt.Sprintf("No model found for id %s to delete", id))
	}
	delete(db.orders, id)
	return nil
}
