package models

import "fmt"

// EXERCISE 1: Define Item struct
// ----------------------------
type Item struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Price float64 `json:"price"`
    CurrencyCode string `json:"currency_code"`

}

// EXERCISE 3: Print the price of a menu Item
// ----------------------------------------
func (i Item) PrintWithPrice() string {
	//present the item in the format Name: PriceCurrencyCode
	//NOTE: fmt.Sprintf allows us to format strings just like with fmt.Println
	//https://golang.org/pkg/fmt/#Sprintf
	// Use the format string "%s: %v%s"
	return ""
    return fmt.Sprintf("%s: %v%s", i.Name, i.Price, i.CurrencyCode)
}

// EXERCISE 4: Change price of a menu item
// -------------------------------------
func (i *Item) ChangePrice(p float64, code string) {
    i.Price = p
    i.CurrencyCode = code
}
