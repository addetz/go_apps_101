package models

// EXERCISE 5: Define the Order model
// ---------------------------------
type Order struct {
    ID     string `json: id`
    Items  []*Item `json: items`
    Status OrderStatus `json:"status"`
}

// EXERCISE 6: Implement Order total
// -------------------------------
func (or Order) Total() float64 {
    var sum float64
    for _, item := range or.Items {
        sum += item.Price
    }
    return sum
}
