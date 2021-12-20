package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"last.com/controller"
)
func main() {
	r := mux.NewRouter()
	// Route Handlers endpoint
	r.HandleFunc("/student", controller.GetStudent).Methods("GET")
	r.HandleFunc("/student/{id}", controller.GetStudentByID).Methods("GET")
	r.HandleFunc("/student", controller.CreateStudent).Methods("POST")
	r.HandleFunc("/student/{id}", controller.DeleteStudent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
