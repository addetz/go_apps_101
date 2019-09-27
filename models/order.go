package models

// EXERCISE 5: Define the Order model
// ---------------------------------
type Order struct {
	ID string
	// Status as OrderStatus
	// Items as slice of item pointers to avoid copying them
}

// EXERCISE 6: Implement Order total
// -------------------------------
func (or Order) Total() float64 {
	//implement total here
	return 0.0
}
