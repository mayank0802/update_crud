package main
import(
	"last.com/controller"
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
)


func TestCreateEntry(t *testing.T) {

	var jsonStr = []byte(`{"name":"xyz","age":9,"subject":"xyz","class":"abcd"}`)

	req, err := http.NewRequest("POST", "/Student", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateStudent)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}