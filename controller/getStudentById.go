package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"last.com/config"
	"last.com/model"

)

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	db := config.Getdb()
	var std model.Student

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	row, errr := db.Query("SELECT * FROM Student where E_NO=?", id)
	if errr != nil {
		panic(err)
	}

	for row.Next(){
	err = row.Scan(&std.ID, &std.Name, &std.Age, &std.Subject, &std.Class)
	if err != nil {
		panic(err)
		fmt.Println(err)
	}
	}
	w.Header().Set("Contet-Type", "application/json")
	json.NewEncoder(w).Encode(std)

}
