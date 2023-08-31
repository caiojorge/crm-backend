package model

import (
	"fmt"
	"testing"
)

//just one test to show how to mock the database
func setupMockDB() {
	SetDBOnActiveRecord(NewMockDatabase())
	fmt.Println(DB.customers)
}

// TestNewCustomer tests the NewCustomer function
func TestNewCustomer(t *testing.T) {
	setupMockDB()
	fmt.Println(DB.customers)

	name := "John"
	role := "Developer"
	email := "john@example.com"
	phone := "123-456-7890"
	contacted := false

	customer := NewCustomer(name, role, email, phone, contacted)
	if customer.Name != name {
		t.Errorf("Expected Name %v, but got %v", name, customer.Name)
	}
	if customer.Role != role {
		t.Errorf("Expected Role %v, but got %v", role, customer.Role)
	}
}

// TestCustomerCRUD tests the Customer CRUD operations
func TestCustomerCRUD(t *testing.T) {
	setupMockDB()
	fmt.Println(DB.customers)

	customer := NewCustomer("John", "Developer", "john@example.com", "123-456-7890", false)
	customer2 := NewCustomer("John2", "Developer", "john2@example.com", "123-456-7891", false)
	created := customer.Create()
	if !created {
		t.Errorf("Expected customer to be created, but got %v", created)
	}

	// Check idempotency
	createdAgain := customer.Create()
	if createdAgain {
		t.Errorf("Expected customer not to be created again, but got %v", createdAgain)
	}

	created2 := customer2.Create()
	if !created2 {
		t.Errorf("Expected customer to be created, but got %v", created)
	}

	fmt.Println(DB.customers)
	c1, e1 := FindCustomerByID(customer.ID)
	if !e1 {
		t.Errorf("error to read customer %v", customer.ID)
	}
	c2, e2 := FindCustomerByID(customer2.ID)
	if !e2 {
		t.Errorf("error to read customer %v", customer2.ID)
	}
	fmt.Println(c1)
	fmt.Println(c2)

	fmt.Println(len(ListAllCustomers()))
	if len(ListAllCustomers()) != 2 {
		t.Errorf("Expected 2 customers, but got %v", len(ListAllCustomers()))
	}

	customer.Name = "John Doe blah blah"
	updated := customer.Update()
	if !updated {
		t.Errorf("Expected customer to be updated, but got %v", updated)
	}

	fmt.Println(customer)

	d1 := c2.Delete()
	if !d1 {
		t.Errorf("error to delete customer %v", c2.ID)
	}

	if len(ListAllCustomers()) != 1 {
		t.Errorf("Expected 2 customers, but got %v", len(ListAllCustomers()))
	}
}
