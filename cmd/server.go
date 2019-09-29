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
	// EXERCISE 2: Create & print some items
	// ------------------------------------
	burger := models.Item{
		ID:           "item-burger",
		Name:         "Burger",
		Price:        float64(10.0),
		CurrencyCode: "GBP",
	}
	pizza := models.Item{
		ID:           "item-pizza",
		Name:         "Pizza",
		Price:        float64(12.5),
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
	order := models.Order{
		Items:  []*models.Item{&burger, &pizza},
		Status: models.ACCEPTED,
	}
	fmt.Printf("%v \n", order)
	orderRepo.Upsert(&order)

	// Server config
	router := handlers.ConfigureServer(handler)
	fmt.Println("Listening on localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", router))

}
