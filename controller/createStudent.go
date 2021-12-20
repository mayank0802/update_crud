package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"last.com/config"
	"last.com/model"
	"math/rand"

)
func CreateStudent(w http.ResponseWriter, r *http.Request,) {
	db := config.Getdb()
	var student model.Student

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&student)
	student.ID = rand.Intn(1000000)
	
		_,e := db.Query("INSERT INTO Student (NAME, AGE, SUBJECT, CLASS) VALUES(?, ?,?,?)", student.Name, student.Age, student.Subject, student.Class)
		if e != nil {
			fmt.Println(e)
		}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	
	json.NewEncoder(w).Encode(student)
}