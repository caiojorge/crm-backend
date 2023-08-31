package model

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

// MockDatabase represents the mock database using a map and a mutex.
type MockDatabase struct {
	mu        sync.Mutex
	customers map[uuid.UUID]Customer
}

// NewMockDatabase creates a new MockDatabase.
func NewMockDatabase() *MockDatabase {
	return &MockDatabase{
		customers: make(map[uuid.UUID]Customer),
	}
}

// Create creates a new customer.
func (db *MockDatabase) Create(customer *Customer) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, exists := db.customers[customer.ID]; exists {
		return false
	}
	_customer := *customer
	db.customers[customer.ID] = _customer
	fmt.Println("Created customer", customer.ID)
	return true
}

// Read reads a customer by id.
func (db *MockDatabase) Read(id uuid.UUID) (Customer, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	customer, exists := db.customers[id]
	return customer, exists
}

// ListAll lists all customers.
func (db *MockDatabase) ListAll() []Customer {
	db.mu.Lock()
	defer db.mu.Unlock()

	var all []Customer
	for _, customer := range db.customers {
		all = append(all, customer)
	}
	return all
}

// Update updates a customer by id.
func (db *MockDatabase) Update(id uuid.UUID, customer *Customer) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	fmt.Println("Updating customer", id, "with", customer)
	if _, exists := db.customers[id]; !exists {
		return false
	}
	_customer := *customer
	db.customers[id] = _customer
	return true
}

// Delete deletes a customer by id.
func (db *MockDatabase) Delete(id uuid.UUID) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.customers[id]; !exists {
		return false
	}
	delete(db.customers, id)
	return true
}
