package main

import(
	"testing"
	"net/http/httptest"
	"net/http"
	"last.com/controller"
	"github.com/gorilla/mux"
	"fmt"
)


func TestGetStudentByID(t *testing.T){
	req, err := http.NewRequest("GET", "/student", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	vars := map[string] string{
		"id":"498084",
	}
	req = mux.SetURLVars(req,vars)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetStudentByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	fmt.Println("rr.body in id = ",rr.Body.String())

	// Check the response body is what we expect.
	/*
	expected := `{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb2405@gmail.com","phone_number":"0987654321"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}*/
	
}


