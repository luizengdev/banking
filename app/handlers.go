package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luizengdev/banking/service"
)

// CustomerHandlers contains the services needed to handle customer requests.
type CustomerHandlers struct {
	service service.CustomerService
}

// getAllCustomers handles the HTTP GET request to fetch all customers.
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"John", "New York", "10001"},
	// 	{"Jane", "San Francisco", "94101"},
	// 	{"Jill", "Los Angeles", "90001"},
	// }

	customers, _ := ch.service.GetAllCustomers()

	// Sets the response content type based on the request header
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

// getCustomer handles the HTTP GET request to fetch a customer.
func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	customer, err := ch.service.GetCustomer(customerId)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, err.Message)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
