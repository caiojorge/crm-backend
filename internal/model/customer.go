package model

import (
	"github.com/google/uuid"
)

/*
The Active Record pattern is an architectural pattern found in software engineering
that stores in-memory object data in relational databases. In this pattern, an object
instance is tied to a single row in the database table. The object's member variables
correspond to the columns of the table and the behavior of the object maps to CRUD operations
in the database.
*/

type Customer struct {
	ID        uuid.UUID
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

// New creates a new Customer.
func New(name, role, email, phone string, contacted bool) *Customer {
	return &Customer{
		ID:        uuid.New(),
		Name:      name,
		Role:      role,
		Email:     email,
		Phone:     phone,
		Contacted: contacted,
	}
}
