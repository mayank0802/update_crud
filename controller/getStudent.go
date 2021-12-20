package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"last.com/config"
	"last.com/model"


)

func GetStudent(w http.ResponseWriter, r *http.Request) {
	var students model.Student
	var arrStudents []model.Student
db := config.Getdb()
	row, err := db.Query("SELECT E_NO, NAME,AGE,CLASS,SUBJECT FROM Student")

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