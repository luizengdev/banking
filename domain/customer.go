package domain

import (
	"github.com/luizengdev/banking/dto"
	"github.com/luizengdev/banking/errs"
)

// Customer represents a customer with personal and contact information.
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

// statusAsText returns the customer status as text.
func (c Customer) statusAsText() string {
	statusAsText := "active"

	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

// ToDto converts the customer to a CustomerResponse.
func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

// CustomerRepository defines the operations available for working with customer data.
type CustomerRepository interface {
	// status == 1 is active, status == 0 is inactive, status == "" is all
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
