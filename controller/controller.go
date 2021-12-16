package controller

import (
	"encoding/json"
	"log"
	"net/http"

	//".../config"
	//"../model"
	"github.com/mayank0802/crud/config"

	"github.com/mayank0802/crud/model"
)

// git token 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
func getStudent(w http.ResponseWriter, r *http.Request) {
	var students model.Student
	var arrStudents []model.Student

	db := config.Connect()
	defer db.close()
	row, err := db.Query("SELECT E_NO, NAME FROM Students")

	if err != nil {
		log.Fatal(err.Error())
	}

	for row.Next() {
		err = row.Scan(&students.ID, &students.Name)
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
func createStudent(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	err := r.ParseMultiplepartForm(4096)
	if err != nil {
		panic(err)
	}
	name := r.FormValue("name")
	class := r.FormValue("class")
	subject := r.FormValue("student")
	age := r.FormValue("age")

	_, err = db.Exec("INSERT INTO STUDENTS (NAME, AGE, SUBJECT, CLASS)VALUES(?, ?,?,?)", name, age, subject, class)
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode("insert successfully")
}
