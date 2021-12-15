package main;

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

// STUDENT STRUCT 
type Student struct {
	ID  int `json:"id"`
	Name string `json:"name"`
	Age int `json: "age"`
	Subject string `json: "subject"`
	Class string `json: "class"`

}
//Init student var as slice 
var students []Student;
//get all students

func getStudents(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students);
}

// get a student detail
func getStudent(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r) // get parameter
	id, err := strconv.Atoi(params["id"]);
	if err != nil{
		fmt.Println(err);
	}
	w.Header().Set("Content-Type", "application/json")
	for _, item := range students{
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(&Student{})
}

// create a new student
func createStudent(w http.ResponseWriter, r *http.Request){

}

// delete a student
func deleteStudents(w http.ResponseWriter, r *http.Request){

}
//get
func main()  {
	// Init ROuter
	r := mux.NewRouter();

	//mock data
	students  = append(students, Student{ID:1, Name:"MAYANK",Age: 45, Subject :"HINDI", Class: "9"});
	students  = append(students, Student{ID:2, Name:"JOHN",Age: 44, Subject :"math", Class: "9"});
	// Route Handlers endpoint
	r.HandleFunc("/students/students", getStudents).Methods("GET")
	r.HandleFunc("/students/student/{id}", getStudent).Methods("GET")
	r.HandleFunc("/students/student", createStudent).Methods("POST")
	r.HandleFunc("/students/student/{id}", deleteStudents).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000",r))

}