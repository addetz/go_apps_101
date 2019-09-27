package models

// EXERCISE 1: Define Item struct
// ----------------------------
type Item struct {
	// define ID as string,
	// Name as string
	// Price as float64,
	// CurrencyCode as string
}

// EXERCISE 3: Print the price of a menu Item
// ----------------------------------------
func (i Item) PrintWithPrice() string {
	//present the item in the format Name: PriceCurrencyCode
	//NOTE: fmt.Sprintf allows us to format strings just like with fmt.Println
	//https://golang.org/pkg/fmt/#Sprintf
	// Use the format string "%s: %v%s"
	return ""
}

// EXERCISE 4: Change price of a menu item
// -------------------------------------
//func (? ???) ChangePrice(price float64, code string) {
// Should this method take a pointer receiver?
// assign new price & currency code here
//}
