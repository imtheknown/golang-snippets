package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AllItemsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Implemented It")
}

func FindItemEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Implemented It")
}

func CreateItemEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Implemented It")
}

func UpdateItemEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Implemented It")
}

func DeleteItemEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Implemented It")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/item", AllItemsEndpoint).Methods("GET")
	r.HandleFunc("/item", CreateItemEndpoint).Methods("POST")
	r.HandleFunc("/item", UpdateItemEndpoint).Methods("PUT")
	r.HandleFunc("/item", DeleteItemEndpoint).Methods("DELETE")
	r.HandleFunc("/item/{id}", FindItemEndpoint).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
