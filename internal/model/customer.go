package model

import (
	"fmt"
	"github.com/google/uuid"
)

// Customer is a ActiveRecord
type Customer struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Contacted bool      `json:"contacted"`
}

// DB MockDatabase represents the mock database using a map and a mutex.
var DB *MockDatabase

// NewCustomer creates a new Customer.
func NewCustomer(name, role, email, phone string, contacted bool) *Customer {
	return &Customer{
		ID:        uuid.New(),
		Name:      name,
		Role:      role,
		Email:     email,
		Phone:     phone,
		Contacted: contacted,
	}
}

// SetDBOnActiveRecord sets the database on the ActiveRecord
func SetDBOnActiveRecord(db *MockDatabase) {
	DB = db
}

//
func (c *Customer) Create() bool {
	fmt.Println(c.ID)

	if c.ID.String() == "00000000-0000-0000-0000-000000000000" {
		c.ID = uuid.New()
	}

	fmt.Println(c)
	return DB.Create(c)
}

// FindCustomerByID finds a customer by id
func FindCustomerByID(id uuid.UUID) (Customer, bool) {
	customer, exists := DB.Read(id)
	return customer, exists
}

// ListAllCustomers lists all customers
func ListAllCustomers() []Customer {
	return DB.ListAll()
}

// Update UpdateCustomer updates a customer
func (c *Customer) Update() bool {
	return DB.Update(c.ID, c)
}

// Delete DeleteCustomer deletes a customer
func (c *Customer) Delete() bool {
	return DB.Delete(c.ID)
}
