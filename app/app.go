package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luizengdev/banking/domain"
	"github.com/luizengdev/banking/service"
)

// Start configures and starts the HTTP server.
func Start() {

	// create new router lib mux.
	router := mux.NewRouter()

	// Configures request handlers.
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods("GET")

	// starting server
	log.Fatal(http.ListenAndServe(":8000", router))
}
