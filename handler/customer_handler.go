package handler

import (
	"encoding/json"
	"fmt"
	"github.com/caiojorge/crm-backend/internal/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

// DB MockDatabase represents the mock database using a map and a mutex.
var DB *model.MockDatabase

// GetAllCustomers SetDBOnActiveRecord sets the database on the ActiveRecord
func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(DB.ListAll()) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error": "there is no customer."}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.ListAllCustomers())
}

// UpdateCustomer updates a customer by id.
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	uid, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Printf("Failed to parse UUID: %v\n", err)
		return
	}
	fmt.Println(uid, id)
	var customer model.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, exist := model.FindCustomerByID(uid)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error": "customer not found."}`)
		return
	}
	fmt.Println(customer)
	customer.ID = uid
	ok := customer.Update()
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "error to update customer."}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

// CreateCustomer creates a new customer.
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var customer model.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(customer)
	ok := customer.Create()
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "error to create customer."}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

// FindCustomerById finds a customer by id.
func FindCustomerById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	uid, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Printf("Failed to parse UUID: %v\n", err)
		return
	}
	fmt.Println(uid, id)
	customer, exist := model.FindCustomerByID(uid)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error": "customer not found."}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

// DeleteCustomer deletes a customer by id.
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	uid, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Printf("Failed to parse UUID: %v\n", err)
		return
	}
	fmt.Println(uid, id)
	customer, exist := model.FindCustomerByID(uid)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error": "customer not found."}`)
		return
	}

	ok := customer.Delete()
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "error to delete customer."}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}
