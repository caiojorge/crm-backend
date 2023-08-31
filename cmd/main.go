package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caiojorge/crm-backend/handler"
	"github.com/caiojorge/crm-backend/internal/model"

	// Import the third-party gorilla/mux package
	"github.com/gorilla/mux"
)

var dictionary = map[string]string{
	"Go":     "A programming language created by Google.",
	"Gopher": "A software engineer who builds with Go.",
	"Golang": "Another name for Go.",
}

func getDictionary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dictionary)
}

// using the same structure as the dictionary, based on course class
func main() {
	// all handles will share de DB connection
	handler.DB = model.NewMockDatabase()

	//create 3 example customers
	createCustomer()

	// Instantiate a new router by invoking the "NewRouter" handler
	router := mux.NewRouter()

	// Rather calling http.HandleFunc, call the equivalent router.HandleFunc
	// This gives us access to method-based routing
	router.HandleFunc("/crm/api/v1/dictionary", getDictionary).Methods("GET")
	router.HandleFunc("/crm/api/v1/customer/{id}", handler.FindCustomerById).Methods("GET")
	router.HandleFunc("/crm/api/v1/customer", handler.GetAllCustomers).Methods("GET")
	router.HandleFunc("/crm/api/v1/customer", handler.CreateCustomer).Methods("POST")
	router.HandleFunc("/crm/api/v1/customer/{id}", handler.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/crm/api/v1/customer/{id}", handler.DeleteCustomer).Methods("DELETE")
	// Serve static files from the 'static' directory
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	fmt.Println("Server is starting on port 3000...")
	// The second argument for "ListenAndServe" was previously nil
	/// Now that we are using our own custom router, we pass it along to "ListenAndServe" as its second argument
	http.ListenAndServe(":3000", router)
}

//
func createCustomer() {
	customer := model.NewCustomer("Caio", "Developer", "caio@yellowdroid.com", "+ 55 48 999339999", false)
	customer2 := model.NewCustomer("Junior", "Dog-Developer", "dog@yellowdroid.com", "+ 55 48 999339999", false)
	customer3 := model.NewCustomer("Patty", "Developer", "patty@eyellowdroid.com", "+ 55 48 999339999", false)

	model.SetDBOnActiveRecord(handler.DB)

	customer.Create()
	customer2.Create()
	customer3.Create()

	fmt.Println(model.ListAllCustomers())
}
