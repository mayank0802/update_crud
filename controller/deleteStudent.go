package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"last.com/config"


)

func DeleteStudent(w http.ResponseWriter, r *http.Request){
	db := config.Getdb()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	row, errr := db.Query("DELETE FROM Student where E_NO=?", id)
	if errr != nil {
		fmt.Println(errr)
	}
	w.Header().Set("Contet-Type", "application/json")
	json.NewEncoder(w).Encode(row)

}