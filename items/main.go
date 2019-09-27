package main

import (
    "fmt"
    "github.com/go_apps_101/models"
)

func main() {
    // EXERCISE 2: Create & print some items
    // ------------------------------------
    burger := models.Item{
        ID:           "item-burger",
        Name:         "Burger",
        Price:        10.0,
        CurrencyCode: "GBP",
    }
    pizza := models.Item{
        ID:           "item-pizza",
        Name:         "Pizza",
        Price:        12.5,
        CurrencyCode: "GBP",
    }
    fmt.Printf("%v \n", burger)
    fmt.Printf("%v \n", pizza)
}
