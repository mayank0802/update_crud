package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"last.com/config"
	"last.com/model"
	"math/rand"
)

func GetStudent(w http.ResponseWriter, r *http.Request) {
	var students model.Student
	var arrStudents []model.Student

	db := config.Connect()
	defer db.Close()
	row, err := db.Query("SELECT E_NO, NAME,AGE,CLASS,SUBJECT FROM Students")

	if err != nil {
		log.Fatal(err.Error())
	}

	for row.Next() {
		err = row.Scan(&students.ID, &students.Name, &students.Age, &students.Class, &students.Subject)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrStudents = append(arrStudents, students)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(arrStudents)

}

//creating a new student
func CreateStudent(w http.ResponseWriter, r *http.Request,) {
	db := config.Connect()
	defer db.Close()
	

	var student model.Student
	//rr := mux.NewRouter()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&student)
	student.ID = rand.Intn(1000000)
	fmt.Println(student.Name)

	err := r.ParseForm() 
	if err != nil{
    	   panic(err)
	}

	 name := r.FormValue("name")
	 fmt.Println(name)
	 
		_,e := db.Query("INSERT INTO Students (E_NO,NAME, AGE, SUBJECT, CLASS) VALUES(?,?, ?,?,?)",student.ID, student.Name, student.Age, student.Subject, student.Class)
		if e != nil {
			fmt.Println(e)
		}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	
	json.NewEncoder(w).Encode(student)
}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var std model.Student

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	row, errr := db.Query("SELECT * FROM Students where E_NO=?", id)
	if errr != nil {
		panic(err)
	}

	for row.Next(){
	err = row.Scan(&std.ID, &std.Name, &std.Age, &std.Subject, &std.Class)
	if err != nil {
		fmt.Println(err)
	}
	}
	w.Header().Set("Contet-Type", "application/json")
	json.NewEncoder(w).Encode(std)

}

func DeleteStudents(w http.ResponseWriter, r *http.Request){
	db := config.Connect()
	defer db.Close()

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	row, errr := db.Query("DELETE FROM Students where E_NO=?", id)
	if errr != nil {
		fmt.Println(errr)
	}
	w.Header().Set("Contet-Type", "application/json")
	json.NewEncoder(w).Encode(row)

}