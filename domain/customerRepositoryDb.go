package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/luizengdev/banking/errs"
)

// CustomerRepositoryDb is a CustomerRepository implementation that uses a MySQL database.
type CustomerRepositoryDb struct {
	client *sql.DB
}

// FindAll fetches all customers from the database.
func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		c := Customer{}
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

// FindById fetches a customer from the database.
func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	findByIdSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(findByIdSql, id)
	c := Customer{}
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer table ", err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

// NewCustomerRepositoryDb creates and returns a new instance of CustomerRepositoryDb.
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3) // Configures the maximum lifetime of connections
	client.SetMaxOpenConns(10)                 // Configure the maximum number of open connections
	client.SetMaxIdleConns(10)                 // Configure the maximum number of inactive connections

	return CustomerRepositoryDb{client}
}
