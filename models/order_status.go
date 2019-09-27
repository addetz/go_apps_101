package models


// EXERCISE 3: Define the OrderStatus constant using iota
// ----------------------------------------------------

type OrderStatus int
const (
    ACCEPTED OrderStatus = iota
    PREPARING
    DELIVERING
    COMPLETED
    REJECTED
)