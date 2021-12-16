package controller


import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
 
    "../model"
 
    "../config"
)

func getStudent(w http.ResponseWriter, r *http.Request){
	var students model.Student
	var arrStudents []model.Student

	db =: config.Connect()
	defer db.close()
	 row, err := db.Query("SELECT E_NO, NAME FROM Students")
	 
	 if err != nil{
		 	log.Fatal(err.Error())
	 }

	 for row.Next() {
		err = row.Scan(&students.ID,&students.Name)
		if(err != nil)
			log.Fatal(err.Error())
		else 
			arrStudents  = append(arrStudents, students)			
	 }
	 else {
		 arrStudents =append(arrStudents, employee)
	 }
	 w.Header().Set("Content-Type", "application/json")
	 json.NewEncoder(w).Encode(arrStudents)
/*	params := mux.Vars(r) // get parameter
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
	*/
}
