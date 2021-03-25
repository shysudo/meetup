package add

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateEntry(t *testing.T) {

	var jsonStr = []byte(`{
    "name" : "Gireesh K H",
    "age" : 10,
    "number_of_guest" : 1,
    "profession" : "Student",
    "address" : "gireesh kademane bangalore",
    "locality" : "Bangalore",
    "dob" : "1992-07-02T13:45:28Z"
}`)

	req, err := http.NewRequest("POST", "/meetup/participants", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RegisterParticipantHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{
    "name" : "Gireesh K H",
    "age" : 10,
    "number_of_guest" : 1,
    "profession" : "Student",
    "address" : "gireesh kademane bangalore",
    "locality" : "Bangalore",
    "dob" : "1992-07-02T13:45:28Z"
}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
