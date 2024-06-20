package service

import (
	"github.com/luizengdev/banking/domain"
	"github.com/luizengdev/banking/dto"
	"github.com/luizengdev/banking/errs"
)

// CustomerService defines the service operations available to customers.
type CustomerService interface {
	GetAllCustomers(string) (*[]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

// DefaultCustomerService is the default implementation of CustomerService.
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomers fetches all customers using the associated repository and returns a CustomerResponse slice.
func (s DefaultCustomerService) GetAllCustomers(status string) (*[]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	var responses []dto.CustomerResponse
	for _, c := range customers {
		responses = append(responses, c.ToDto())
	}

	return &responses, nil
}

// GetCustomer fetches a customer using the associated repository and returns a CustomerResponse.
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

// NewCustomerService creates a new instance of the DefaultCustomerService with the provided repository.
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
