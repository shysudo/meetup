package get

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetParticipantListHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/meetup/participants", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetParticipantListHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"name":"Gireesh","age":28,"dob":"1992-06-04T00:00:00Z","profession":"Employed","locality":"Bangalore","number_of_guest":1,"address":"HSR layout, Bangalore"},{"name":"Veerendra","age":27,"dob":"1993-07-10T00:00:00Z","profession":"Student","locality":"Bangalore","number_of_guest":2,"address":"Banashankari, Bangalore"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
