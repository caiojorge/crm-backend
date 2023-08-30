package infra

import (
	"github.com/caiojorge/crm-backend/internal/model"
	"github.com/google/uuid"
	"sync"
)

// MockDatabase represents the mock database using a map and a mutex.
type MockDatabase struct {
	mu        sync.Mutex
	customers map[uuid.UUID]model.Customer
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{
		customers: make(map[uuid.UUID]model.Customer),
	}
}

func (db *MockDatabase) Create(customer model.Customer) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, exists := db.customers[customer.ID]; exists {
		return false
	}
	db.customers[customer.ID] = customer
	return true
}

func (db *MockDatabase) Read(id uuid.UUID) (model.Customer, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	customer, exists := db.customers[id]
	return customer, exists
}

func (db *MockDatabase) ListAll() []model.Customer {
	db.mu.Lock()
	defer db.mu.Unlock()

	var all []model.Customer
	for _, customer := range db.customers {
		all = append(all, customer)
	}
	return all
}

func (db *MockDatabase) Update(id uuid.UUID, customer model.Customer) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.customers[id]; !exists {
		return false
	}
	db.customers[id] = customer
	return true
}

func (db *MockDatabase) Delete(id uuid.UUID) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.customers[id]; !exists {
		return false
	}
	delete(db.customers, id)
	return true
}
