package main

import(
	"testing"
	"net/http/httptest"
	"net/http"
	"last.com/controller"
	"fmt"
	"github.com/gorilla/mux"

)
func TestDeleteStudent(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/student/498089", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	vars := map[string]string{
		"id":"498089",
	}
	req = mux.SetURLVars(req,vars)
	handler := http.HandlerFunc(controller.DeleteStudent)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	fmt.Println(rr.Body.String())
	
}