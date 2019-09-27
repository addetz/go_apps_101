package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/go_apps_101/models"
	"github.com/go_apps_101/repo"
	"github.com/gorilla/mux"
)

type Handler struct {
	Repo *repo.OrderRepo
}

// Index is invoked by HTTP GET /
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	// Send an HTTP status & a hardcoded message
	writeResponse(w, http.StatusOK, fmt.Sprintf("Welcome to the Orders Service!"))
}

// EXERCISE 9: Implement OrderIndex & OrderDelete
// ----------------------------------------------

// OrderIndex is invoked by HTTP GET /orders
func (h *Handler) OrderIndex(w http.ResponseWriter, r *http.Request) {
	// Call the repository method corresponding to the operation

	// Send an HTTP success status & the return value from the repo
}

// OrderShow is invoked by HTTP GET /orders/orderId
func (h *Handler) OrderShow(w http.ResponseWriter, r *http.Request) {
	//Get the orderId as a route variable
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	// Call the repository method corresponding to the operation
	order, err := h.Repo.Find(orderId)
	// Handle any errors & write an error HTTP status & response
	if err != nil {
		writeResponse(w, http.StatusNotFound, fmt.Errorf("no order with id %s found:%v", orderId, err))
	}
	// Send an HTTP success status & the return value from the repo
	writeResponse(w, http.StatusOK, order)
}

// EXERCISE 9: Implement OrderIndex & OrderDelete
// ----------------------------------------------

// OrderDelete is invoked by HTTP DELETE /orders/orderId
func (h *Handler) OrderDelete(w http.ResponseWriter, r *http.Request) {
	//Get the orderId as a route variable
	//vars := mux.Vars(r)
	//orderId := vars["orderId"]
	// Call the repository method corresponding to the operation

	// Handle any errors & write an error HTTP error status & response

	// Send an HTTP success status & response
}

// OrderShow is invoked by HTTP POST /orders
func (h *Handler) OrderUpsert(w http.ResponseWriter, r *http.Request) {
	// Initialize an order to unmarshal request body into
	var order models.Order
	// Read the request body
	body, err := readRequestBody(r)
	// Handle any errors & write an error HTTP status & response
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, fmt.Errorf("invalid order body:%v", err))
	}
	// Unmarshal response to order var
	// Handle any errors & write an error HTTP status & response
	if err := json.Unmarshal(body, &order); err != nil {
		writeResponse(w, http.StatusUnprocessableEntity, fmt.Errorf("invalid order body:%v", err))
	}
	// Call the repository method corresponding to the operation
	err = h.Repo.DB.UpsertOrder(order)
	// Send an HTTP success status & the return value from the repo
	writeResponse(w, http.StatusOK, err)
}

// writeResponse is a helper method that allows to write and HTTP status & response
func writeResponse(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Fprintf(w, fmt.Sprintf("error encoding resp %v:%s", resp, err))
	}
}

// readRequestBody is a helper method that allows to read a request body and return any errors
func readRequestBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return []byte{}, err
	}
	if err := r.Body.Close(); err != nil {
		return []byte{}, err
	}
	return body, err
}
