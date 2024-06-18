package service

import (
	"github.com/luizengdev/banking/domain"
	"github.com/luizengdev/banking/errs"
)

// CustomerService defines the service operations available to customers.
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

// DefaultCustomerService is the default implementation of CustomerService.
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomers fetches all customers using the associated repository.
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// GetCustomer fetches a customer using the associated repository.
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.FindById(id)
}

// NewCustomerService creates a new instance of the DefaultCustomerService with the provided repository.
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}