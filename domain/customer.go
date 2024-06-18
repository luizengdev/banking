package domain

import "github.com/luizengdev/banking/errs"

// Customer represents a customer with personal and contact information.
type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// CustomerRepository defines the operations available for working with customer data.
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(string) (*Customer, *errs.AppError)
}
