package main


import (
	"testing"
	"net/http/httptest"
	"net/http"
	"last.com/controller"
	"fmt"
)

func TestGetStudent(t *testing.T){
	req,err := http.NewRequest("GET","/student",nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetStudent)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK{
		t.Errorf("handler return wrong status code: got %v want %v",status,http.StatusOK)
	}
	fmt.Println("rr =",rr)
}
