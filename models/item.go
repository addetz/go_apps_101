package models

// EXERCISE 1: Define Item struct
// ----------------------------
type Item struct {
    ID string
    Name string
    Price float64
    CurrencyCode string
}

// EXERCISE 3: Print the price of a menu Item
// ----------------------------------------
func (i Item) PrintWithPrice() string {
	//present the item in the format Name: PriceCurrency
	//NOTE: fmt.Sprintf allows us to format strings just like with fmt.Println
	//https://golang.org/pkg/fmt/#Sprintf
	return ""
}

// EXERCISE 4: Change price of a menu item
// -------------------------------------
//func (? ???) ChangePrice(price float64, code string) {
// Should this method take a pointer receiver?
// assign new price & currency code here
//}
