package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	// Import the third-party gorilla/mux package
	"github.com/gorilla/mux"
)

var dictionary = map[string]string{
	"Go":     "A programming language created by Google.",
	"Gopher": "A software engineer who builds with Go.",
	"Golang": "Another name for Go.",
}

var customerDB = map[string]string{
	"Go":     "A programming language created by Google.",
	"Gopher": "A software engineer who builds with Go.",
	"Golang": "Another name for Go.",
}

func getDictionary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dictionary)
}

func main() {
	// Instantiate a new router by invoking the "NewRouter" handler
	router := mux.NewRouter()

	// Rather calling http.HandleFunc, call the equivalent router.HandleFunc
	// This gives us access to method-based routing
	router.HandleFunc("/dictionary", getDictionary).Methods("GET")
	router.HandleFunc("/model/{id}", getDictionary).Methods("GET")
	router.HandleFunc("/model", getDictionary).Methods("GET")
	router.HandleFunc("/model", getDictionary).Methods("POST")
	router.HandleFunc("/model", getDictionary).Methods("PUT")
	router.HandleFunc("/model/{id}", getDictionary).Methods("DELETE")

	fmt.Println("Server is starting on port 3000...")
	// The second argument for "ListenAndServe" was previously nil
	/// Now that we are using our own custom router, we pass it along to "ListenAndServe" as its second argument
	http.ListenAndServe(":3000", router)
}
