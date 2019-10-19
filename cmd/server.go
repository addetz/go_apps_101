package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go_apps_101/db"
	"github.com/go_apps_101/handlers"
	"github.com/go_apps_101/models"
	"github.com/go_apps_101/repo"
)

func main() {
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

	orderRepo := &repo.OrderRepo{
		DB: db.InitDB(),
	}
	handler := &handlers.Handler{
		Repo: orderRepo,
	}

	// EXERCISE 8: Create & save an order
	// --------------------------------
	// An order for a starting point

	// Server config
	router := handlers.ConfigureServer(handler)
	fmt.Println("Listening on localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", router))

}
