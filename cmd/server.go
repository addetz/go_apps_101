package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go_apps_101/db"
	"github.com/go_apps_101/handlers"
	"github.com/go_apps_101/repo"
)

func main() {
	orderRepo := &repo.OrderRepo{
		DB: db.InitDB(),
	}
	handler := &handlers.Handler{
		Repo: orderRepo,
	}

	// EXERCISE 8: Create & save an order
	// --------------------------------

	// Server config
	router := handlers.ConfigureServer(handler)
	fmt.Println("Listening on localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", router))

}
