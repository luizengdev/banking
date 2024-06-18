package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/luizengdev/banking/errs"
	"github.com/luizengdev/banking/logger"
)

// CustomerRepositoryDb is a CustomerRepository implementation that uses a MySQL database.
type CustomerRepositoryDb struct {
	client *sqlx.DB
}

// FindAll fetches all customers from the database.
func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}
	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return customers, nil
}

// FindById fetches a customer from the database.
func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	findByIdSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	c := Customer{}
	err := d.client.Get(&c, findByIdSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

// NewCustomerRepositoryDb creates and returns a new instance of CustomerRepositoryDb.
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "user:password@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3) // Configures the maximum lifetime of connections
	client.SetMaxOpenConns(10)                 // Configure the maximum number of open connections
	client.SetMaxIdleConns(10)                 // Configure the maximum number of inactive connections

	return CustomerRepositoryDb{client}
}
