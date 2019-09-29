package models

// EXERCISE 3: Define the OrderStatus constant using iota
// ----------------------------------------------------
type OrderStatus int //Underlying type
const (
	ACCEPTED = iota
	PREPARING
	DELIVERING
	COMPLETED
	REJECTED
)
