package models

import "fmt"

// EXERCISE 1: Define Item struct
// ----------------------------
type Item struct {
	ID           string
	Name         string
	Price        float64
	CurrencyCode string
}

// EXERCISE 3: Print the price of a menu Item
// ----------------------------------------
func (i Item) PrintWithPrice() string {
	return fmt.Sprintf("%s: %v%s", i.Name, i.Price, i.CurrencyCode)
}

// EXERCISE 4: Change price of a menu item
// -------------------------------------
func (i *Item) ChangePrice(p float64, code string) {
	i.Price = p
	i.CurrencyCode = code
}
