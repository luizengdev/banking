package domain

import "github.com/luizengdev/banking/errs"

// Customer represents a customer with personal and contact information.
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

// CustomerRepository defines the operations available for working with customer data.
type CustomerRepository interface {
	// status == 1 is active, status == 0 is inactive, status == "" is all
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
