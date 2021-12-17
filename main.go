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
	r.HandleFunc("/students", controller.GetStudent).Methods("GET")
	r.HandleFunc("/students/student/{id}", controller.GetStudentByID).Methods("GET")
	r.HandleFunc("/students/student", controller.CreateStudent).Methods("POST")
	r.HandleFunc("/students/student/{id}", controller.DeleteStudents).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
