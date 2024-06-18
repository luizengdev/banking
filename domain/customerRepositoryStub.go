package domain

// CustomerRepositoryStub is a test implementation of the CustomerRepository interface.
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// NewCustomerRepositoryStub creates and returns a new instance of CustomerRepositoryStub with pre-defined customers.
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "John", City: "New York", Zipcode: "10001", DateOfBirth: "01/01/2000", Status: "1"},
		{Id: "2", Name: "Jane", City: "San Francisco", Zipcode: "94101", DateOfBirth: "01/01/2000", Status: "1"},
		{Id: "3", Name: "Jill", City: "Los Angeles", Zipcode: "90001", DateOfBirth: "01/01/2000", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
