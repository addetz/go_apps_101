package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := strings.TrimLeft(html.EscapeString(r.URL.Path), "/")
		fmt.Fprintf(w, "%s %s %s ... ",
			strings.ToUpper(p), strings.Title(p), strings.ToLower(p))
	})
	fmt.Println("Starting server... ")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
