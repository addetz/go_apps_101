package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my first Go server!")
	})
	fmt.Println("Starting server on 3030...")
	log.Fatal(http.ListenAndServe(":3030", nil))
}
